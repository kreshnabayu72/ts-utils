package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/goccy/go-json"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/config"
	"github.com/segmentio/kafka-go"
)

// CDCKind represents the type of database change captured by Change Data Capture.
type CDCKind string

const (
	CDCKindInsert CDCKind = "insert" // A new row was inserted
	CDCKindUpdate CDCKind = "update" // An existing row was modified
	CDCKindDelete CDCKind = "delete" // A row was deleted

	CDCCustomKindMove CDCKind = "move" // A row was moved from one schema to another
)

// CDCMessage represents a Change Data Capture message produced by PeerDB.
// PeerDB captures PostgreSQL WAL (Write-Ahead Log) changes and publishes them to Kafka.
//
// Example message from topic "tagsamurai_iam.users":
//
//	{
//	    "kind": "insert",
//	    "new": {
//	        "id": 2029086878094856200,
//	        "email": "tyler@example.com",
//	        "full_name": "Tyler Lang"
//	    },
//	    "source": "qtera.users"
//	}
//
// Fields:
//   - Kind: The type of change (insert, update, delete)
//   - New:  The new row data (present for insert and update)
//   - Old:  The old row data (present for delete; contains the row before deletion)
//   - Source: The schema-qualified table name in the source database (e.g., "qtera.users")
//     The source schema part ("qtera" or "public") is used in combination with ReplicationFlow
//     to determine the actual target schema.
type CDCMessage struct {
	Kind   CDCKind        `json:"kind"`
	New    map[string]any `json:"new,omitempty"`
	Old    map[string]any `json:"old,omitempty"`
	Source string         `json:"source"`
}

// ReplicationConfig holds configuration for a single CDC replication instance.
// One CDCReplicator gets one ReplicationConfig.
type ReplicationConfig struct {
	SourceDBName    string                        // Source database name, used to build Kafka topic names (e.g., "tagsamurai_iam")
	TableMappings   map[string]TargetTableMapping // Maps source table name (e.g., "users") -> how to replicate it
	ConsumerGroup   string                        // Kafka consumer group ID for offset tracking
	ValidClients    []string                      // List of valid clients for this table.
	ReplicationFlow ReplicationFlow
}

// TargetTableMapping describes how to replicate a single table.
// It defines the target table structure and optional transformation logic.
type TargetTableMapping struct {
	PrimaryKey      string                              // Primary key column name. Defaults to "id" if empty.
	Columns         []string                            // Columns to replicate. Only these columns are extracted from CDC messages.
	ColumnTypes     map[string]string                   // Maps column name -> SQL type (e.g., "id" -> "BIGINT"). Used for table creation.
	TargetTableName string                              // The table name in the target database to write to.
	Transform       func(map[string]any) map[string]any // Optional: transforms the CDC data before writing. If nil, data is used as-is.
}

// primaryKeyOrDefault returns the primary key or "id" if not set.
func (tm *TargetTableMapping) primaryKeyOrDefault() string {
	if tm.PrimaryKey == "" {
		return "id"
	}
	return tm.PrimaryKey
}

// columnsWithPrimaryKey returns the columns including the primary key.
func (tm *TargetTableMapping) columnsWithPrimaryKey() []string {
	pk := tm.primaryKeyOrDefault()
	hasPK := slices.Contains(tm.Columns, pk)
	if hasPK {
		return tm.Columns
	}
	colsWithPK := make([]string, 0, len(tm.Columns)+1)
	colsWithPK = append(colsWithPK, tm.Columns...)
	colsWithPK = append(colsWithPK, pk)
	return colsWithPK
}

// CDCReplicatorInterface defines the contract for a CDC replicator.
// The default implementation is CDCReplicator. This interface exists
// primarily to allow mocking in tests.
type CDCReplicatorInterface interface {
	Start() error // Blocks and consumes messages until context is cancelled
	Stop()        // Cancels the context, causing Start() to return

	getMappedTopics() []string
	processMessage(topic string, value []byte) error
	parseTopic(topic string) (source string, err error)
	upsertRecord(ctx context.Context, mapping TargetTableMapping, data map[string]any, schemaName string) error
	deleteRecord(ctx context.Context, mapping TargetTableMapping, data map[string]any, schemaName string) error
}

