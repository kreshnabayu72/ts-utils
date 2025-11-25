package provider

import (
	"gitea.qwertysystem.net/BETS/ts-utils/v2/db"
	"gitea.qwertysystem.net/BETS/ts-utils/v2/service"
	"github.com/stretchr/testify/mock"
)

// MockDBProvider implements DBProviderer
//
// Used to mock the DBProvider implementation.
type MockDBProvider struct {
	mock.Mock
}

func (m *MockDBProvider) Init() {
	// Pass
}

func (m *MockDBProvider) MakeMongoDBService(dbName db.DBName, collection string) service.MongoDBServicer {
	args := m.Called(dbName, collection)
	return args.Get(0).(service.MongoDBServicer)
}

func (m *MockDBProvider) MakePostgreSQLService(dbName db.DBName, options ...db.PostgreSQLServiceOptions) service.PostgreSQLServicer {
	args := m.Called(dbName, options)
	return args.Get(0).(service.PostgreSQLServicer)
}

func (m *MockDBProvider) Close() {
	// Pass
}
