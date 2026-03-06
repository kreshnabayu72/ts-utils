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
//     The schema name ("qtera") tells us which tenant this data belongs to.
//     If the schema is "public", the data is non-tenant-specific.
type CDCMessage struct {
	Kind   CDCKind        `json:"kind"`
	New    map[string]any `json:"new,omitempty"`
	Old    map[string]any `json:"old,omitempty"`
	Source string         `json:"source"`
}

// ReplicationConfig holds configuration for a single CDC replication instance.
// One CDCReplicator gets one ReplicationConfig.
type ReplicationConfig struct {
	SourceDBName  string                        // Source database name, used to build Kafka topic names (e.g., "tagsamurai_iam")
	TableMappings map[string]TargetTableMapping // Maps source table name (e.g., "users") -> how to replicate it
	ConsumerGroup string                        // Kafka consumer group ID for offset tracking
	ValidClients  []string                      // List of valid clients for this table.

	// TODO: uncomment if in the future need replication from all tenant schemas to 1 "public" schema.
	// Search global: "enable target public"
	// TargetPublic bool // Whether the table should be replicated to the public schema
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
	upsertRecord(mapping TargetTableMapping, data map[string]any, tenantID string) error
	deleteRecord(mapping TargetTableMapping, data map[string]any, tenantID string) error
}

