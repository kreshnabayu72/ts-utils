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
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/iam_service/replica"
)

type ReplicatedTable interface {
	GetTableName() string
	GetTableNameWithSchema() string
	GetColumnTypes() map[string]string
	GetColumns() []string
}

var ReplicatedTables = map[string][]ReplicatedTable{
	constant.IAM_SERVICE_NAME: {
		&replica.Rfid{},
	},
	constant.HARDWARE_TAGGING_SERVICE_NAME: {
		// db.FixedAssetGroupTableName,
		// db.HandheldStockTableName,
		// db.SkuTableName,
		// db.SmartFixedStockTableName,
		// db.SupplyAssetGroupTableName,
		// db.TabletStockTableName,
		// db.UserTableName,
	},
	constant.ENTERPRISE_MANAGEMENT_SERVICE_NAME: nil,
	constant.MASTER_ASSET_REGISTRY_SERVICE_NAME: nil,
	constant.ANALYTICS_REPORTING_SERVICE_NAME:   nil,
	constant.APPROVAL_SERVICE_NAME:              nil,
	constant.FILE_MANAGER_SERVICE_NAME:          {
		// db.AssetCustomFieldTableName,
		// db.AssetTableName,
		// db.CustomFieldTableName,
		// db.FixedAssetDisposalTransactionTableName,
		// db.FixedAssetGroupTableName,
		// db.UserTableName,
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
	replicators []CDCReplicatorInterface
	cfg         config.ConfigOptions
	dbFactory   func(string) PostgreSQLServicer
}

// NewReplicationManager creates a new ReplicationManager.
// dbFactory is a function that returns a PostgreSQLServicer for a given tenant.
func NewReplicationManager(
	cfg config.ConfigOptions,
	dbFactory func(string) PostgreSQLServicer,
) *ReplicationManager {
	return &ReplicationManager{
		replicators: make([]CDCReplicatorInterface, 0),
		cfg:         cfg,
		dbFactory:   dbFactory,
	}
}

// StartReplication starts replication for the given tenants.
// targetServiceName: The name of the current service (e.g., "iam-service"). Used to find replicated tables.
// sourceDBSuffix: The suffix of the source database (e.g., "tagsamurai_hardware_tagging").
// isTenantDependent: Whether the source database is tenant-dependent.
func (rm *ReplicationManager) StartReplication(
	tenants []string,
	targetServiceName string,
	sourceDBSuffix string,
	isTenantDependent bool,
) error {
	replicatedTables := ReplicatedTables[targetServiceName]
	if len(replicatedTables) == 0 {
		slog.Warn("No replicated tables found", "service_name", targetServiceName)
		return nil
	}

	tableMappings := make(map[string]TargetTableMapping, len(replicatedTables))
	for _, table := range replicatedTables {
		tableMappings[table.GetTableNameWithSchema()] = TargetTableMapping{
			PrimaryKey:      "id",
			Columns:         table.GetColumns(),
			TargetTableName: table.GetTableName(),
			ColumnTypes:     table.GetColumnTypes(),
		}
	}

	for _, tenantID := range tenants {
		// Create PostgreSQL service for the tenant
		pgService := rm.dbFactory(tenantID)

		sourceDBName := sourceDBSuffix
		if isTenantDependent {
			sourceDBName = tenantID + "_" + sourceDBSuffix
		}
		// Replace hyphens with underscores for consumer group name consistency
		consumerGroup := fmt.Sprintf("%s_replicator", strings.ReplaceAll(targetServiceName, "-", "_"))

		replicator := NewCDCReplicator(
			context.Background(),
			pgService,
			rm.cfg.KAFKA_BROKERS,
			&ReplicationConfig{
				TenantID:      tenantID,
				ConsumerGroup: consumerGroup,
				SourceDBName:  sourceDBName,
				TableMappings: tableMappings,
			},
		)

		rm.replicators = append(rm.replicators, replicator)

		go func(id string, r CDCReplicatorInterface) {
			slog.Info("Starting replicator service", "tenant_id", id)
			if err := r.Start(); err != nil {
				slog.Error("Error starting replicator", "tenant_id", id, "error", err)
			}
		}(tenantID, replicator)
	}

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
