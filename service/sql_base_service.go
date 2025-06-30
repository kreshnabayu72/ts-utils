package service

import (
	"context"
	"errors"
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

// Select array of data, equivalent of Find() or Aggregate()
func (s *BasePostgreSqlService) SelectQuery(v any, ctx context.Context, queryString string, args ...any) error {
	operationType := sql_query.GetSQLOperationType(queryString)

	if operationType != sql_query.SQLOperationsTypeEnumFind && operationType != sql_query.SQLOperationsTypeEnumPagination {
		return errors.New("Invalid query string: SelectQuery require 'SELECT' operation without 'LIMIT 1' or Pagination Query")
	}

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

func (s *BasePostgreSqlService) InsertOneQuery(ctx context.Context, queryString string, args ...any) (interface{}, error) {
	var resultId int

	err2 := s.Pool.QueryRow(context.Background(), queryString, args...).Scan(&resultId)
	if err2 != nil {
		return nil, err2
	}

	return resultId, nil
}
