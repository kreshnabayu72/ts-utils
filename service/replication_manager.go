package service

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"sync"
	"time"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/config"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/constant"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/db"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/enterprise_management_service/model"
	file_manager_service "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/file_manager_service/replica"
	hardware_tagging_replica "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/hardware_tagging_service/replica"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/iam_service/replica"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
)

type ReplicatedTable interface {
	GetTableName() string
	GetColumnTypes() map[string]string
	GetColumns() []string
}

// ReplicatedTables - db name with prefix `true` means to use tenant-specific schema, otherwise use public schema
var ReplicatedTables = map[string]map[string][]ReplicatedTable{
	constant.IAM_SERVICE_NAME: { // Listener/Consumer
		fmt.Sprintf("%v.%s", true, db.HardwareTaggingDBName): { // Publisher
			&replica.Rfid{},
		},
	},
	constant.HARDWARE_TAGGING_SERVICE_NAME: { // Listener/Consumer
		fmt.Sprintf("%v.%s", false, db.HardwareTaggingDBName): { // Publisher
			&hardware_tagging_replica.DeviceName{},
			&hardware_tagging_replica.HandheldStock{},
			&hardware_tagging_replica.Sku{},
			&hardware_tagging_replica.TabletStock{},
			&hardware_tagging_replica.SmartFixedStock{},
		},
		fmt.Sprintf("%v.%s", true, db.IAMDBName): { // Publisher
			&hardware_tagging_replica.FixedAssetGroup{},
			&hardware_tagging_replica.SupplyAssetGroup{},
			&hardware_tagging_replica.User{},
		},
	},
	constant.ENTERPRISE_MANAGEMENT_SERVICE_NAME: nil,
	constant.MASTER_ASSET_REGISTRY_SERVICE_NAME: nil,
	constant.ANALYTICS_REPORTING_SERVICE_NAME:   nil,
	constant.APPROVAL_SERVICE_NAME:              nil,
	constant.FILE_MANAGER_SERVICE_NAME: { // Listener/Consumer
		fmt.Sprintf("%v.%s", true, db.MasterAssetRegistryDBName): { // Publisher
			&file_manager_service.AssetCustomField{},
			&file_manager_service.Asset{},
			&file_manager_service.CustomField{},
		},
		fmt.Sprintf("%v.%s", true, db.AssetTransactionDBName): { // Publisher
			&file_manager_service.FixedAssetDisposalTransaction{},
		},
		fmt.Sprintf("%v.%s", true, db.IAMDBName): { // Publisher
			&file_manager_service.FixedAssetGroup{},
			&file_manager_service.User{},
		},
	},
	constant.IMPORT_SERVICE_NAME:               nil,
	constant.LANGUAGE_TRANSLATION_SERVICE_NAME: nil,
	constant.NOTIFICATION_SERVICE_NAME:         nil,
	constant.PLANNING_STRATEGY_SERVICE_NAME:    nil,
	constant.RESOURCE_ALLOCATION_SERVICE_NAME:  nil,
	constant.ASSET_ADDON_SERVICE_NAME:          nil,
	constant.ASSET_TRANSACTION_SERVICE_NAME:    {
		// db.FixedAssetGroupTableName,
		// db.FixedAssetUserTransactionRoleTableName,
		// db.UserTableName,
	},
}

// ReplicationManager manages CDC replicators for multiple tenants.
type ReplicationManager struct {
	replicators     []CDCReplicatorInterface
	cfg             config.ConfigOptions
	dbSvc           PostgreSQLServicer
	temporalService *TemporalService
}

// NewReplicationManager creates a new ReplicationManager.
// dbFactory is a function that returns a PostgreSQLServicer for a given tenant.
func NewReplicationManager(
	cfg config.ConfigOptions,
	temporalService *TemporalService,
	dbSvc PostgreSQLServicer,
) *ReplicationManager {
	return &ReplicationManager{
		replicators:     make([]CDCReplicatorInterface, 0),
		cfg:             cfg,
		dbSvc:           dbSvc,
		temporalService: temporalService,
	}
}

// StartReplication starts replication for the given tenants.
// targetServiceName: The name of the current service (e.g., "iam-service"). Used to find replicated tables.
// sourceDBSuffix: The suffix of the source database (e.g., "tagsamurai_hardware_tagging").
// isTenantDependent: Whether the source database is tenant-dependent.
func (rm *ReplicationManager) StartReplication(
	targetServiceName string,
	sourceDBSuffix db.DBName,
	isTenantDependent bool,
) error {
	slog.Debug("Entering StartReplication", "target_service", targetServiceName)
	clients := rm.GetClients()
	slog.Debug("GetClients returned", "clients_count", len(clients))

	if len(clients) == 0 {
		slog.Warn("No clients found", "service_name", targetServiceName)
		return nil
	}

	replicatedTables := ReplicatedTables[targetServiceName][fmt.Sprintf("%v.%s", isTenantDependent, sourceDBSuffix)]
	if len(replicatedTables) == 0 {
		slog.Warn("No replicated tables found", "service_name", targetServiceName)
		return nil
	}

	tableMappings := make(map[string]TargetTableMapping, len(replicatedTables))
	for _, table := range replicatedTables {
		tableMappings[table.GetTableName()] = TargetTableMapping{
			PrimaryKey:      "id",
			Columns:         table.GetColumns(),
			TargetTableName: table.GetTableName(),
			ColumnTypes:     table.GetColumnTypes(),
		}
	}

	// Replace hyphens with underscores for consumer group name consistency
	sanitizedTarget := strings.ReplaceAll(targetServiceName, "-", "_")
	sanitizedSource := strings.ReplaceAll(string(sourceDBSuffix), "-", "_")
	consumerGroup := fmt.Sprintf("%s_%s_replicator", sanitizedTarget, sanitizedSource)

	replicator := NewCDCReplicator(
		context.Background(),
		rm.dbSvc,
		rm.cfg.KAFKA_BROKERS,
		&ReplicationConfig{
			ConsumerGroup: consumerGroup,
			SourceDBName:  string(sourceDBSuffix),
			TableMappings: tableMappings,
		},
	)

	rm.replicators = append(rm.replicators, replicator)

	go func(r CDCReplicatorInterface) {
		slog.Info("Starting replicator service")
		if err := r.Start(); err != nil {
			slog.Error("Error starting replicator", "error", err)
		}
	}(replicator)

	return nil
}

// StopReplication stops all replicators gracefully.
func (rm *ReplicationManager) StopReplication() {
	var wg sync.WaitGroup
	for _, r := range rm.replicators {
		wg.Add(1)
		go func(r CDCReplicatorInterface) {
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
	case <-time.After(5 * time.Second):
		slog.Warn("Timeout waiting for replicators to stop")
	}
}

func (rm *ReplicationManager) GetClients() []string {
	wo := client.StartWorkflowOptions{
		TaskQueue: "enterprise-management-service",
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 1,
			NonRetryableErrorTypes: []string{
				"WORKFLOW_PANIC",
			},
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
