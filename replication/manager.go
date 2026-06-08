package replication

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"sync"
	"time"

	"gitea.qwertysystem.net/Wangsit-Developer/be-core-utils/config"
	core_replication "gitea.qwertysystem.net/Wangsit-Developer/be-core-utils/replication"
	"gitea.qwertysystem.net/Wangsit-Developer/be-core-utils/replication/api"
	"gitea.qwertysystem.net/Wangsit-Developer/be-core-utils/service"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/constant"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/enterprise_management_service/model"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/replication/registry"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
)

// ReplicationManager orchestrates CDC (Change Data Capture) replication for a service.
type ReplicationManager struct {
	replicators     []api.CDCReplicatorInterface // All active replicators managed by this manager
	cfg             config.ConfigOptions         // Application config (contains KAFKA_BROKERS, etc.)
	dbSvc           service.PostgreSQLServicer   // Target database connection - where replicated data is written to
	temporalService *service.TemporalService     // Used to fetch the list of tenant client IDs via Temporal workflow
}

// NewReplicationManager creates a new ReplicationManager.
func NewReplicationManager(
	cfg config.ConfigOptions,
	temporalService *service.TemporalService,
	dbSvc service.PostgreSQLServicer,
) *ReplicationManager {
	return &ReplicationManager{
		replicators:     make([]api.CDCReplicatorInterface, 0),
		cfg:             cfg,
		dbSvc:           dbSvc,
		temporalService: temporalService,
	}
}

// StartReplication starts CDC replication for a specific source database.
func (rm *ReplicationManager) StartReplication(
	targetServiceName string,
	sourceDBName string,
	replicationFlow api.ReplicationFlow,
) error {
	slog.Debug("Entering StartReplication", "target_service", targetServiceName)
	clients := rm.GetClients()
	slog.Debug("GetClients returned", "clients_count", len(clients))

	if len(clients) == 0 {
		slog.Warn("No clients found", "service_name", targetServiceName)
		return nil
	}

	// Internal lookup from Registry
	replicatedTables := registry.ReplicatedTables[targetServiceName][api.ReplicationKey{
		Flow:         replicationFlow,
		SourceDBName: sourceDBName,
	}]

	if len(replicatedTables) == 0 {
		slog.Warn("No replicated tables provided", "service_name", targetServiceName, "flow", replicationFlow)
		return nil
	}

	tableMappings := make(map[string]core_replication.TargetTableMapping, len(replicatedTables))
	for _, table := range replicatedTables {
		tableMappings[table.GetTableName()] = core_replication.TargetTableMapping{
			PrimaryKey:      "id",
			Columns:         table.GetColumns(),
			TargetTableName: table.GetTableName(),
			ColumnTypes:     table.GetColumnTypes(),
		}
	}

	sanitizedTarget := strings.ReplaceAll(targetServiceName, "-", "_")
	sanitizedSource := strings.ReplaceAll(string(sourceDBName), "-", "_")
	consumerGroup := fmt.Sprintf("%s_%s_replicator", sanitizedTarget, sanitizedSource)

	replicator := core_replication.NewCDCReplicator(
		context.Background(),
		rm.dbSvc,
		rm.cfg.KAFKA_BROKERS,
		&core_replication.ReplicationConfig{
			ConsumerGroup:   consumerGroup,
			SourceDBName:    string(sourceDBName),
			TableMappings:   tableMappings,
			ValidClients:    clients,
			ReplicationFlow: replicationFlow,
		},
	)

	rm.replicators = append(rm.replicators, replicator)

	go func(r api.CDCReplicatorInterface) {
		slog.Info("Starting replicator service")
		if err := r.Start(); err != nil {
			slog.Error("Error starting replicator", "error", err)
		}
	}(replicator)

	return nil
}

// StopReplication stops all active replicators gracefully.
func (rm *ReplicationManager) StopReplication() {
	var wg sync.WaitGroup
	for _, r := range rm.replicators {
		wg.Add(1)
		go func(r api.CDCReplicatorInterface) {
			defer wg.Done()
			r.Stop()
		}(r)
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		slog.Info("All replicators stopped gracefully")
	case <-time.After(15 * time.Second):
		slog.Warn("Timeout waiting for replicators to stop")
	}
}

// GetClients fetches the list of all clients from enterprise-management-service.
func (rm *ReplicationManager) GetClients() []string {
	wo := client.StartWorkflowOptions{
		TaskQueue: "enterprise-management-service",
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 3,
			NonRetryableErrorTypes: []string{
				"WORKFLOW_PANIC",
			},
			BackoffCoefficient: 2,
			InitialInterval:    time.Second,
			MaximumInterval:    time.Second,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	wr, executionErr := rm.temporalService.Client.ExecuteWorkflow(
		ctx,
		wo,
		constant.GetClientsWorkflow,
	)
	if executionErr != nil {
		slog.Error("Failed to execute GetClientsWorkflow", "error", executionErr)
		return []string{}
	}

	var clients []model.Client
	if workflowErr := wr.Get(ctx, &clients); workflowErr != nil {
		slog.Error("Failed to get result from GetClientsWorkflow", "error", workflowErr)
		return []string{}
	}

	result := make([]string, 0, len(clients))
	for _, c := range clients {
		result = append(result, c.ClientID)
	}

	return result
}
