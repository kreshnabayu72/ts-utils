package service

import (
	"context"
	"errors"
	"log"

	"gitea.qwertysystem.net/BETS/ts-utils/db"
	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"
	"gitea.qwertysystem.net/BETS/ts-utils/sql_query/common_builders"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// Base Service PostgreSQL
type PostgreSqlService interface {
	Debug(level ...int)
	GetPool() PgxPoolInterface
	GetTableName() string
	SetTransaction(tx pgx.Tx)

	Count(ctx context.Context, queryString string, args ...any) (int, error)
	CountWithFilter(ctx context.Context, filter map[string]sql_query.SQLCondition) (int, error)

	SelectMany(v any, ctx context.Context, queryString string, args ...any) error
	SelectOne(v any, ctx context.Context, queryString string, args ...any) error

	InsertOne(ctx context.Context, queryString string, args ...any) (interface{}, error)
	InsertOneWithData(ctx context.Context, body interface{}) (interface{}, error)
	InsertMany(ctx context.Context, queryString string, args ...any) (int64, error)
	InsertManyWithData(ctx context.Context, body interface{}) (interface{}, error)

	UpdateOne(ctx context.Context, queryString string, args ...any) (interface{}, error)
	UpdateOneWithData(ctx context.Context, query map[string]sql_query.SQLCondition, body interface{}) (interface{}, error)
	UpdateOneWithDataAndReturning(v any, ctx context.Context, query map[string]sql_query.SQLCondition, body interface{}, returning ...string) error
	UpdateMany(ctx context.Context, queryString string, args ...any) (int64, error)
	UpdateManyWithData(ctx context.Context, query map[string]sql_query.SQLCondition, body interface{}) (int64, error)
	UpdateEachWithData(ctx context.Context, rowIdentifier string, query map[string]sql_query.SQLCondition, body interface{}) (int64, error)

	DeleteOne(ctx context.Context, queryString string, args ...any) (interface{}, error)
	DeleteOneWithFilter(ctx context.Context, filter map[string]sql_query.SQLCondition) (interface{}, error)
	DeleteMany(ctx context.Context, queryString string, args ...any) (int64, error)
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

	debugLevel int
}

func MakePostgreSqlService(dbName db.DBName, tableName string) PostgreSqlService {
	pool := db.ConnectPostgres(dbName)

	return &BasePostgreSqlService{Pool: pool, TableName: tableName}
}

// To show query for debugging in dev mode; 0 = off, 1 = query only, 2 = query + args, other than that is fallback to default 1
// Remove its usage in production
func (s *BasePostgreSqlService) Debug(level ...int) {
	if len(level) > 0 && sql_query.ArrayIncludes([]int{1, 2}, level[0]) {
		s.debugLevel = level[0]
		return
	}

	s.debugLevel = 1
}

// Used for getting pool
func (s *BasePostgreSqlService) GetPool() PgxPoolInterface {
	return s.Pool
}

// Used for getting tablename of this service
func (s *BasePostgreSqlService) GetTableName() string {
	return s.TableName
}

func (s *BasePostgreSqlService) SetTransaction(tx pgx.Tx) {
	s.Transaction = tx
}

func (s *BasePostgreSqlService) Count(ctx context.Context, queryString string, args ...any) (int, error) {
	shouldShowQuery(s.debugLevel, queryString, args...)

	var count int
	var err error

	if s.Transaction != nil {
		err = s.Transaction.QueryRow(ctx, queryString, args...).Scan(&count)
	} else {
		err = s.Pool.QueryRow(ctx, queryString, args...).Scan(&count)
	}

	if err != nil {
		log.Println("Count query failed:", err)
		return 0, err
	}

	return count, nil
}

func (s *BasePostgreSqlService) CountWithFilter(ctx context.Context, filter map[string]sql_query.SQLCondition) (int, error) {
	queryString, args := common_builders.CountBuilder(s.TableName, filter)
	shouldShowQuery(s.debugLevel, queryString, args...)

	return s.Count(ctx, queryString, args...)
}

// Select array of data, equivalent of FindOne()
func (s *BasePostgreSqlService) SelectOne(v any, ctx context.Context, queryString string, args ...any) error {
	shouldShowQuery(s.debugLevel, queryString, args...)

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

// Select array of data, equivalent of Find() or Aggregate()
func (s *BasePostgreSqlService) SelectMany(v any, ctx context.Context, queryString string, args ...any) error {
	shouldShowQuery(s.debugLevel, queryString, args...)

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

func (s *BasePostgreSqlService) InsertOne(ctx context.Context, queryString string, args ...any) (interface{}, error) {
	shouldShowQuery(s.debugLevel, queryString, args...)

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

func (s *BasePostgreSqlService) InsertOneWithData(ctx context.Context, body interface{}) (interface{}, error) {
	queryString, args := common_builders.InsertBuilder(s.TableName, body)
	shouldShowQuery(s.debugLevel, queryString, args...)

	return s.InsertOne(ctx, queryString, args...)
}

func (s *BasePostgreSqlService) InsertMany(ctx context.Context, queryString string, args ...any) (int64, error) {
	shouldShowQuery(s.debugLevel, queryString, args...)

	var commandTag pgconn.CommandTag
	var err error

	if s.Transaction != nil {
		commandTag, err = s.Transaction.Exec(ctx, queryString, args...)
	} else {
		commandTag, err = s.Pool.Exec(ctx, queryString, args...)
	}

	if err != nil {
		return 0, err
	}

	return commandTag.RowsAffected(), nil
}

func (s *BasePostgreSqlService) InsertManyWithData(ctx context.Context, body interface{}) (interface{}, error) {
	queryString, args := common_builders.InsertBuilder(s.TableName, body)
	shouldShowQuery(s.debugLevel, queryString, args...)

	return s.InsertMany(ctx, queryString, args...)
}

func (s *BasePostgreSqlService) UpdateOne(ctx context.Context, queryString string, args ...any) (interface{}, error) {
	shouldShowQuery(s.debugLevel, queryString, args...)

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

func (s *BasePostgreSqlService) UpdateOneWithData(ctx context.Context, query map[string]sql_query.SQLCondition, updateData interface{}) (interface{}, error) {
	queryString, args := common_builders.UpdateBuilder(s.TableName,
		query,
		updateData,
	)
	shouldShowQuery(s.debugLevel, queryString, args...)

	return s.UpdateOne(ctx, queryString, args...)
}

func (s *BasePostgreSqlService) UpdateOneWithDataAndReturning(v any, ctx context.Context, query map[string]sql_query.SQLCondition, updateData interface{}, returningColumn ...string) error {
	queryString, args := common_builders.UpdateBuilder(s.TableName,
		query,
		updateData,
		returningColumn...,
	)
	shouldShowQuery(s.debugLevel, queryString, args...)

	return s.SelectOne(v, ctx, queryString, args...)
}

func (s *BasePostgreSqlService) UpdateMany(ctx context.Context, queryString string, args ...any) (int64, error) {
	shouldShowQuery(s.debugLevel, queryString, args...)

	var commandTag pgconn.CommandTag
	var err error

	if s.Transaction != nil {
		commandTag, err = s.Transaction.Exec(ctx, queryString, args...)
	} else {
		commandTag, err = s.Pool.Exec(ctx, queryString, args...)
	}

	if err != nil {
		return 0, err
	}

	return commandTag.RowsAffected(), nil
}

func (s *BasePostgreSqlService) UpdateManyWithData(ctx context.Context, query map[string]sql_query.SQLCondition, updateData interface{}) (int64, error) {
	queryString, args := common_builders.UpdateBuilder(s.TableName,
		query,
		updateData,
	)
	shouldShowQuery(s.debugLevel, queryString, args...)

	return s.UpdateMany(ctx, queryString, args...)
}

func (s *BasePostgreSqlService) UpdateEachWithData(ctx context.Context, rowIdentifier string, query map[string]sql_query.SQLCondition, updateData interface{}) (int64, error) {
	queryString, args := common_builders.UpdateEachBuilder(s.TableName,
		rowIdentifier,
		query,
		updateData,
	)
	shouldShowQuery(s.debugLevel, queryString, args...)

	return s.UpdateMany(ctx, queryString, args...)
}

func (s *BasePostgreSqlService) DeleteOne(ctx context.Context, queryString string, args ...any) (interface{}, error) {
	shouldShowQuery(s.debugLevel, queryString, args...)

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

func (s *BasePostgreSqlService) DeleteOneWithFilter(ctx context.Context, filter map[string]sql_query.SQLCondition) (interface{}, error) {
	queryString, args := common_builders.DeleteBuilder(s.TableName, filter)
	shouldShowQuery(s.debugLevel, queryString, args...)

	return s.DeleteOne(ctx, queryString, args...)
}

func (s *BasePostgreSqlService) DeleteMany(ctx context.Context, queryString string, args ...any) (int64, error) {
	shouldShowQuery(s.debugLevel, queryString, args...)

	var commandTag pgconn.CommandTag
	var err error

	if s.Transaction != nil {
		commandTag, err = s.Transaction.Exec(ctx, queryString, args...)
	} else {
		commandTag, err = s.Pool.Exec(ctx, queryString, args...)
	}

	if err != nil {
		return 0, err
	}

	return commandTag.RowsAffected(), nil
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

func shouldShowQuery(level int, query string, args ...any) {
	switch level {
	case 1:
		log.Println("query", query)
	case 2:
		log.Println("query and args", query, args)
	}
}
