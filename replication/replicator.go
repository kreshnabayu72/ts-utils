package replication

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"sync"
	"time"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/config"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/replication/api"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/service"
	"github.com/segmentio/kafka-go"
)

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
//
// Implementation Note: The struct is private to the package, but exposed via NewCDCReplicator
// which returns an api.CDCReplicatorInterface.
// NewCDCReplicator creates a new CDCReplicator instance.
//
// Parameters:
//   - parentCtx: Parent context. Replicator creates a child context with cancel for lifecycle management.
//   - postgreService: PostgreSQL service for the TARGET database. The replicator uses its connection pool.
//   - brokers: Kafka/Redpanda broker addresses (e.g., "localhost:9092" or "redpanda:9092").
//   - replicationConfig: Defines which topics to consume and how to map tables.
func NewCDCReplicator(
	parentCtx context.Context,
	postgreService service.PostgreSQLServicer,
	brokers string,
	replicationConfig *ReplicationConfig,
	configOptions ...config.ConfigOptions,
) api.CDCReplicatorInterface {
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
				if processErr := r.processMessage(m.Topic, m.Value); processErr != nil {
					slog.Debug(
						"Error processing message. Will NOT commit offset.",
						slog.String("topic", m.Topic),
						slog.Any("error", processErr),
					)
				} else {
					// Commit offset only if processing is successful
					if commitErr := reader.CommitMessages(r.ctx, m); commitErr != nil {
						slog.Debug("[%s] Error committing offset: %v", m.Topic, commitErr)
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
//
// For TenantToTenant and NonTenantToTenant flow, each client gets its own topic per table:
// "{sourceDBName}_{clientId}.{tableName}" (e.g., "tagsamurai_hardware_tagging_80000001.rfids").
// One topic is generated per client per table, matching the per-client destination
// naming used by PeerDB mirrors.
//
// For all other flows, a single shared topic per table is used:
// "{sourceDBName}.{tableName}" (e.g., "tagsamurai_iam.users").
func (r *CDCReplicator) getMappedTopics() []string {
	slog.Debug("Getting all topics for service:", slog.String("service_name", r.replicationConfig.SourceDBName))
	topics := []string{}
	for source := range r.replicationConfig.TableMappings {
		if r.replicationConfig.ReplicationFlow == ReplicationFlowTenantToTenant ||
			r.replicationConfig.ReplicationFlow == ReplicationFlowNonTenantToTenant {
			for _, clientId := range r.replicationConfig.ValidClients {
				topic := fmt.Sprintf("%s_%s.%s", r.replicationConfig.SourceDBName, clientId, source)
				topics = append(topics, topic)
			}
		} else {
			topic := fmt.Sprintf("%s.%s", r.replicationConfig.SourceDBName, source)
			topics = append(topics, topic)
		}
	}
	return topics
}
