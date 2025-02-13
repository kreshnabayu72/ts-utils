package provider

import (
	"github.com/stretchr/testify/mock"

	"gitea.qwertysystem.net/BETS/ts-utils/db"
	"gitea.qwertysystem.net/BETS/ts-utils/service"
)

type ServiceProvider interface {
	MakeService(dbName db.DBName, collection string) service.Service
}

type ServiceProviderImpl struct{}

func (sf *ServiceProviderImpl) MakeService(
	dbName db.DBName,
	collection string,
) service.Service {
	return service.MakeService(dbName, collection)
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
