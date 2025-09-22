package db

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/ssh"
)

var (
	pools         = make(map[string]*pgxpool.Pool)
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
	if pools[string(dbName)] != nil {
		pool := pools[string(dbName)]
		err := pool.Ping(context.Background())
		if err == nil {
			return pool
		}
	}

	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresUsername := os.Getenv("POSTGRES_USERNAME")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresSslMode := os.Getenv("POSTGRES_SSLMODE")

	if postgresSslMode == "" {
		postgresSslMode = "disable"
	}

	if postgresHost == "" {
		postgresHost = "localhost"
		postgresPort = "5432"
	}

	postgresUri := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		postgresUsername,
		postgresPassword,
		postgresHost,
		postgresPort,
		dbName,
		postgresSslMode,
	)

	config, err := pgxpool.ParseConfig(postgresUri) //or (dsn)
	if err != nil {
		log.Fatalf("Unable to parse PostgreSQL URI: %v", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to PostgreSQL: %v", err)
	}

	// ===== SSH Tunnel Configuration =====
	sshHost := os.Getenv("SSH_HOST")
	sshPort := "22"
	sshUser := os.Getenv("SSH_USER")
	sshPassword := os.Getenv("SSH_PASSWORD")

	if sshHost != "" {
		// SSH client configuration using password authentication
		sshConfig := &ssh.ClientConfig{
			User: sshUser,
			Auth: []ssh.AuthMethod{
				ssh.Password(sshPassword),
			},
			// IMPORTANT: In production, use a proper HostKeyCallback.
			// ssh.InsecureIgnoreHostKey() is convenient for development but insecure.
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Timeout:         5 * time.Second,
		}

		// Establish the SSH connection
		sshClient, err := ssh.Dial("tcp", net.JoinHostPort(sshHost, sshPort), sshConfig)
		if err != nil {
			log.Fatalf("Failed to dial SSH server: %v", err)
		}
		defer sshClient.Close()
		fmt.Println("✅ SSH connection established successfully!")
		// This dialer function will be used by pgx to establish connections
		// through the SSH tunnel.
		dialer := func(ctx context.Context, network, addr string) (net.Conn, error) {
			// The addr is ignored because we are tunneling to the dbHost and dbPort defined above.
			return sshClient.Dial("tcp", net.JoinHostPort(postgresHost, postgresPort))
		}

		// Set the custom dialer
		config.ConnConfig.DialFunc = dialer
	}

	err = pool.Ping(context.Background())
	if err != nil {
		notExisingDB := fmt.Sprintf("database \"%s\" does not exist", dbName)
		if strings.Contains(err.Error(), notExisingDB) {
			err := MigrateDB(
				dbName,
				postgresHost,
				postgresUsername,
				postgresPassword,
				postgresSslMode,
			)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatalf("Unable to ping PostgreSQL: %v", err)
		}

	}

	fmt.Println("PostgreSQL pool initialized!")

	pools[string(dbName)] = pool
	log.Printf("Connected to PostgreSQL database: %s\n", dbName)

	return pools[string(dbName)]
}
