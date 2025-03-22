package provider

import (
	"github.com/stretchr/testify/mock"

	"gitea.qwertysystem.net/BETS/ts-utils/service"
)

type AMQPServiceProvider interface {
	MakeAMQPService() service.AMQPService
}

type AMQPServiceProviderImpl struct{}

func (asf *AMQPServiceProviderImpl) MakeAMQPService() service.AMQPService {
	return service.MakeAMQPService()
}

type MockAMQPServiceProvider struct {
	mock.Mock
}

func (m *MockServiceProvider) MakeAMQPService() service.AMQPService {
	args := m.Called()
	return args.Get(0).(service.AMQPService)
}