// CDCReplicator consumes CDC messages from Kafka and replicates them into PostgreSQL.
//
// Simple Data Flow:
//
//	[Database A] -> [PeerDB] -> [Kafka Topic] -> [CDC Replicator] -> [Database B]
//
// How it works (The Simple Version):
//  1. PeerDB watches Database A for any changes (Insert, Update, Delete).
//  2. PeerDB sends that change as a JSON message to a Kafka topic.
//  3. CDC Replicator reads the message and looks at "where" it came from.
//  4. Replicator picks the right folder (schema) in Database B to write to.
//  5. Replicator runs the change (Upsert or Delete) in Database B.
//
// Smart Routing (Where does the data go?):
//
//	This replicator is smart. It knows how to send data to the right tenant:
//
//	- Tenant to Tenant:
//	  Data from "qtera.users" goes directly to folder "qtera" in the new DB.
//
//	- Global to All Tenants (Broadcast):
//	  Data like "public.skus" (global items) has no specific owner (has no column/field client_id in data).
//	  The replicator sends a copy to EVERY active tenant (from r.replicationConfig.ValidClients) automatically.
//
//	- Targeted Tenant (by Field):
//	  Data like "public.stocks" has a "client_id" field in the data.
//	  The replicator reads that field and sends the data to that specific tenant.
//
//	- Moving Tenants (The "Move" Kind):
//	  If a record changes its "client_id" (e.g., from Apple to Google), the replicator:
//	  1. Deletes the old record from the "Apple" folder.
//	  2. Inserts the new record into the "Google" folder.
//
//	- Skip Rule:
//	  If a record is supposed to have a client_id but the field is empty or null, it's ignored.
type CDCReplicator struct {
	ctx               context.Context        // Context for managing lifecycle; cancelled on Stop()
	cancel            context.CancelFunc     // Cancel function to stop the replicator
	pool              PgxPoolInterface       // Postgres connection pool for the TARGET database
	brokers           string                 // Kafka brokers list (e.g., "redpanda:9092"). Comma-separated for multiple.
	replicationConfig *ReplicationConfig     // Configuration: which topics to consume, how to map tables
	configOptions     []config.ConfigOptions // Optional configuration options
	sem               chan struct{}          // Semaphore to limit concurrent DB operations
}

// NewCDCReplicator creates a new CDCReplicator instance.
//
// Parameters:
//   - parentCtx: Parent context. Replicator creates a child context with cancel for lifecycle management.
//   - postgreService: PostgreSQL service for the TARGET database. The replicator uses its connection pool.
//   - brokers: Kafka/Redpanda broker addresses (e.g., "localhost:9092" or "redpanda:9092").
//   - replicationConfig: Defines which topics to consume and how to map tables.
func NewCDCReplicator(
	parentCtx context.Context,
	postgreService PostgreSQLServicer,
	brokers string,
	replicationConfig *ReplicationConfig,
	configOptions ...config.ConfigOptions,
) CDCReplicatorInterface {
	slog.Info(
		"Initializing CDC Replicator for service",
		slog.String("service_name", replicationConfig.SourceDBName),
	)

	ctx, cancel := context.WithCancel(parentCtx)

	return &CDCReplicator{
		ctx:               ctx,
		cancel:            cancel,
		pool:              postgreService.GetPool(),
		brokers:           brokers,
		replicationConfig: replicationConfig,
		configOptions:     configOptions,
		sem:               make(chan struct{}, 10), // Limit total concurrent replications to 10
	}
}

