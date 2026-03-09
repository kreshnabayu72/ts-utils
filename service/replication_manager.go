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

// ReplicatedTable defines the contract for a table that can be replicated via CDC.
// Each struct that implements this interface represents one database table
// and describes its structure so the replicator knows which columns to read and write.
//
// These methods are used in two places:
//  1. Here in the replicator - to build table mappings for CDC consumption (column selection, upsert queries)
//  2. In ts-migrations - to validate that migration schemas match the expected replica table structure.
//     The validator checks that columns and types defined here are consistent with the actual SQL migrations,
//     ensuring schema integrity between the source and target databases.
type ReplicatedTable interface {
	GetTableName() string              // The SQL table name (e.g., "users", "rfids")
	GetColumnTypes() map[string]string // Maps column name -> SQL type (e.g., "id" -> "BIGINT")
	GetColumns() []string              // Ordered list of column names to replicate
}

// ReplicatedTables is the central registry that declares which tables each service
// needs to replicate (consume) from other services (publishers).
//
// Structure:  map[ConsumerServiceName] -> map["ReplicationFlow.sourceDBName"] -> []ReplicatedTable
//
// The map key format "Tenant to Tenant.tagsamurai_iam" or "Non-Tenant to Tenant.tagsamurai_hardware_tagging" encodes:
//   - ReplicationFlow: The configured data routing flow for this replication
//   - sourceDBName: The source database name (e.g., "tagsamurai_iam")
//
// Example: hardware-tagging-service needs to replicate:
//   - From itself in schema public (ReplicationFlowNonTenantToTenant): device_names, handheld_stocks, skus, etc.
//   - From IAM in schema tenant (ReplicationFlowTenantToTenant): users, fixed_asset_groups, supply_asset_groups
//
// How to add a new replicated table:
//  1. Create a struct in the appropriate replica package (e.g., module/iam_service/replica/)
//  2. Implement the ReplicatedTable interface (GetTableName, GetColumnTypes, GetColumns)
//  3. Add it to the correct consumer -> source entry in this map
//  4. Ensure the target database has the matching table in the relevant target schemas
//     (tenant or public) via migrations. This is done by running the migration in the target database.
var ReplicatedTables = map[string]map[string][]ReplicatedTable{
	constant.IAM_SERVICE_NAME: { // Listener/Consumer
		fmt.Sprintf("%v.%s", ReplicationFlowTenantToTenant, db.HardwareTaggingDBName): { // Publisher
			&replica.Rfid{},
		},
	},
	constant.HARDWARE_TAGGING_SERVICE_NAME: { // Listener/Consumer
		fmt.Sprintf("%v.%s", ReplicationFlowNonTenantToTenant, db.HardwareTaggingDBName): { // Publisher
			&hardware_tagging_replica.DeviceName{},
			&hardware_tagging_replica.HandheldStock{},
			&hardware_tagging_replica.Sku{},
			&hardware_tagging_replica.TabletStock{},
			&hardware_tagging_replica.SmartFixedStock{},
		},
		fmt.Sprintf("%v.%s", ReplicationFlowTenantToTenant, db.IAMDBName): { // Publisher
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
		fmt.Sprintf("%v.%s", ReplicationFlowTenantToTenant, db.MasterAssetRegistryDBName): { // Publisher
			&file_manager_service.AssetCustomField{},
			&file_manager_service.Asset{},
			&file_manager_service.CustomField{},
		},
		fmt.Sprintf("%v.%s", ReplicationFlowTenantToTenant, db.AssetTransactionDBName): { // Publisher
			&file_manager_service.FixedAssetDisposalTransaction{},
		},
		fmt.Sprintf("%v.%s", ReplicationFlowTenantToTenant, db.IAMDBName): { // Publisher
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

// ReplicationManager orchestrates CDC (Change Data Capture) replication for a service.
//
// It is the top-level coordinator that:
//  1. Looks up which tables the current service needs to replicate (from ReplicatedTables)
//  2. Creates CDCReplicator instances (one per source database) to consume Kafka messages
//  3. Manages the lifecycle of all replicators (start/stop)
//
// Usage in a service's main.go:
//
//	replicationManager := service.NewReplicationManager(cfgOptions, temporalService, dbProvider.MakePostgreSQLService(db.MyDBName))
//
//	// Start consuming CDC events from IAM service (e.g., ReplicationFlowTenantToTenant)
//	replicationManager.StartReplication(constant.MY_SERVICE_NAME, db.IAMDBName, service.ReplicationFlowTenantToTenant)
//
//	// Start consuming CDC events from own database but public to tenant (e.g., ReplicationFlowNonTenantToTenant)
//	replicationManager.StartReplication(constant.MY_SERVICE_NAME, db.MyDBName, service.ReplicationFlowNonTenantToTenant)
//
//	// On shutdown:
//	replicationManager.StopReplication()
type ReplicationManager struct {
	replicators     []CDCReplicatorInterface // All active replicators managed by this manager
	cfg             config.ConfigOptions     // Application config (contains KAFKA_BROKERS, etc.)
	dbSvc           PostgreSQLServicer       // Target database connection - where replicated data is written to
	temporalService *TemporalService         // Used to fetch the list of tenant client IDs via Temporal workflow
}

// NewReplicationManager creates a new ReplicationManager.
//
// Parameters:
//   - cfg: Application config containing KAFKA_BROKERS and other settings
//   - temporalService: Used to call GetClientsWorkflow to discover tenant client IDs
//   - dbSvc: The PostgreSQL service for the TARGET database (where data will be written)
//
// Important: dbSvc should point to the current service's own database, NOT the source database.
// For example, if hardware-tagging-service replicates from IAM, dbSvc should be
// dbProvider.MakePostgreSQLService(db.HardwareTaggingDBName) - the replicator writes INTO `this` database.
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

type ReplicationFlow string

const (
	ReplicationFlowTenantToTenant       ReplicationFlow = "Tenant to Tenant"
	ReplicationFlowNonTenantToTenant    ReplicationFlow = "Non-Tenant to Tenant"
	ReplicationFlowNonTenantToNonTenant ReplicationFlow = "Non-Tenant to Non-Tenant"
	ReplicationFlowTenantToNonTenant    ReplicationFlow = "Tenant to Non-Tenant"
)

// StartReplication starts CDC replication for a specific source database.
//
// This method does the following:
//  1. Fetches the list of valid clients (via Temporal workflow to enterprise-management-service).
//     This list is passed to the CDCReplicator to validate target schemas when evaluating ReplicationFlow rules.
//  2. Looks up which tables need to be replicated using the ReplicatedTables registry.
//  3. Creates a CDCReplicator that subscribes to Kafka topics for those tables.
//  4. Starts the replicator in a background goroutine.
//
// Parameters:
//   - targetServiceName: The current service's name (e.g., constant.HARDWARE_TAGGING_SERVICE_NAME).
//     Used as the key to look up ReplicatedTables.
//   - sourceDBName: The source database name (e.g., db.IAMDBName = "tagsamurai_iam").
//     This determines which Kafka topics to subscribe to (topic = "sourceDBName.tableName").
//   - replicationFlow: The configured data routing flow for this replication (e.g., ReplicationFlowTenantToTenant).
//     This enum dictates both how tenant validation occurs and where data is safely routed.
//
// Consumer Group Naming:
//
//	The consumer group is named: "{target_service_name}_{source_db_name}_replicator" (with hyphens replaced by underscores).
//	Example: "hardware_tagging_service_tagsamurai_iam_replicator"
//	This ensures each (target_service_name, source_db_name) pair has its own consumer group,
//	so multiple services can independently consume from the same topics.
//
// Call this method once per source database. For example, if hardware-tagging-service
// needs data from both IAM and its own DB:
//
//	replicationManager.StartReplication(constant.HARDWARE_TAGGING_SERVICE_NAME, db.HardwareTaggingDBName, service.ReplicationFlowNonTenantToNonTenant)
//	replicationManager.StartReplication(constant.HARDWARE_TAGGING_SERVICE_NAME, db.IAMDBName, service.ReplicationFlowTenantToTenant)
func (rm *ReplicationManager) StartReplication(
	targetServiceName string,
	sourceDBName db.DBName,
	replicationFlow ReplicationFlow,
) error {
	slog.Debug("Entering StartReplication", "target_service", targetServiceName)
	clients := rm.GetClients()
	slog.Debug("GetClients returned", "clients_count", len(clients))

	if len(clients) == 0 {
		slog.Warn("No clients found", "service_name", targetServiceName)
		return nil
	}

	// Look up which tables this service needs to consume from the given source.
	// The key format is "ReplicationFlow.sourceDBName" (e.g., "Tenant to Tenant.tagsamurai_iam")
	replicatedTables := ReplicatedTables[targetServiceName][fmt.Sprintf("%v.%s", replicationFlow, sourceDBName)]
	if len(replicatedTables) == 0 {
		slog.Warn("No replicated tables found", "service_name", targetServiceName, "flow", replicationFlow)
		return nil
	}

	// Build table mappings from the ReplicatedTable interface implementations.
	// Each mapping tells the replicator how to handle data for a specific table.
	tableMappings := make(map[string]TargetTableMapping, len(replicatedTables))
	for _, table := range replicatedTables {
		tableMappings[table.GetTableName()] = TargetTableMapping{
			PrimaryKey:      "id",
			Columns:         table.GetColumns(),
			TargetTableName: table.GetTableName(),
			ColumnTypes:     table.GetColumnTypes(),
		}
	}

	// Build consumer group name: replace hyphens with underscores for Kafka compatibility.
	sanitizedTarget := strings.ReplaceAll(targetServiceName, "-", "_")
	sanitizedSource := strings.ReplaceAll(string(sourceDBName), "-", "_")
	consumerGroup := fmt.Sprintf("%s_%s_replicator_local", sanitizedTarget, sanitizedSource)

	replicator := NewCDCReplicator(
		context.Background(),
		rm.dbSvc,
		rm.cfg.KAFKA_BROKERS,
		&ReplicationConfig{
			ConsumerGroup:   consumerGroup,
			SourceDBName:    string(sourceDBName),
			TableMappings:   tableMappings,
			ValidClients:    clients,
			ReplicationFlow: replicationFlow,

			// TODO: uncomment if in the future need replication from all tenant schemas to 1 "public" schema.
			// Search global: "enable target public"
			// TargetPublic:  isTargetPublic,
		},
	)

	rm.replicators = append(rm.replicators, replicator)

	// Start the replicator in a background goroutine.
	// r.Start() is a BLOCKING call - it runs an infinite loop reading from Kafka.
	// It only returns when the context is cancelled (via StopReplication) or on fatal error.
	go func(r CDCReplicatorInterface) {
		slog.Info("Starting replicator service")
		if err := r.Start(); err != nil {
			slog.Error("Error starting replicator", "error", err)
		}
	}(replicator)

	return nil
}

// StopReplication stops all active replicators gracefully.
// It cancels the context of each replicator (which stops the Kafka readers),
// then waits up to 5 seconds for all of them to finish.
// This should be called during application shutdown (e.g., on SIGINT/SIGTERM).
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

// GetClients fetches the list of all clients from enterprise-management-service
// via a Temporal workflow call. This is needed at startup to know which clients exist,
// so the replicator can validate destination schemas where required by the ReplicationFlow.
//
// If no clients are returned (e.g., enterprise-management-service is not running),
// replication will be skipped with a warning log.
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
