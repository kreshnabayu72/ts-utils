package mqtt

import "github.com/stretchr/testify/mock"

type MockMQTTService struct {
	mock.Mock
}

func (m *MockMQTTService) Connect() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockMQTTService) Subscribe(topic string, handler MQTTMessageHandler) error {
	args := m.Called(topic, handler)
	return args.Error(0)
}

func (m *MockMQTTService) Publish(topic string, message []byte) error {
	args := m.Called(topic, message)
	return args.Error(0)
}

func (m *MockMQTTService) Unsubscribe(topic string) error {
	args := m.Called(topic)
	return args.Error(0)
}
