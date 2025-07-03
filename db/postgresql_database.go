package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/bwmarrin/snowflake"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	pools         = make(map[string]*pgxpool.Pool)
	dbOnce        sync.Once
	snowflakeOnce sync.Once
	Node          *snowflake.Node
)

func InitSnowflake() {
	snowflakeOnce.Do(func() {
		var err error
		Node, err = snowflake.NewNode(1) // 1 = node ID (range: 0–1023)
		if err != nil {
			log.Fatalf("failed to initialize snowflake node: %v", err)
		}
		fmt.Println("Snowflake node initialized")
	})
}

// ConnectPostgres initializes the PostgreSQL connection pool once
func ConnectPostgres(dbName DBName) *pgxpool.Pool {
	InitSnowflake()

	dbOnce.Do(func() {
		postgresHost := os.Getenv("POSTGRES_HOST")
		postgresUsername := os.Getenv("POSTGRES_USERNAME")
		postgresPassword := os.Getenv("POSTGRES_PASSWORD")
		postgresSslMode := os.Getenv("POSTGRES_SSLMODE")

		if postgresSslMode == "" {
			postgresSslMode = "disable"
		}

		if postgresHost == "" {
			postgresHost = "localhost:5432"
		}

		postgresUri := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", postgresUsername, postgresPassword, postgresHost, dbName, postgresSslMode)

		config, err := pgxpool.ParseConfig(postgresUri)
		if err != nil {
			log.Fatalf("Unable to parse PostgreSQL URI: %v", err)
		}

		pool, err := pgxpool.NewWithConfig(context.Background(), config)
		if err != nil {
			log.Fatalf("Unable to connect to PostgreSQL: %v", err)
		}

		err = pool.Ping(context.Background())
		if err != nil {
			log.Fatalf("Unable to ping PostgreSQL: %v", err)
		}

		fmt.Println("PostgreSQL pool initialized!")

		pools[string(dbName)] = pool
		log.Printf("Connected to PostgreSQL database: %s\n", dbName)
	})

	return pools[string(dbName)]
}
