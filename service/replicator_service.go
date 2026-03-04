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
	"github.com/gofiber/fiber/v2/log"
	"github.com/segmentio/kafka-go"
)

type CDCKind string

const (
	CDCKindInsert CDCKind = "insert"
	CDCKindUpdate CDCKind = "update"
	CDCKindDelete CDCKind = "delete"
)

// CDCMessage represents a Change Data Capture message structure. From PeerDB format.
type CDCMessage struct {
	Kind   CDCKind        `json:"kind"`
	New    map[string]any `json:"new,omitempty"`
	Old    map[string]any `json:"old,omitempty"`
	Source string         `json:"source"`
}

// ReplicationConfig holds configuration for a single tenant replication.
type ReplicationConfig struct {
	SourceDBName  string                        // source database name (e.g., "tagsamurai_iam")
	TableMappings map[string]TargetTableMapping // key: source table name (e.g., "users")
	ConsumerGroup string
}

type TargetTableMapping struct {
	PrimaryKey      string
	Columns         []string
	ColumnTypes     map[string]string
	TargetTableName string
	Transform       func(map[string]any) map[string]any // Optional transformation function. If nil, no transformation is applied.
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

type CDCReplicatorInterface interface {
	Start() error
	Stop()

	getMappedTopics() []string
	processMessage(topic string, value []byte) error
	parseTopic(topic string) (source string, err error)
	upsertRecord(mapping TargetTableMapping, data map[string]any, tenantID string) error
	deleteRecord(mapping TargetTableMapping, data map[string]any, tenantID string) error
}

// CDCReplicator struct handles CDC replication from Kafka to Postgres.
type CDCReplicator struct {
	ctx               context.Context        // Context for managing lifecycle
	cancel            context.CancelFunc     // Cancel function to stop the replicator
	pool              PgxPoolInterface       // Postgres connection pool
	brokers           string                 // Kafka brokers list (e.g., "redpanda:9092"). Can be comma-separated for multiple brokers.
	replicationConfig *ReplicationConfig     // Replication configuration for a single tenant
	configOptions     []config.ConfigOptions // Optional configuration options (e.g., environment variables)
}

// NewCDCReplicator creates a new CDCReplicator instance.
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

// Start starts the CDC replication process.
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
				CommitInterval: 0, // Disable auto-commit
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
				// Fetch message from Kafka
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

// Stop stops the CDC replication process.
func (r *CDCReplicator) Stop() {
	slog.Info(
		"Stopping CDC Replicator for service",
		slog.String("service_name", r.replicationConfig.SourceDBName),
	)
	// Currently, there's no explicit stop logic as readers will stop when context is cancelled.
	// Just cancel the context.
	if r.cancel != nil {
		r.cancel()
	}
}

func (r *CDCReplicator) getMappedTopics() []string {
	slog.Debug("Getting all topics for service:", slog.String("service_name", r.replicationConfig.SourceDBName))
	topics := []string{}
	tenantCfg := r.replicationConfig
	for source := range tenantCfg.TableMappings {
		topic := fmt.Sprintf("%s.%s", tenantCfg.SourceDBName, source)
		topics = append(topics, topic)

		r.ensureTableExists(tenantCfg.TableMappings[source])
	}
	return topics
}

func (r *CDCReplicator) ensureTableExists(mapping TargetTableMapping) error {
	if len(mapping.ColumnTypes) == 0 {
		return nil // Skip if no type info provided
	}

	var definitions []string

	// Iterate over Columns to ensure order
	for _, col := range mapping.Columns {
		sqlType, ok := mapping.ColumnTypes[col]
		if !ok {
			sqlType = "TEXT" // Fallback default
		}

		def := fmt.Sprintf("%s %s", col, sqlType)
		if col == mapping.PrimaryKey {
			def += " PRIMARY KEY"
		}
		definitions = append(definitions, def)
	}

	query := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %s (%s)",
		mapping.TargetTableName,
		strings.Join(definitions, ", "),
	)

	_, err := r.pool.Exec(r.ctx, query)
	return err
}