// CDCReplicator consumes CDC messages from Kafka and replicates them into PostgreSQL.
//
// Data Flow Overview:
//
//	[Source DB] -> [PeerDB WAL Capture] -> [Kafka/Redpanda Topic] -> [CDCReplicator] -> [Target DB]
//
// How it works step by step:
//  1. PeerDB watches the source database's WAL (Write-Ahead Log) for changes
//  2. When a row is inserted/updated/deleted, PeerDB publishes a CDCMessage to a Kafka topic
//  3. Topic naming: "{sourceDBName}.{tableName}" (e.g., "tagsamurai_iam.users")
//  4. CDCReplicator subscribes to these topics via Kafka consumer groups
//  5. For each message, it extracts the tenant schema from CDCMessage.Source
//     (e.g., "qtera.users" -> schema is "qtera")
//  6. It then runs SET LOCAL search_path = {schema} inside a transaction
//  7. Finally, it executes an UPSERT or DELETE into the target database
//
// Schema-aware routing (tenant isolation):
//   - Tenant-dependent data: source field is like "qtera.users" -> writes to schema "qtera"
//   - Non-tenant data: source field is like "public.device_names" -> writes to schema "public"
type CDCReplicator struct {
	ctx               context.Context        // Context for managing lifecycle; cancelled on Stop()
	cancel            context.CancelFunc     // Cancel function to stop the replicator
	pool              PgxPoolInterface       // Postgres connection pool for the TARGET database
	brokers           string                 // Kafka brokers list (e.g., "redpanda:9092"). Comma-separated for multiple.
	replicationConfig *ReplicationConfig     // Configuration: which topics to consume, how to map tables
	configOptions     []config.ConfigOptions // Optional configuration options
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

// processMessage handles a single CDC message from Kafka.
//
// Flow:
//  1. Parse the Kafka topic name to extract the source table name
//     Example: topic "tagsamurai_iam.users" -> table "users"
//  2. Look up the TargetTableMapping for this table
//  3. Parse the JSON message into a CDCMessage struct
//  4. Extract the schema name from CDCMessage.Source
//     Example: source "qtera.users" -> schema "qtera" (tenant-specific)
//     Example: source "public.device_names" -> schema "" (public, non-tenant)
//  5. Route to upsertRecord (for insert/update) or deleteRecord (for delete)
//     with the extracted schema name, so data is written to the correct tenant schema
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
	// We use the schema name to set search_path before executing queries,
	// so data goes into the correct tenant-specific schema in the target database.
	splitSource := strings.Split(msg.Source, ".")
	var tenantIdSource string
	if len(splitSource) == 2 && splitSource[0] != "" {
		tenantIdSource = splitSource[0] // Tenant schema (e.g., "qtera")
	} else {
		slog.Error("Invalid source format", slog.String("source", msg.Source))
		return fmt.Errorf("invalid source format: %s", msg.Source)
	}

	slog.Debug("Processing CDC message for table:", slog.String("table", source), slog.Any("message", msg))

	// Scenario is either:
	// 1. Tenant to tenant (must same tenant) or Non-tenant(public) to Non-tenant(public)
	// 2. Non-tenant to Tenant (schema public to any valid tenant)
	// 3. Tenant to Non-tenant (hypothetical based on assumption, not implemented yet)

	isValidClients := func(clientId string) bool {
		return slices.Contains(r.replicationConfig.ValidClients, clientId)
	}

	var targetSchema string
	switch {
	case isValidClients(tenantIdSource): // Cover both scenario 1 & 3
		// TODO: uncomment if in the future need replication from all tenant schemas to 1 "public" schema.
		// Search global: "enable target public"
		// if r.replicationConfig.TargetPublic {
		// 	// Scenario 3 (Tenant to Non-Tenant)
		// 	targetSchema = "public"
		// } else {
		// Scenario 1 (Tenant to Tenant)
		targetSchema = tenantIdSource
		// }
	case !isValidClients(tenantIdSource) && tenantIdSource == "public": // Cover both scenario 1 & 2
		data := msg.New
		if msg.Kind == CDCKindDelete {
			data = msg.Old
		}

		clientId, exists := data["client_id"]

		// TODO: uncomment if in the future need replication from all tenant schemas to 1 "public" schema.
		// Search global: "enable target public"
		// if !exists || replicationConfig.TargetPublic { // Scenario 1 (Non-tenant to Non-tenant)
		if !exists { // Scenario 1 (Non-tenant to Non-tenant)
			targetSchema = "public"
		} else { // Scenario 2 (Non-tenant to Tenant)
			targetSchema = fmt.Sprintf("%v", clientId)
			if !isValidClients(targetSchema) {
				slog.Debug("Skipping unmapped client_id", slog.String("client_id", targetSchema))
				return nil
			}
		}
	default:
		slog.Error(
			"Invalid source format",
			slog.String("source", msg.Source),
		)
		return nil
	}

	switch msg.Kind {
	case "insert", "update":
		return r.upsertRecord(tableMap, msg.New, targetSchema)
	case "delete":
		return r.deleteRecord(tableMap, msg.Old, targetSchema)
	default:
		slog.Debug("Unknown Kind", slog.String("kind", string(msg.Kind)))
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
//  2. Filters out records that don't belong to the current tenant (via client_id check)
//  3. Builds an UPSERT query: INSERT ... ON CONFLICT (pk) DO UPDATE SET ...
//  4. Executes within a transaction with SET LOCAL search_path for tenant isolation
//
// Parameters:
//   - mapping: Defines which columns to extract and the target table name
//   - data: The row data from the CDC message (CDCMessage.New)
//   - tenantID: The tenant schema name (e.g., "qtera"). Empty string means public schema.
//
// Tenant filtering (client_id check):
//
//	For non-tenant-dependent sources (e.g., hardware_tagging public schema), a single CDC topic
//	may contain rows from ALL tenants. The client_id field is used to filter:
//	- If client_id exists in data and doesn't match tenantID -> skip (not our tenant)
//	- If client_id is missing -> replicate to all tenants (shared data like device_names)
func (r *CDCReplicator) upsertRecord(
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
	tx, err := r.pool.Begin(r.ctx)
	if err != nil {
		return fmt.Errorf("begin transaction failed: %w", err)
	}
	defer tx.Rollback(r.ctx)

	// Set search_path to the tenant schema.
	// Since tenantID comes from the CDC source field (not user input), SQL injection isn't a concern here
	_, err = tx.Exec(r.ctx, fmt.Sprintf("SET LOCAL search_path = %s", schemaName))
	if err != nil {
		return fmt.Errorf("set session search path failed: %w", err)
	}
	_, err = tx.Exec(r.ctx, query, vals...)
	if err != nil {
		return fmt.Errorf("upsert failed: %w", err)
	}
	tx.Commit(r.ctx)

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
//  1. Filters by client_id for non-tenant-dependent sources
//  2. Runs within a transaction with SET LOCAL search_path for tenant isolation
//  3. Executes DELETE FROM {table} WHERE {pk} = $1
//
// Parameters:
//   - mapping: Defines the target table and primary key
//   - data: The old row data from the CDC message (CDCMessage.Old)
//   - tenantID: The tenant schema name. Empty string means public schema.
func (r *CDCReplicator) deleteRecord(
	mapping TargetTableMapping,
	data map[string]any,
	schemaName string,
) error {
	primaryKey := mapping.primaryKeyOrDefault()
	// Tenant filtering: skip records that belong to a different tenant

	pkValue, ok := data[primaryKey]
	if !ok {
		return fmt.Errorf("primary key %s not found", primaryKey)
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE %s = $1", mapping.TargetTableName, primaryKey)
	tx, err := r.pool.Begin(r.ctx)
	if err != nil {
		return fmt.Errorf("begin transaction failed: %w", err)
	}
	defer tx.Rollback(r.ctx)

	// Since tenantID comes from the CDC source field (not user input), SQL injection isn't a concern here
	_, err = tx.Exec(r.ctx, fmt.Sprintf("SET LOCAL search_path = %s", schemaName))
	if err != nil {
		return fmt.Errorf("set session search path failed: %w", err)
	}
	_, err = tx.Exec(r.ctx, query, pkValue)
	if err != nil {
		return fmt.Errorf("delete failed: %w", err)
	}
	tx.Commit(r.ctx)

	slog.Debug(
		"Deleted record from table:",
		slog.String("table", mapping.TargetTableName),
		slog.Any("pk", pkValue),
	)
	return nil
}
