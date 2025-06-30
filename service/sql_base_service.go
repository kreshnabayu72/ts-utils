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
	SelectMany(v any, ctx context.Context, queryString string, args ...any) error
	SelectOne(v any, ctx context.Context, queryString string, args ...any) error
	InsertOne(ctx context.Context, queryString string, args ...any) (interface{}, error)
	UpdateOne(ctx context.Context, queryString string, args ...any) (interface{}, error)
	DeleteOne(ctx context.Context, queryString string, args ...any) (interface{}, error)
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
func (s *BasePostgreSqlService) SelectMany(v any, ctx context.Context, queryString string, args ...any) error {
	rows, err := s.Pool.Query(context.Background(), queryString, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	err = sql_query.ScanRowsArray(v, rows)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// Select array of data, equivalent of FindOne()
func (s *BasePostgreSqlService) SelectOne(v any, ctx context.Context, queryString string, args ...any) error {
	rows, err := s.Pool.Query(context.Background(), queryString, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	err = sql_query.ScanRowObject(v, rows)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *BasePostgreSqlService) InsertOne(ctx context.Context, queryString string, args ...any) (interface{}, error) {
	var resultId int

	err := s.Pool.QueryRow(context.Background(), queryString, args...).Scan(&resultId)
	if err != nil {
		return nil, err
	}

	return resultId, nil
}

func (s *BasePostgreSqlService) UpdateOne(ctx context.Context, queryString string, args ...any) (interface{}, error) {
	var resultId int

	err := s.Pool.QueryRow(context.Background(), queryString, args...).Scan(&resultId)
	if err != nil {
		return nil, err
	}

	return resultId, nil
}

func (s *BasePostgreSqlService) DeleteOne(ctx context.Context, queryString string, args ...any) (interface{}, error) {
	var resultId int

	err := s.Pool.QueryRow(context.Background(), queryString, args...).Scan(&resultId)
	if err != nil {
		return nil, err
	}

	return resultId, nil
}