// Start begins consuming CDC messages from Kafka. This is a BLOCKING call.
// It creates one Kafka reader goroutine per topic, all sharing the same consumer group.
// Each reader runs an infinite loop: fetch message -> process -> commit offset.
//
// The method only returns when:
//   - The context is cancelled (via Stop())
//   - All reader goroutines have exited
//
// Message processing guarantees:
//   - Offsets are only committed AFTER successful processing (at-least-once delivery)
//   - If processing fails, the offset is NOT committed, so the message will be retried
func (r *CDCReplicator) Start() error {
	topics := r.getMappedTopics()
	slog.Debug("Starting CDC Replicator", slog.Any("topics", topics))

	// Use WaitGroup to manage multiple topic readers
	var wg sync.WaitGroup
	for _, topic := range topics {
		wg.Add(1)
		go func(topic string) {
			defer wg.Done()
			reader := kafka.NewReader(kafka.ReaderConfig{
				Brokers:        strings.Split(r.brokers, ","),
				GroupID:        r.replicationConfig.ConsumerGroup,
				Topic:          topic,
				CommitInterval: 0, // Disable auto-commit; we commit manually after successful processing
				MinBytes:       1,
				MaxBytes:       10e6,

				// Custom logger, integrate with slog
				Logger: kafka.LoggerFunc(
					func(msg string, args ...any) {
						if len(args) > 0 {
							msg = fmt.Sprintf(msg, args...)
						}
						slog.Debug(msg)
					},
				),
			})
			defer reader.Close() // Ensure reader is closed when goroutine exits

			for {
				// Fetch message from Kafka (blocks until a message is available or context is cancelled)
				m, err := reader.FetchMessage(r.ctx)
				if err != nil {
					if errors.Is(err, context.Canceled) {
						slog.Info(
							"Kafka reader stopped due to context cancellation.",
							slog.String("topic", topic),
						)
						return
					}

					slog.Debug(
						"Kafka fetch error",
						slog.String("topic", topic),
						slog.Any("error", err),
					)
					time.Sleep(time.Second)
					continue
				}

				// Topic & Value extraction
				msgTopic := m.Topic
				value := m.Value

				if processErr := r.processMessage(msgTopic, value); processErr != nil {
					slog.Debug(
						"Error processing message. Will NOT commit offset.",
						slog.String("topic", msgTopic),
						slog.Any("error", processErr),
					)
				} else {
					// Commit offset only if processing is successful
					if commitErr := reader.CommitMessages(r.ctx, m); commitErr != nil {
						slog.Debug("[%s] Error committing offset: %v", msgTopic, commitErr)
					}
				}
			}
		}(topic)
	}

	wg.Wait()
	return nil
}

// Stop gracefully stops the CDC replication process.
// It cancels the context, which causes all Kafka readers (in Start()) to exit
// their FetchMessage loop with a context.Canceled error.
func (r *CDCReplicator) Stop() {
	slog.Info(
		"Stopping CDC Replicator for service",
		slog.String("service_name", r.replicationConfig.SourceDBName),
	)
	if r.cancel != nil {
		r.cancel()
	}
}

// getMappedTopics returns the list of Kafka topic names this replicator should subscribe to.
// Topic format: "{sourceDBName}.{tableName}" (e.g., "tagsamurai_iam.users").
// This matches the topic naming convention used by PeerDB when publishing CDC events.
func (r *CDCReplicator) getMappedTopics() []string {
	slog.Debug("Getting all topics for service:", slog.String("service_name", r.replicationConfig.SourceDBName))
	topics := []string{}
	tenantCfg := r.replicationConfig
	for source := range tenantCfg.TableMappings {
		topic := fmt.Sprintf("%s.%s", tenantCfg.SourceDBName, source)
		topics = append(topics, topic)
	}
	return topics
}

type CDCSchemaPair struct {
	oldSchema string
	newSchema string
}

