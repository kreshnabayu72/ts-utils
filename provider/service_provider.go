package provider

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/mock"

	"gitea.qwertysystem.net/BETS/ts-utils/db"
	"gitea.qwertysystem.net/BETS/ts-utils/service"
)

type ServiceProvider interface {
	MakeService(dbName db.DBName, collection string) service.Service
	MakePostgreSqlService(pool *pgxpool.Pool, tableName string) service.PostgreSqlService
}

type ServiceProviderImpl struct{}

func (sf *ServiceProviderImpl) MakeService(
	dbName db.DBName,
	collection string,
) service.Service {
	return service.MakeService(dbName, collection)
}

func (sf *ServiceProviderImpl) MakePostgreSqlService(
	pool *pgxpool.Pool,
	tableName string,
) service.PostgreSqlService {
	return service.MakePostgreSqlService(pool, tableName)
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
	pool *pgxpool.Pool,
	tableName string,
) service.PostgreSqlService {
	args := m.Called(pool, tableName)
	return args.Get(0).(service.PostgreSqlService)
}
