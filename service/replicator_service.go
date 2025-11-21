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

	"gitea.qwertysystem.net/BETS/ts-utils/v2/config"
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
	TenantID      string                        // company code (e.g., "0001")
	SourceDBName  string                        // source database name (e.g., "0001_global_settings_service")
	TableMappings map[string]TargetTableMapping // key: source table name (e.g., "public.users")
	ConsumerGroup string
}

type TargetTableMapping struct {
	PrimaryKey      string
	Columns         []string
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
	parseTopic(topic string) (tenantID string, source string, err error)
	upsertRecord(mapping TargetTableMapping, data map[string]any) error
	deleteRecord(mapping TargetTableMapping, data map[string]any) error
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
		"Initializing CDC Replicator for tenant",
		slog.String("tenant_id", replicationConfig.TenantID),
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
		"Stopping CDC Replicator for tenant",
		slog.String("tenant_id", r.replicationConfig.TenantID),
	)
	// Currently, there's no explicit stop logic as readers will stop when context is cancelled.
	// Just cancel the context.
	if r.cancel != nil {
		r.cancel()
	}
}

func (r *CDCReplicator) getMappedTopics() []string {
	slog.Debug("Getting all topics for tenant:", slog.String("tenant_id", r.replicationConfig.TenantID))
	topics := []string{}
	tenantCfg := r.replicationConfig
	for source := range tenantCfg.TableMappings {
		topic := fmt.Sprintf("%s.%s", tenantCfg.SourceDBName, source)
		topics = append(topics, topic)
	}
	return topics
}

// processMessage processes a single CDC message.
func (r *CDCReplicator) processMessage(topic string, value []byte) error {
	// Parse topic to get source table name (ignore tenant parsing)
	_, source, err := r.parseTopic(topic)
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

	slog.Debug("Processing CDC message for table:", slog.String("table", source), slog.Any("message", msg))

	switch msg.Kind {
	case "insert", "update":
		return r.upsertRecord(tableMap, msg.New)
	case "delete":
		return r.deleteRecord(tableMap, msg.Old)
	default:
		log.Debug("Unknown kind: %s", msg.Kind)
	}

	return nil
}

func (r *CDCReplicator) parseTopic(topic string) (tenantID string, source string, err error) {
	// Expect format: 0001_dbname.public.table_name into [0001_dbname, public.table_name]
	parts := strings.SplitN(
		topic,
		".",
		2,
	)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid topic format: %s", topic)
	}

	dbName := parts[0]
	source = parts[1]

	dbParts := strings.Split(dbName, "_")
	if len(dbParts) >= 2 {
		tenantID = dbParts[0]
	} else {
		return "", "", fmt.Errorf("cannot extract tenant ID from: %s", dbName)
	}

	return tenantID, source, nil
}

func (r *CDCReplicator) upsertRecord(mapping TargetTableMapping, data map[string]any) error {
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
	_, err := r.pool.Exec(r.ctx, query, vals...)
	if err != nil {
		return fmt.Errorf("upsert failed: %w", err)
	}

	slog.Debug(
		"Upserted record into table:",
		slog.String("table", mapping.TargetTableName),
		slog.Any("pk", row[primaryKey]),
	)
	return nil
}

// deleteRecord deletes a record based on the primary key.
func (r *CDCReplicator) deleteRecord(mapping TargetTableMapping, data map[string]any) error {
	primaryKey := mapping.primaryKeyOrDefault()
	pkValue, ok := data[primaryKey]
	if !ok {
		return fmt.Errorf("primary key %s not found", primaryKey)
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE %s = $1", mapping.TargetTableName, primaryKey)
	_, err := r.pool.Exec(r.ctx, query, pkValue)
	if err != nil {
		return fmt.Errorf("delete failed: %w", err)
	}

	slog.Debug(
		"Deleted record from table:",
		slog.String("table", mapping.TargetTableName),
		slog.Any("pk", pkValue),
	)
	return nil
}