// processMessage handles a single CDC message from Kafka.
//
// Flow:
//  1. Parse the Kafka topic name to extract the source table name
//     Example: topic "tagsamurai_iam.users" -> table "users"
//  2. Look up the TargetTableMapping for this table
//  3. Parse the JSON message into a CDCMessage struct
//  4. Extract the source schema from CDCMessage.Source
//     Example: source "qtera.users" -> schema "qtera"
//     Example: source "public.device_names" -> schema "public"
//  5. Determine the actual target schema using the configured ReplicationFlow.
//     For example, ReplicationFlowNonTenantToTenant extracts a client_id from the record
//     to route to a specific schema instead of using the source schema.
//  6. Route to upsertRecord (for insert/update) or deleteRecord (for delete)
//     with the target schema, so data is written to the correct location.
func (r *CDCReplicator) processMessage(topic string, value []byte) error {
	// Parse topic to get source table name
	source, err := r.parseTopic(topic)
	if err != nil {
		return err
	}

	tableMap, ok := r.replicationConfig.TableMappings[source]
	if !ok {
		return fmt.Errorf("table mapping for %s not found", source)
	}
	slog.Debug("Table mapping found for %s: %+v", source, tableMap)

	var msg CDCMessage
	if err := json.Unmarshal(value, &msg); err != nil {
		return fmt.Errorf("parse failed: %w", err)
	}

	// Extract schema name from the CDC message source field.
	// PeerDB sets source to "{schema}.{table}" (e.g., "qtera.users" or "public.device_names").
	// We use the source schema name with the ReplicationFlow to determine
	// the actual target schema and set search_path before executing queries.
	splitSource := strings.Split(msg.Source, ".")
	var tenantIdSource string
	if len(splitSource) == 2 && splitSource[0] != "" {
		tenantIdSource = splitSource[0] // Tenant schema (e.g., "qtera")
	} else {
		slog.Error("Invalid source format", slog.String("source", msg.Source))
		return fmt.Errorf("invalid source format: %s", msg.Source)
	}

	slog.Debug("Processing CDC message for table:", slog.String("table", source), slog.Any("message", msg))

	isValidClients := func(clientId string) bool {
		return slices.Contains(r.replicationConfig.ValidClients, clientId)
	}

	var targetSchemas []CDCSchemaPair
	switch r.replicationConfig.ReplicationFlow {
	case ReplicationFlowTenantToTenant: // Target is same as source (tenant to tenant)
		if !isValidClients(tenantIdSource) {
			return nil
		}
		// Ensure both oldSchema and newSchema assigned for upsert & delete
		targetSchemas = []CDCSchemaPair{
			{
				oldSchema: tenantIdSource,
				newSchema: tenantIdSource,
			},
		}

	case ReplicationFlowNonTenantToNonTenant: // Source should be "public" since this is `from non-tenant` flow
		if tenantIdSource != "public" {
			return nil
		}
		// Ensure both oldSchema and newSchema assigned for upsert & delete
		targetSchemas = []CDCSchemaPair{
			{
				oldSchema: "public",
				newSchema: "public",
			},
		}

	case ReplicationFlowNonTenantToTenant:
		newData := msg.New
		oldData := msg.Old

		newClientIdVal, newExists := newData["client_id"]
		oldClientIdVal, oldExists := oldData["client_id"]

		if !newExists && !oldExists {
			// Rule: No client_id field AND flow from non-tenant (public) -> Broadcast to all tenants (e.g., device_names & skus)
			if tenantIdSource == "public" {
				targetSchemas = make([]CDCSchemaPair, len(r.replicationConfig.ValidClients))
				for i, v := range r.replicationConfig.ValidClients {
					targetSchemas[i] = CDCSchemaPair{
						oldSchema: v, // To handle deletion for this Rule (e.g., delete from schema "qtera")
						newSchema: v, // To handle insertion for this Rule (e.g., insert into schema "qtera")
					}
				}
			}
		} else {
			// Rule: client_id field exists but empty/null -> Skip this row (skipped by if else conditioning)
			newClientIdStr := fmt.Sprintf("%v", newClientIdVal)
			oldClientIdStr := fmt.Sprintf("%v", oldClientIdVal)

			if msg.Kind == CDCKindDelete { // Pure Deletion, no need newSchema
				if oldExists && !isValidClients(oldClientIdStr) {
					return nil
				}
				targetSchemas = []CDCSchemaPair{
					{
						oldSchema: oldClientIdStr, // e.g., delete from schema "qtera"
						newSchema: "",
					},
				}
			} else if newExists && oldExists && newClientIdStr != oldClientIdStr { // Move, both field in new & old exist but value differ
				if !isValidClients(newClientIdStr) || !isValidClients(oldClientIdStr) {
					return nil
				}
				targetSchemas = []CDCSchemaPair{
					{
						oldSchema: oldClientIdStr, // e.g., delete from schema "qtera"
						newSchema: newClientIdStr, // e.g., insert into schema "0001"
					},
				}
				msg.Kind = CDCCustomKindMove
			} else if newExists && isValidClients(newClientIdStr) { // Pure Insertion, no need oldSchema
				targetSchemas = []CDCSchemaPair{
					{
						oldSchema: "",
						newSchema: newClientIdStr, // e.g., insert into schema "0001"
					},
				}
			}
		}

	case ReplicationFlowTenantToNonTenant: // Target is "public"
		if !isValidClients(tenantIdSource) {
			return nil
		}
		// Ensure both oldSchema and newSchema assigned for upsert & delete
		targetSchemas = []CDCSchemaPair{
			{
				newSchema: "public",
				oldSchema: "public",
			},
		}

	default:
		slog.Error("Unsupported replication flow", slog.String("flow", string(r.replicationConfig.ReplicationFlow)))
		return nil
	}

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(r.ctx)
	defer cancel()

	errCh := make(chan error, len(targetSchemas))

	for _, schemaPair := range targetSchemas {
		wg.Add(1)
		go func(sp CDCSchemaPair) {
			defer wg.Done()

			// Acquire semaphore with context awareness
			select {
			case r.sem <- struct{}{}:
				defer func() { <-r.sem }()
			case <-ctx.Done():
				return // Stop before starting if an error occurred elsewhere
			}

			var execErr error
			switch msg.Kind {
			case CDCKindInsert, CDCKindUpdate:
				if sp.newSchema != "" { // Just defensive programming, pretty sure newSchema assignment already handled above
					execErr = r.upsertRecord(ctx, tableMap, msg.New, sp.newSchema)
				}
			case CDCKindDelete:
				if sp.oldSchema != "" { // Just defensive programming, pretty sure oldSchema assignment already handled above
					execErr = r.deleteRecord(ctx, tableMap, msg.Old, sp.oldSchema)
				}
			case CDCCustomKindMove:
				// Validation empty string is just defensive programming, pretty sure both newSchema and oldSchema assignment already handled above
				if sp.oldSchema != "" {
					execErr = r.deleteRecord(ctx, tableMap, msg.Old, sp.oldSchema)
				}
				if execErr == nil && sp.newSchema != "" {
					execErr = r.upsertRecord(ctx, tableMap, msg.New, sp.newSchema)
				}
			default:
				slog.Debug("Unknown CDC kind recorded", slog.String("kind", string(msg.Kind)))
				return
			}

			if execErr != nil {
				cancel() // Fail Fast: Cancel all other goroutines immediately
				select {
				case errCh <- fmt.Errorf("replication failed for old schema %s & new schema %s: %w", sp.oldSchema, sp.newSchema, execErr):
				default:
				}
			}
		}(schemaPair)
	}

	wg.Wait()
	close(errCh)

	// Collect first error if any occurred
	for e := range errCh {
		if e != nil {
			return e
		}
	}

	return nil
}

