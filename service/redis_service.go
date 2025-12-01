package service

import (
	"context"
	"log/slog"
	"time"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/db"
	"github.com/go-redis/redis/v8"
)

// RedisServicer provides an interface for implementing RedisService
type RedisServicer interface {
	Init()
	GetClient() *redis.Client
	Set(ctx context.Context, key string, value any, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, keys ...string) error
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	Expire(ctx context.Context, key string, expiration time.Duration) error
	Close()
}

// RedisService implements RedisServicer
type RedisService struct {
	Client *redis.Client
}

// MakeRedisService creates a new RedisServicer instance.
func MakeRedisService(options ...db.RedisServiceOptions) *RedisService {
	client := db.ConnectRedis(options...)
	return &RedisService{
		Client: client,
	}
}

func (s *RedisService) Init() {
	// Pass
}

// GetClient returns the underlying redis client.
func (s *RedisService) GetClient() *redis.Client {
	return s.Client
}

// Set sets a key-value pair in Redis with an expiration time.
func (s *RedisService) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	slog.Debug("Redis Set", "key", key)
	return s.Client.Set(ctx, key, value, expiration).Err()
}

// Get retrieves a value from Redis by key.
func (s *RedisService) Get(ctx context.Context, key string) (string, error) {
	slog.Debug("Redis Get", "key", key)
	return s.Client.Get(ctx, key).Result()
}

// Del deletes one or more keys from Redis.
func (s *RedisService) Del(ctx context.Context, keys ...string) error {
	slog.Debug("Redis Del", "keys", keys)
	return s.Client.Del(ctx, keys...).Err()
}

// HGetAll retrieves all fields and values of the hash stored at key.
func (s *RedisService) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	slog.Debug("Redis HGetAll", "key", key)
	return s.Client.HGetAll(ctx, key).Result()
}

// Expire sets a timeout on key.
func (s *RedisService) Expire(ctx context.Context, key string, expiration time.Duration) error {
	slog.Debug("Redis Expire", "key", key, "expiration", expiration)
	return s.Client.Expire(ctx, key, expiration).Err()
}

func (s *RedisService) Close() {
	slog.Debug("Closing Redis service...")
	if s.Client != nil {
		s.Client.Close()
	}
	slog.Debug("Redis service closed.")
}
