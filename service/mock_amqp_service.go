package service

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/mock"
)

type MockAMQPService struct {
	mock.Mock
}

func (m *MockAMQPService) Publish(
	ctx context.Context,
	exchange ExchangeDeclareArgs,
	pub PublishArgs,
) error {
	args := m.Mock.Called(ctx, exchange, pub)
	return args.Error(0)
}

func (m *MockAMQPService) Subscribe(
	exchange ExchangeDeclareArgs,
	queue QueueDeclareArgs,
	bind QueueBindArgs,
	consume ConsumeArgs,
	handleFunc func(delivery <-chan amqp.Delivery),
) error {
	args := m.Mock.Called(exchange, queue, bind, consume, handleFunc)
	return args.Error(0)
}

func (m *MockAMQPService) SimplePublish(
	ctx context.Context,
	exchange string,
	key string,
	msg amqp.Publishing,
) error {
	args := m.Mock.Called(ctx, exchange, key, msg)
	return args.Error(0)
}

func (m *MockAMQPService) SimpleSubscribe(
	exchange string,
	queue string,
	key string,
	consumer string,
	handleFunc func(delivery <-chan amqp.Delivery),
) error {
	args := m.Mock.Called(exchange, queue, key, consumer, handleFunc)
	return args.Error(0)
}