// parseTopic extracts the table name from a Kafka topic name.
//
// Topic format: "{sourceDBName}.{tableName}"
// Example: "tagsamurai_iam.users" -> returns "users"
// Example: "tagsamurai_hardware_tagging.handheld_stocks" -> returns "handheld_stocks"
//
// The returned table name is used to look up the TargetTableMapping.
func (r *CDCReplicator) parseTopic(topic string) (source string, err error) {
	parts := strings.SplitN(topic, ".", 2)
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid topic format: %s", topic)
	}

	return parts[1], nil
}

// upsertRecord inserts or updates a record in the target database.
//
// This method:
//  1. Extracts only the columns defined in the TargetTableMapping from the CDC data
//  2. Builds an UPSERT query: INSERT ... ON CONFLICT (pk) DO UPDATE SET ...
//  3. Executes within a transaction with SET LOCAL search_path for schema isolation
//
// Parameters:
//   - mapping: Defines which columns to extract and the target table name
//   - data: The row data from the CDC message (CDCMessage.New)
//   - schemaName: The target schema name (e.g., "qtera" or "public")
func (r *CDCReplicator) upsertRecord(
	ctx context.Context,
	mapping TargetTableMapping,
	data map[string]any,
	schemaName string,
) error {
	var row map[string]any

	// Apply transformation if provided, otherwise extract only mapped columns
	if mapping.Transform != nil {
		row = mapping.Transform(data)
	} else {
		row = make(map[string]any)
		for _, col := range mapping.columnsWithPrimaryKey() {
			if val, ok := data[col]; ok {
				row[col] = val
			}
		}
	}

	// Ensure primary key is always included in the row data
	primaryKey := mapping.primaryKeyOrDefault()
	if _, ok := row[primaryKey]; !ok {
		if val, ok2 := data[primaryKey]; ok2 {
			row[primaryKey] = val
		}
	}

	// Build the UPSERT query dynamically from the mapped columns
	cols := []string{}
	vals := []any{}
	placeholders := []string{}
	updates := []string{}
	i := 1

	for _, col := range mapping.columnsWithPrimaryKey() {
		if val, ok := row[col]; ok {
			cols = append(cols, col)
			vals = append(vals, val)
			placeholders = append(placeholders, fmt.Sprintf("$%d", i))
			if col != primaryKey {
				updates = append(updates, fmt.Sprintf("%s = EXCLUDED.%s", col, col))
			}
			i++
		}
	}

	// Fallback: if Columns are not defined (empty), use all keys in row
	if len(cols) == 0 {
		i = 1
		for col, val := range row {
			cols = append(cols, col)
			vals = append(vals, val)
			placeholders = append(placeholders, fmt.Sprintf("$%d", i))
			if col != primaryKey {
				updates = append(updates, fmt.Sprintf("%s = EXCLUDED.%s", col, col))
			}
			i++
		}
	}

	// Build the final SQL query.
	// If no update columns exist (only PK), use DO NOTHING to avoid syntax error.
	var query string
	if len(updates) == 0 {
		query = fmt.Sprintf(
			"INSERT INTO %s (%s) VALUES (%s) ON CONFLICT (%s) DO NOTHING",
			mapping.TargetTableName,
			strings.Join(cols, ", "),
			strings.Join(placeholders, ", "),
			primaryKey,
		)
	} else {
		query = fmt.Sprintf(
			"INSERT INTO %s (%s) VALUES (%s) ON CONFLICT (%s) DO UPDATE SET %s",
			mapping.TargetTableName,
			strings.Join(cols, ", "),
			strings.Join(placeholders, ", "),
			primaryKey,
			strings.Join(updates, ", "),
		)
	}

	// Execute within a transaction with tenant-specific search_path.
	// SET LOCAL search_path only affects the current transaction, so it's safe
	// for concurrent use - other queries on the same connection pool are not affected.
	slog.Debug("Pool: ", slog.Any("pool", r.pool))
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction failed: %w", err)
	}
	defer tx.Rollback(ctx)

	// Set search_path to the tenant schema.
	// Since tenantID comes from the CDC source field (not user input), SQL injection isn't a concern here
	_, err = tx.Exec(ctx, fmt.Sprintf("SET LOCAL search_path = %s", schemaName))
	if err != nil {
		return fmt.Errorf("set session search path failed: %w", err)
	}
	_, err = tx.Exec(ctx, query, vals...)
	if err != nil {
		return fmt.Errorf("upsert failed: %w", err)
	}
	tx.Commit(ctx)

	slog.Debug(
		"Upserted record into table:",
		slog.String("table", mapping.TargetTableName),
		slog.Any("pk", row[primaryKey]),
	)
	return nil
}