// processMessage processes a single CDC message.
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
	if err := json.Unmarshal(value, &msg); err != nil { // Parse JSON message
		return fmt.Errorf("parse failed: %w", err)
	}

	schemaSource := msg.Source
	splitSource := strings.Split(schemaSource, ".")
	var schemaName string
	if splitSource[0] == "public" {
		schemaName = ""
	} else if splitSource[0] != "" {
		schemaName = splitSource[0]
	}

	slog.Debug("Processing CDC message for table:", slog.String("table", source), slog.Any("message", msg))

	switch msg.Kind {
	case "insert", "update":
		return r.upsertRecord(tableMap, msg.New, schemaName)
	case "delete":
		return r.deleteRecord(tableMap, msg.Old, schemaName)
	default:
		log.Debug("Unknown kind: %s", msg.Kind)
	}

	return nil
}

// parseTopic extracts the source table identifier from a topic.
// Topic format: {dbName}.{schema}.{table} → returns "{schema}.{table}"
func (r *CDCReplicator) parseTopic(topic string) (source string, err error) {
	// Expect format: service_name.table_name → split into [service_name, table_name]
	parts := strings.SplitN(topic, ".", 2)
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid topic format: %s", topic)
	}

	return parts[1], nil
}

func (r *CDCReplicator) upsertRecord(
	mapping TargetTableMapping,
	data map[string]any,
	tenantID string,
) error {
	var row map[string]any

	// Apply transformation if provided, otherwise use data as is
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

	// For non-tenant-dependent sources, filter by client_id when present.
	// If client_id is missing or null, the record is shared and should be replicated to all tenants.
	// Only skip when client_id exists, is non-nil, and doesn't match this tenant.
	//
	// Example replicating from non-tenant hardware tagging that contains all stocks all of tenants.
	// If column client_id is present (table handheld_stocks), it should be skipped if it doesn't match this tenant (tenant-dependent rows).
	// If column client_id is missing (table device_names), it should be replicated to all tenants (non-tenant-dependent rows).
	if tenantID != "" {
		if clientID, exists := data["client_id"]; exists {
			if fmt.Sprintf("%v", clientID) != tenantID {
				return nil
			}
		}
	}

	// Ensure primary key is always included in row (paranoia, though columnsWithPrimaryKey covers almost all cases)
	primaryKey := mapping.primaryKeyOrDefault()
	if _, ok := row[primaryKey]; !ok {
		if val, ok2 := data[primaryKey]; ok2 {
			row[primaryKey] = val
		}
	}

	cols := []string{}
	vals := []any{}
	placeholders := []string{}
	updates := []string{}
	i := 1

	// Always use columnsWithPrimaryKey to guarantee primary key is included in column set
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

	// Fallback: if Columns are not defined (empty), use all keys in row (but this is rare with enforced pk column)
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

	// FINAL DEFENSE: If update-set is empty (all-only PK present), do NOT place "SET " to avoid syntax error
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

	slog.Debug("Pool: ", slog.Any("pool", r.pool))
	tx, err := r.pool.Begin(r.ctx)
	if err != nil {
		return fmt.Errorf("begin transaction failed: %w", err)
	}
	defer tx.Rollback(r.ctx)
	if tenantID == "" {
		tenantID = "public"
	}
	_, err = tx.Exec(r.ctx, fmt.Sprintf("SET LOCAL search_path = %s", tenantID))
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

// deleteRecord deletes a record based on the primary key.
func (r *CDCReplicator) deleteRecord(
	mapping TargetTableMapping,
	data map[string]any,
	tenantID string,
) error {
	primaryKey := mapping.primaryKeyOrDefault()
	// For non-tenant-dependent sources, filter by client_id when present.
	if tenantID != "" {
		if clientID, exists := data["client_id"]; exists {
			if fmt.Sprintf("%v", clientID) != tenantID {
				return nil
			}
		}
	}

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
	if tenantID == "" {
		tenantID = "public"
	}
	_, err = tx.Exec(r.ctx, fmt.Sprintf("SET LOCAL search_path = %s", tenantID))
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
