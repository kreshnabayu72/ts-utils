package service

import (
	"context"
	"log"

	"gitea.qwertysystem.net/BETS/ts-utils/db"
	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Base Service PostgreSQL
type PostgreSqlService interface {
	SelectQuery(v any, ctx context.Context, queryString string, args ...any) error
}

type BasePostgreSqlService struct {
	Pool      *pgxpool.Pool
	TableName string
}

func MakePostgreSqlService(dbName db.DBName, tableName string) *BasePostgreSqlService {
	pool := db.ConnectPostgres(dbName)

	return &BasePostgreSqlService{Pool: pool, TableName: tableName}
}

func (s *BasePostgreSqlService) SelectQuery(v any, ctx context.Context, queryString string, args ...any) error {
	rows, err := s.Pool.Query(context.Background(), queryString, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	err = sql_query.ScanRows(v, rows)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
