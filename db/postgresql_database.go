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
	sshClients    = make(map[string]*ssh.Client)
)

type PostgreSQLServiceOptions struct {
	Host     string
	Port     int
	User     string
	Password string
	SSLMode  string

	SSHHost     string
	SSHPort     int
	SSHUser     string
	SSHPassword string
}

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
func ConnectPostgres(dbName DBName, options ...PostgreSQLServiceOptions) *pgxpool.Pool {
	key := string(dbName)

	// 1. Check for an existing, healthy pool.
	if pool, ok := pools[key]; ok && pool != nil {
		if err := pool.Ping(context.Background()); err == nil {
			return pool // Pool is healthy, reuse it.
		}
		// Pool is unhealthy. Close it and its associated SSH client before creating a new one.
		log.Printf("Unhealthy pool for '%s' detected, closing old resources.", key)
		pool.Close()
		if client, ok := sshClients[key]; ok && client != nil {
			client.Close()
		}
		delete(pools, key)
		delete(sshClients, key)
	}

	// 2. Gather all configuration details first.
	var (
		postgresHost     string
		postgresUsername string
		postgresPassword string
		postgresSslMode  string
		sshHost          string
		sshPort          string = "22"
		sshUser          string
		sshPassword      string
	)

	if len(options) > 0 {
		opt := options[0]
		postgresHost = fmt.Sprintf("%s:%d", opt.Host, opt.Port)
		postgresUsername = opt.User
		postgresPassword = opt.Password
		postgresSslMode = opt.SSLMode

		sshHost = opt.SSHHost
		if opt.SSHPort != 0 {
			sshPort = fmt.Sprintf("%d", opt.SSHPort)
		}
		sshUser = opt.SSHUser
		sshPassword = opt.SSHPassword
	} else {
		host := os.Getenv("POSTGRES_HOST")
		if host == "" {
			host = "localhost"
		}
		port := os.Getenv("POSTGRES_PORT")
		if port == "" {
			port = "5432"
		}
		postgresHost = fmt.Sprintf("%s:%s", host, port)

		postgresUsername = os.Getenv("POSTGRES_USERNAME")
		postgresPassword = os.Getenv("POSTGRES_PASSWORD")
		postgresSslMode = os.Getenv("POSTGRES_SSLMODE")

		sshHost = os.Getenv("SSH_HOST")
		sshUser = os.Getenv("SSH_USER")
		sshPassword = os.Getenv("SSH_PASSWORD")
	}

	if postgresSslMode == "" {
		postgresSslMode = "disable"
	}
	if postgresHost == "" || postgresHost == ":0" {
		postgresHost = "localhost:5432"
	}

	postgresUri := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=%s",
		postgresUsername, postgresPassword, postgresHost, dbName, postgresSslMode,
	)

	config, err := pgxpool.ParseConfig(postgresUri)
	if err != nil {
		log.Fatalf("Unable to parse PostgreSQL URI: %v", err)
	}

	var sshClient *ssh.Client // Declare here to be accessible for storage later.

	// 3. Configure the SSH tunnel (if needed) and apply it to the config object.
	if sshHost != "" {

		sshConfig := &ssh.ClientConfig{
			User:            sshUser,
			Auth:            []ssh.AuthMethod{ssh.Password(sshPassword)},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Timeout:         5 * time.Second,
		}

		splittedHost := strings.Split(postgresHost, ":")
		sshClient, err = ssh.Dial("tcp", net.JoinHostPort(sshHost, sshPort), sshConfig)
		if err != nil {
			log.Fatalf("Failed to dial SSH server: %v", err)
		}
		// DO NOT use defer sshClient.Close() here. The client must live on.
		fmt.Println("✅ SSH connection established successfully!")

		dialer := func(ctx context.Context, network, addr string) (net.Conn, error) {
			return sshClient.Dial("tcp", net.JoinHostPort(splittedHost[0], splittedHost[1]))
		}
		config.ConnConfig.DialFunc = dialer
	}

	// 4. Apply health check settings to the config. This is always a good practice.
	// Old max conn is 4 or runtime.NumCPU() (default by pgx itself). For DB-level tenant this is fine, but for schema-level tenant this is bad (lower concurrency); adjust to 30 as a safe bet.
	config.MaxConns = 30
	// Keep 5 for warm up
	config.MinConns = 5
	config.MaxConnIdleTime = 5 * time.Minute
	config.MaxConnLifetime = 2 * time.Hour
	config.HealthCheckPeriod = 1 * time.Minute

	// 5. Now, create the pool using the fully prepared config.
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		if sshClient != nil {
			sshClient.Close() // Clean up the SSH client if pool creation fails.
		}
		// Handle "database does not exist" error.
		notExisingDB := fmt.Sprintf("database \"%s\" does not exist", dbName)
		if strings.Contains(err.Error(), notExisingDB) {
			// ... your migration logic here ...
			log.Fatalf("Database '%s' does not exist. Please run migrations and restart.", dbName)
		} else {
			log.Fatalf("Unable to connect to PostgreSQL: %v", err)
		}
	}

	// 6. Store the new pool and its SSH client for future use.
	pools[key] = pool
	if sshClient != nil {
		sshClients[key] = sshClient
	}

	log.Printf("Connected to PostgreSQL database: %s\n", dbName)
	return pool
}
