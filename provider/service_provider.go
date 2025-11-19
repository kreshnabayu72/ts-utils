package provider

import (
	"github.com/stretchr/testify/mock"

	"gitea.qwertysystem.net/BETS/ts-utils/v2/db"
	"gitea.qwertysystem.net/BETS/ts-utils/v2/service"
)

type ServiceProvider interface {
	MakeService(dbName db.DBName, collection string) service.Service
	MakePostgreSqlService(dbName db.DBName) service.PostgreSqlService
}

type ServiceProviderImpl struct{}

func (sf *ServiceProviderImpl) MakeService(
	dbName db.DBName,
	collection string,
) service.Service {
	return service.MakeService(dbName, collection)
}

func (sf *ServiceProviderImpl) MakePostgreSqlService(
	dbName db.DBName,
) service.PostgreSqlService {
	return service.MakePostgreSqlService(dbName)
}

type MockServiceProvider struct {
	mock.Mock
}

func (m *MockServiceProvider) MakeService(
	dbName db.DBName,
	collection string,
) service.Service {
	args := m.Called(dbName, collection)
	return args.Get(0).(service.Service)
}

func (m *MockServiceProvider) MakePostgreSqlService(
	dbName db.DBName,
) service.PostgreSqlService {
	args := m.Called(dbName)
	return args.Get(0).(service.PostgreSqlService)
}
