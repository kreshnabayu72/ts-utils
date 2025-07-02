package service

import (
	"context"
	"errors"
	"log"

	"gitea.qwertysystem.net/BETS/ts-utils/db"
	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// Base Service PostgreSQL
type PostgreSqlService interface {
	GetPool() PgxPoolInterface
	SelectMany(v any, ctx context.Context, queryString string, args ...any) error
	SelectOne(v any, ctx context.Context, queryString string, args ...any) error
	InsertOne(ctx context.Context, queryString string, args ...any) (interface{}, error)
	UpdateOne(ctx context.Context, queryString string, args ...any) (interface{}, error)
	DeleteOne(ctx context.Context, queryString string, args ...any) (interface{}, error)
	SetTransaction(tx pgx.Tx)
}

type PgxPoolInterface interface {
	// Connection control
	Begin(ctx context.Context) (pgx.Tx, error)

	// Query execution
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)

	Close()
}

type BasePostgreSqlService struct {
	Pool        PgxPoolInterface
	Transaction pgx.Tx
	TableName   string
}

func MakePostgreSqlService(dbName db.DBName, tableName string) *BasePostgreSqlService {
	pool := db.ConnectPostgres(dbName)

	return &BasePostgreSqlService{Pool: pool, TableName: tableName}
}

// Used for getting pool
func (s *BasePostgreSqlService) GetPool() PgxPoolInterface {
	return s.Pool
}

// Select array of data, equivalent of Find() or Aggregate()
func (s *BasePostgreSqlService) SelectMany(v any, ctx context.Context, queryString string, args ...any) error {
	var rows pgx.Rows
	var err error

	if s.Transaction != nil {
		rows, err = s.Transaction.Query(ctx, queryString, args...)
	} else {
		rows, err = s.Pool.Query(ctx, queryString, args...)
	}

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
	var rows pgx.Rows
	var err error

	if s.Transaction != nil {
		rows, err = s.Transaction.Query(ctx, queryString, args...)
	} else {
		rows, err = s.Pool.Query(ctx, queryString, args...)
	}

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
	var err error

	if s.Transaction != nil {
		err = s.Transaction.QueryRow(ctx, queryString, args...).Scan(&resultId)
	} else {
		err = s.Pool.QueryRow(ctx, queryString, args...).Scan(&resultId)
	}

	if err != nil {
		return nil, err
	}

	return resultId, nil
}

func (s *BasePostgreSqlService) UpdateOne(ctx context.Context, queryString string, args ...any) (interface{}, error) {
	var resultId int
	var err error

	if s.Transaction != nil {
		err = s.Transaction.QueryRow(ctx, queryString, args...).Scan(&resultId)
	} else {
		err = s.Pool.QueryRow(ctx, queryString, args...).Scan(&resultId)
	}

	if err != nil {
		return nil, err
	}

	return resultId, nil
}

func (s *BasePostgreSqlService) DeleteOne(ctx context.Context, queryString string, args ...any) (interface{}, error) {
	var resultId int
	var err error

	if s.Transaction != nil {
		err = s.Transaction.QueryRow(ctx, queryString, args...).Scan(&resultId)
	} else {
		err = s.Pool.QueryRow(ctx, queryString, args...).Scan(&resultId)
	}

	if err != nil {
		return nil, err
	}

	return resultId, nil
}

func (s *BasePostgreSqlService) SetTransaction(tx pgx.Tx) {
	s.Transaction = tx
}

func UseTransactions(ctx context.Context, pool PgxPoolInterface, fn func(tx pgx.Tx) (any, error)) (any, error) {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return nil, errors.New("Can't start transactions: " + err.Error())
	}

	defer func() {
		tx.Rollback(ctx)
	}()

	result, err := fn(tx)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return result, nil
}
