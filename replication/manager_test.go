package replication

import (
	"testing"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/mocks"
	"gitea.tagsamurai.local/Wangsit-Developer/be-core-utils/config"
	"gitea.tagsamurai.local/Wangsit-Developer/be-core-utils/service"
	"github.com/stretchr/testify/assert"
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