// deleteRecord deletes a record from the target database based on the primary key.
//
// Similar to upsertRecord, this method:
//  1. Runs within a transaction with SET LOCAL search_path for schema isolation
//  2. Executes DELETE FROM {table} WHERE {pk} = $1
//
// Parameters:
//   - mapping: Defines the target table and primary key
//   - data: The old row data from the CDC message (CDCMessage.Old)
//   - schemaName: The target schema name (e.g., "qtera" or "public")
func (r *CDCReplicator) deleteRecord(
	ctx context.Context,
	mapping TargetTableMapping,
	data map[string]any,
	schemaName string,
) error {
	primaryKey := mapping.primaryKeyOrDefault()

	pkValue, ok := data[primaryKey]
	if !ok {
		return fmt.Errorf("primary key %s not found", primaryKey)
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE %s = $1", mapping.TargetTableName, primaryKey)
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction failed: %w", err)
	}
	defer tx.Rollback(ctx)

	// Since tenantID comes from the CDC source field (not user input), SQL injection isn't a concern here
	_, err = tx.Exec(ctx, fmt.Sprintf("SET LOCAL search_path = %s", schemaName))
	if err != nil {
		return fmt.Errorf("set session search path failed: %w", err)
	}
	_, err = tx.Exec(ctx, query, pkValue)
	if err != nil {
		return fmt.Errorf("delete failed: %w", err)
	}
	tx.Commit(ctx)

	slog.Debug(
		"Deleted record from table:",
		slog.String("table", mapping.TargetTableName),
		slog.Any("pk", pkValue),
	)
	return nil
}
