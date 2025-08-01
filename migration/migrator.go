package migration

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Model interface {
	TableName() string
	Up(builder *MigrationBuilder)
}

type Migrator struct {
	Host     string
	Username string
	Password string
	SSLMode  string
}

func MakeMigrator(host, username, password, sslMode string) *Migrator {
	return &Migrator{
		Host:     host,
		Username: username,
		Password: password,
		SSLMode:  sslMode,
	}
}

func (m *Migrator) IsDBExist(dbName string) bool {
	postgresUri := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=%s",
		m.Username,
		m.Password,
		m.Host,
		dbName,
		m.SSLMode,
	)

	config, err := pgxpool.ParseConfig(postgresUri)
	if err != nil {
		return false
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return false
	}
	defer pool.Close()

	err = pool.Ping(context.Background())
	if err != nil {
		return false
	}

	return true
}

func (m *Migrator) MigrateDB(dbName string) error {
	defaultDB := "postgres"
	postgresUri := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=%s",
		m.Username,
		m.Password,
		m.Host,
		defaultDB,
		m.SSLMode,
	)

	config, err := pgxpool.ParseConfig(postgresUri)
	if err != nil {
		return err
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return err
	}
	defer pool.Close()

	// Create the database if it does not exist
	createDBQuery := fmt.Sprintf("CREATE DATABASE %s", dbName)
	_, err = pool.Exec(context.Background(), createDBQuery)
	if err != nil {
		return fmt.Errorf("failed to create database %s: %w", dbName, err)
	}

	return nil
}

func (m *Migrator) migrateModel(
	pool *pgxpool.Pool,
	allBuilders []MigrationBuilder,
	builderToMigrate MigrationBuilder,
) error {
	// Ensure reference table exists
	for _, fk := range builderToMigrate.ForeignKeys {
		foundBuilder := MigrationBuilder{}
		for _, b := range allBuilders {
			if b.TableName == fk.RefTable {
				foundBuilder = b
				break
			}
		}

		if foundBuilder.TableName != "" {
			m.migrateModel(pool, allBuilders, foundBuilder)
		}
	}

	query := builderToMigrate.Build()
	_, err := pool.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf(
			"failed to execute migration for table %s: %w",
			builderToMigrate.TableName,
			err,
		)
	}

	return nil
}

func (m *Migrator) MigrateModels(dbName string, models ...Model) error {
	postgresUri := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=%s",
		m.Username,
		m.Password,
		m.Host,
		dbName,
		m.SSLMode,
	)

	config, err := pgxpool.ParseConfig(postgresUri)
	if err != nil {
		return err
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return err
	}
	defer pool.Close()

	migrationBuilders := []MigrationBuilder{}
	for _, model := range models {
		migrationBuilder := &MigrationBuilder{
			TableName: model.TableName(),
		}
		model.Up(migrationBuilder)

		migrationBuilders = append(migrationBuilders, *migrationBuilder)
	}

	for _, builder := range migrationBuilders {
		err := m.migrateModel(pool, migrationBuilders, builder)
		if err != nil {
			return fmt.Errorf("migration failed for model %s: %w", builder.TableName, err)
		}
	}

	return nil
}

func (m *Migrator) Exec(dbName, query string) error {
	postgresUri := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=%s",
		m.Username,
		m.Password,
		m.Host,
		dbName,
		m.SSLMode,
	)

	config, err := pgxpool.ParseConfig(postgresUri)
	if err != nil {
		return err
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return err
	}
	defer pool.Close()

	_, err = pool.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}

func (m *Migrator) ExecFile(dbName, filePath string) error {
	query, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	return m.Exec(dbName, string(query))
}
