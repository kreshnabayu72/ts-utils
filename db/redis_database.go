package db

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisServiceOptions struct {
	Host     string
	Port     string
	Password string
	DB       int
	Timeout  time.Duration
}

func ConnectRedis(options ...RedisServiceOptions) *redis.Client {
	slog.Debug("Connecting to Redis...")

	opts := RedisServiceOptions{
		Host:     "localhost",
		Port:     "6379",
		Password: "",
		DB:       0,
		Timeout:  5 * time.Second,
	}

	if len(options) > 0 {
		slog.Debug("Redis configuration found.")
		opts = options[0]
	}

	addr := fmt.Sprintf("%s:%s", opts.Host, opts.Port)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: opts.Password,
		DB:       opts.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), opts.Timeout)
	defer cancel()
	if err := client.Ping(ctx).Err(); err != nil {
		slog.Error(
			"Failed to connect to Redis.",
			"error", err,
		)

		return nil
	}

	slog.Debug("Connected to Redis.")
	return client
}
