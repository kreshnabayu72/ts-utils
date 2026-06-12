package replication

import (
	"context"
	"testing"
	"time"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/mocks"
	ems_model "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/enterprise_management_service/model"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/replication/registry"
	"gitea.tagsamurai.local/Wangsit-Developer/be-core-utils/config"
	core_replication "gitea.tagsamurai.local/Wangsit-Developer/be-core-utils/replication"
	"gitea.tagsamurai.local/Wangsit-Developer/be-core-utils/replication/api"
	"gitea.tagsamurai.local/Wangsit-Developer/be-core-utils/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// NewReplicationManager Test
func TestNewReplicationManager_ReturnsNonNil(t *testing.T) {
	cfg := config.ConfigOptions{SERVICE_NAME: "test-service"}
	dbSvc := &service.MockPostgreSQLService{}
	mockClient := new(mocks.MockTemporalClient)
	temporalSvc := &service.TemporalService{Client: mockClient}
	rm := NewReplicationManager(cfg, temporalSvc, dbSvc)

	assert.NotNil(t, rm)
}

func TestNewReplicationManager_ReplicatorsInitialized(t *testing.T) {
	cfg := config.ConfigOptions{}
	dbSvc := &service.MockPostgreSQLService{}
	mockClient := new(mocks.MockTemporalClient)
	temporalSvc := &service.TemporalService{Client: mockClient}
	rm := NewReplicationManager(cfg, temporalSvc, dbSvc)

	require.NotNil(t, rm)
	assert.NotNil(t, rm.replicators, "replicators should be initialized, not nil")
	assert.Len(t, rm.replicators, 0, "replicators should start empty")
}

func TestNewReplicationManager_FieldsAssigned(t *testing.T) {
	cfg := config.ConfigOptions{
		SERVICE_NAME: "test-service",
		ENV:          "local",
	}
	dbSvc := &service.MockPostgreSQLService{}

	mockClient := new(mocks.MockTemporalClient)
	temporalSvc := &service.TemporalService{Client: mockClient}

	rm := NewReplicationManager(cfg, temporalSvc, dbSvc)

	require.NotNil(t, rm)
	assert.Equal(t, cfg, rm.cfg)
	assert.Equal(t, dbSvc, rm.dbSvc)
	assert.Same(t, temporalSvc, rm.temporalService)
}

// seedRegistry inserts test tables into the global registry and returns a cleanup function.
func seedRegistry(serviceName, sourceDB string, flow api.ReplicationFlow, tables []api.ReplicatedTable) func() {
	if registry.ReplicatedTables[serviceName] == nil {
		registry.ReplicatedTables[serviceName] = make(map[api.ReplicationKey][]api.ReplicatedTable)
	}
	key := api.ReplicationKey{Flow: flow, SourceDBName: sourceDB}
	registry.ReplicatedTables[serviceName][key] = tables

	return func() {
		delete(registry.ReplicatedTables, serviceName)
	}
}

// StartReplication Test
func TestStartReplication_NoClients(t *testing.T) {
	temporalSvc := mocks.CreateTemporalService(nil, nil, assert.AnError)
	rm := NewReplicationManager(config.ConfigOptions{}, temporalSvc, &service.MockPostgreSQLService{})

	err := rm.StartReplication("any-service", "any-db", api.ReplicationFlowTenantToTenant)

	assert.Nil(t, err)
	assert.Len(t, rm.replicators, 0)
}

func TestStartReplication_NoTables(t *testing.T) {
	temporalSvc := mocks.CreateTemporalService([]ems_model.Client{{ClientID: "c1"}}, nil, nil)
	rm := NewReplicationManager(config.ConfigOptions{}, temporalSvc, &service.MockPostgreSQLService{})

	// "unknown-service" tidak ada di registry — tidak perlu seed
	err := rm.StartReplication("unknown-service", "any-db", api.ReplicationFlowTenantToTenant)

	assert.Nil(t, err)
	assert.Len(t, rm.replicators, 0)
}

func TestStartReplication_Success(t *testing.T) {
	origNewCDC := newCDCReplicator
	defer func() { newCDCReplicator = origNewCDC }()

	var capturedCfg *core_replication.ReplicationConfig
	mockReplicator := new(mocks.MockCDCReplicator)
	startCalled := make(chan struct{})
	mockReplicator.On("Start").Return(nil).Run(func(_ mock.Arguments) {
		close(startCalled)
	})
	newCDCReplicator = func(
		_ context.Context, _ service.PostgreSQLServicer, _ string,
		cfg *core_replication.ReplicationConfig, _ ...config.ConfigOptions,
	) api.CDCReplicatorInterface {
		capturedCfg = cfg
		return mockReplicator
	}

	table := &mocks.MockReplicatedTable{
		Name:     "users",
		Cols:     []string{"id", "email"},
		ColTypes: map[string]string{"id": "BIGINT", "email": "VARCHAR"},
	}
	cleanup := seedRegistry("my-service", "source-db", api.ReplicationFlowTenantToTenant, []api.ReplicatedTable{table})
	defer cleanup()

	temporalSvc := mocks.CreateTemporalService(
		[]ems_model.Client{{ClientID: "c1"}, {ClientID: "c2"}}, nil, nil,
	)
	rm := NewReplicationManager(config.ConfigOptions{KAFKA_BROKERS: "broker:9092"}, temporalSvc, &service.MockPostgreSQLService{})

	err := rm.StartReplication("my-service", "source-db", api.ReplicationFlowTenantToTenant)

	require.Nil(t, err)
	require.Len(t, rm.replicators, 1)

	// Goroutine launched
	select {
	case <-startCalled:
	case <-time.After(time.Second):
		t.Fatal("goroutine not launched within 1s")
	}

	require.NotNil(t, capturedCfg)

	// Consumer group: hyphens → underscores
	assert.Equal(t, "my_service_source_db_replicator", capturedCfg.ConsumerGroup)
	// SourceDBName
	assert.Equal(t, "source-db", capturedCfg.SourceDBName)
	// ValidClients
	assert.Equal(t, []string{"c1", "c2"}, capturedCfg.ValidClients)
	// ReplicationFlow
	assert.Equal(t, api.ReplicationFlowTenantToTenant, capturedCfg.ReplicationFlow)
	// Table mappings
	require.Len(t, capturedCfg.TableMappings, 1)
	m := capturedCfg.TableMappings["users"]
	assert.Equal(t, "users", m.TargetTableName)
	assert.Equal(t, "id", m.PrimaryKey)
	assert.Equal(t, table.Cols, m.Columns)
	assert.Equal(t, table.ColTypes, m.ColumnTypes)
}
