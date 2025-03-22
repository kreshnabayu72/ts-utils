package provider

import (
	"time"

	"github.com/stretchr/testify/mock"

	"gitea.qwertysystem.net/BETS/ts-utils/service"
)

type AMQPServiceProvider interface {
	MakeAMQPService(
		reconnectDelay time.Duration,
		reconnectAttempts int,
		keepAlive bool,
	) service.AMQPService
}

type AMQPServiceProviderImpl struct{}

func (asf *AMQPServiceProviderImpl) MakeAMQPService(
	reconnectDelay time.Duration,
	reconnectAttempts int,
	keepAlive bool,
) service.AMQPService {
	return service.MakeAMQPService(reconnectDelay, reconnectAttempts, keepAlive)
}

type MockAMQPServiceProvider struct {
	mock.Mock
}

func (m *MockServiceProvider) MakeAMQPService(
	reconnectDelay time.Duration,
	reconnectAttempts int,
	keepAlive bool,
) service.AMQPService {
	args := m.Called()
	return args.Get(0).(service.AMQPService)
}
