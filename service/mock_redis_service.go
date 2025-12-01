package service

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/mock"
)

// MockRedisService implements RedisServicer
//
// Used to mock the RedisService implementation.
type MockRedisService struct {
	mock.Mock
}

func (m *MockRedisService) Init() {
	// Pass
}

func (m *MockRedisService) GetClient() *redis.Client {
	args := m.Called()
	return args.Get(0).(*redis.Client)
}

func (m *MockRedisService) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	args := m.Called(ctx, key, value, expiration)
	return args.Error(0)
}

func (m *MockRedisService) Get(ctx context.Context, key string) (string, error) {
	args := m.Called(ctx, key)
	return args.String(0), args.Error(1)
}

func (m *MockRedisService) Del(ctx context.Context, keys ...string) error {
	args := m.Called(ctx, keys)
	return args.Error(0)
}

func (m *MockRedisService) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	args := m.Called(ctx, key)
	return args.Get(0).(map[string]string), args.Error(1)
}

func (m *MockRedisService) Close() {
	m.Called()
}
