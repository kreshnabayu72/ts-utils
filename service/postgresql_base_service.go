package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"gitea.qwertysystem.net/BETS/ts-utils/db"
	"gitea.qwertysystem.net/BETS/ts-utils/dto"
	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"
	"gitea.qwertysystem.net/BETS/ts-utils/sql_query/common_builders"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type ReturningConfig struct {
	Column      []string
	Destination any
}

// Base Service PostgreSQL
type PostgreSqlService interface {
	Debug(level ...int)
	GetPool() PgxPoolInterface
	GetTableName() string
	SetTransaction(tx pgx.Tx)

	Count(ctx context.Context, queryString string, args ...any) (int, error)
	CountWithFilter(ctx context.Context, filter map[string]sql_query.SQLCondition) (int, error)

	SelectOne(v any, ctx context.Context, queryString string, args ...any) error
	SelectMany(v any, ctx context.Context, queryString string, args ...any) error

	InsertOne(ctx context.Context, queryString string, args ...any) (interface{}, error)
	InsertOneWithData(ctx context.Context, body interface{}, returnOption ...ReturningConfig) (interface{}, error)
	InsertMany(ctx context.Context, queryString string, args ...any) (int64, error)
	InsertManyWithData(ctx context.Context, body interface{}, returnOption ...ReturningConfig) (interface{}, error)

	UpdateOne(ctx context.Context, queryString string, args ...any) (interface{}, error)
	UpdateOneWithData(ctx context.Context, query map[string]sql_query.SQLCondition, body interface{}, returnOption ...ReturningConfig) (interface{}, error)
	UpdateMany(ctx context.Context, queryString string, args ...any) (int64, error)
	UpdateManyWithData(ctx context.Context, query map[string]sql_query.SQLCondition, body interface{}, returnOption ...ReturningConfig) (int64, error)
	UpdateEachWithData(ctx context.Context, rowIdentifier string, query map[string]sql_query.SQLCondition, body interface{}) (int64, error)

	SoftDeleteOne(ctx context.Context, filter map[string]sql_query.SQLCondition, returnOption ...ReturningConfig) (interface{}, error)
	SoftDeleteMany(ctx context.Context, filter map[string]sql_query.SQLCondition, returnOption ...ReturningConfig) (int64, error)

	DeleteOne(ctx context.Context, queryString string, args ...any) (interface{}, error)
	DeleteOneWithFilter(ctx context.Context, filter map[string]sql_query.SQLCondition) (interface{}, error)
	DeleteMany(ctx context.Context, queryString string, args ...any) (int64, error)
	DeleteManyWithFilter(ctx context.Context, filter map[string]sql_query.SQLCondition) (int64, error)
}

type PgxPoolInterface interface {
	// Begin starts a new database transaction and returns a pgx.Tx.
	Begin(ctx context.Context) (pgx.Tx, error)

	// Query executes a SQL query and returns pgx.Rows for multiple row results.
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	// QueryRow executes a SQL query that returns a single row.
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	// Exec executes a SQL command (INSERT/UPDATE/DELETE)
	// and returns a pgconn.CommandTag containing rows affected.
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)

	// Close closes the database pool connection.
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

// Debug sets the debug level for printing executed SQL queries.
// Level:
//
//	0 → no logs
//	1 → log queries
//	2 → log queries and arguments
//
// Defaults to level 1 if an invalid level is passed.
// Remove its usage in production
func (s *BasePostgreSqlService) Debug(level ...int) {
	if len(level) > 0 && sql_query.ArrayIncludes([]int{1, 2}, level[0]) {
		s.debugLevel = level[0]
		return
	}

	s.debugLevel = 1
}

// GetPool returns the underlying connection pool (PgxPoolInterface)
// used by this service.
func (s *BasePostgreSqlService) GetPool() PgxPoolInterface {
	return s.Pool
}

// GetTableName returns the table name of the service.
func (s *BasePostgreSqlService) GetTableName() string {
	return s.TableName
}

// SetTransaction assigns an active transaction (pgx.Tx) to the service.
// If a transaction is set, all queries will be executed within it.
func (s *BasePostgreSqlService) SetTransaction(tx pgx.Tx) {
	s.Transaction = tx
}

// Count executes a SELECT COUNT(*) query and returns the number of rows.
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

// CountWithFilter builds a COUNT query using SQLCondition filters
// and executes it.
func (s *BasePostgreSqlService) CountWithFilter(ctx context.Context, filter map[string]sql_query.SQLCondition) (int, error) {
	queryString, args := common_builders.CountBuilder(s.TableName, filter)
	shouldShowQuery(s.debugLevel, queryString, args...)

	return s.Count(ctx, queryString, args...)
}

// SelectOne executes a SELECT query that returns a single row
// and scans the result into the provided struct pointer v
// (e.g., *dto.GetLoggedInUser).
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

// SelectMany executes a SELECT query that returns multiple rows
// and scans the results into the provided slice pointer v
// (e.g., *[]dto.GetCustomFieldsResponse).
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

	if rows.Err() != nil {
		log.Printf("rows error: %v", rows.Err())
		return rows.Err()
	}

	return nil
}

// InsertOne executes an INSERT ... RETURNING id query
// and returns the inserted row ID.
func (s *BasePostgreSqlService) InsertOne(ctx context.Context, queryString string, args ...any) (interface{}, error) {
	shouldShowQuery(s.debugLevel, queryString, args...)

	var resultId string
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

// InsertOneWithData builds and executes an INSERT query for a single row,
// using a struct or map (`body`) as the data source.
//
// Parameters:
//
//	ctx         - Context for timeout/cancellation propagation.
//	body        - Struct or map representing columns and values to insert
//	              (for struct, uses `column` tags).
//	returnOption (optional) - Variadic ReturningConfig to specify:
//	    • Column      - Slice of columns to include in RETURNING clause.
//	    • Destination - Pointer to a struct to scan the inserted row.
//
// Behavior:
//   - If returnOption is not provided → Executes INSERT ... RETURNING id, returns inserted ID.
//   - If returnOption is provided with Destination → Executes INSERT ... RETURNING <columns>
//     and scans the row into the given struct, returning nil as the first return value.
//
// Returns:
//   - interface{} → Inserted ID (string) or nil (if Destination is provided).
//   - error       → Any error encountered during query building, execution, or scanning.
func (s *BasePostgreSqlService) InsertOneWithData(ctx context.Context, body interface{}, returnOption ...ReturningConfig) (interface{}, error) {
	returnColumn := []string{}

	if len(returnOption) > 0 {
		returnColumn = append(returnColumn, returnOption[0].Column...)
	}
	queryString, args := common_builders.InsertBuilder(s.TableName, body, returnColumn...)
	shouldShowQuery(s.debugLevel, queryString, args...)

	if len(returnOption) > 0 && returnOption[0].Destination != nil {
		return nil, s.SelectOne(returnOption[0].Destination, ctx, queryString, args...)
	}
	return s.InsertOne(ctx, queryString, args...)
}

// InsertMany executes an INSERT query for multiple rows
// and returns the number of rows affected.
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

// InsertManyWithData builds and executes an INSERT query for multiple rows,
// using a slice of structs or maps (`body`) as the data source.
//
// Parameters:
//
//	ctx         - Context for timeout/cancellation propagation.
//	body        - Slice of structs or maps representing rows to insert.
//	returnOption (optional) - Variadic ReturningConfig to specify:
//	    • Column      - Slice of columns to include in RETURNING clause.
//	    • Destination - Pointer to a slice to scan all inserted rows.
//
// Behavior:
//   - If returnOption is not provided → Executes INSERT and returns rows affected.
//   - If returnOption is provided with Destination → Executes INSERT ... RETURNING <columns>,
//     scans the rows into Destination, and returns the length of the slice.
//
// Returns:
//   - interface{} → Either int64 (rows affected) or int64 (len of scanned rows).
//   - error       → Any error encountered during query building, execution, or scanning.
//
// Notes:
//   - Destination must be a pointer to a slice (e.g., *[]YourDTO), otherwise it returns an error.
func (s *BasePostgreSqlService) InsertManyWithData(ctx context.Context, body interface{}, returnOption ...ReturningConfig) (interface{}, error) {
	returnColumn := []string{}

	if len(returnOption) > 0 {
		returnColumn = append(returnColumn, returnOption[0].Column...)
	}
	queryString, args := common_builders.InsertBuilder(s.TableName, body, returnColumn...)
	shouldShowQuery(s.debugLevel, queryString, args...)

	if len(returnOption) > 0 && returnOption[0].Destination != nil {
		err := s.SelectMany(returnOption[0].Destination, ctx, queryString, args...)
		val := reflect.ValueOf(returnOption[0].Destination)

		return int64(val.Elem().Len()), err
	}
	return s.InsertMany(ctx, queryString, args...)
}

// UpdateOne executes an UPDATE ... RETURNING id query
// and returns the updated row ID.
func (s *BasePostgreSqlService) UpdateOne(ctx context.Context, queryString string, args ...any) (interface{}, error) {
	shouldShowQuery(s.debugLevel, queryString, args...)

	var resultId string
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

// UpdateOneWithData builds and executes an UPDATE query using a filter map (`query`)
// and a struct or map as the update body (`body`).
//
// Parameters:
//
//	ctx         - The request context for cancellation and deadlines.
//	query       - A map of column names to SQLCondition objects used in the WHERE clause.
//	body        - A struct or map representing the columns to update (uses `column` tags for structs).
//	returnOption (optional) - A variadic argument of ReturningConfig that defines:
//	    • Column      - A slice of column names to include in the RETURNING clause.
//	    • Destination - A pointer to a struct where the returned row will be scanned.
//	                    If Destination is nil, the method returns the updated row's ID as string.
//
// Behavior:
//   - If returnOption is not provided → executes UPDATE ... RETURNING id and returns the ID (string).
//   - If returnOption is provided but Destination is nil → still executes UPDATE and returns the ID.
//   - If returnOption is provided with Destination → scans the updated row into Destination
//     and returns nil as the first return value (caller uses Destination for data).
//
// Returns:
//   - interface{}: The updated row's ID (string) if no Destination is provided.
//     nil if Destination is provided (caller reads from Destination).
//   - error:       Any error encountered during query building, execution, or scanning.
func (s *BasePostgreSqlService) UpdateOneWithData(ctx context.Context, query map[string]sql_query.SQLCondition, body interface{}, returnOption ...ReturningConfig) (interface{}, error) {
	returnColumn := []string{}

	if len(returnOption) > 0 {
		returnColumn = append(returnColumn, returnOption[0].Column...)
	}
	queryString, args := common_builders.UpdateBuilder(s.TableName,
		query,
		body,
		returnColumn...,
	)
	shouldShowQuery(s.debugLevel, queryString, args...)

	if len(returnOption) > 0 && returnOption[0].Destination != nil {
		return nil, s.SelectOne(returnOption[0].Destination, ctx, queryString, args...)
	}
	return s.UpdateOne(ctx, queryString, args...)
}

// UpdateMany executes an UPDATE query that may affect multiple rows
// and returns the number of rows updated.
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

// UpdateManyWithData builds and executes an UPDATE query that may affect multiple rows.
// It supports returning and scanning the updated rows into a destination slice.
//
// Parameters:
//
//	ctx         - The request context for cancellation and deadlines.
//	query       - A map of column names to SQLCondition objects used in the WHERE clause.
//	body        - A struct or map representing the columns to update (uses `column` tags for structs).
//	returnOption (optional) - A variadic argument of ReturningConfig that defines:
//	    • Column      - Slice of column names to include in RETURNING.
//	    • Destination - Pointer to a slice where all updated rows will be scanned.
//
// Behavior:
//   - If no returnOption is provided → executes UPDATE and returns the number of rows affected.
//   - If returnOption is provided with Destination → executes UPDATE ... RETURNING columns
//     and scans all updated rows into Destination (which must be a pointer to a slice).
//     The returned int64 equals the length of the slice.
//   - If Destination is not a pointer to a slice → returns an error.
//
// Returns:
//   - int64: Number of rows affected. If Destination is provided, this equals len(Destination).
//   - error: Any error encountered during query building, execution, or scanning.
func (s *BasePostgreSqlService) UpdateManyWithData(ctx context.Context, query map[string]sql_query.SQLCondition, body interface{}, returnOption ...ReturningConfig) (int64, error) {
	returnColumn := []string{}

	if len(returnOption) > 0 {
		returnColumn = append(returnColumn, returnOption[0].Column...)
	}
	queryString, args := common_builders.UpdateBuilder(s.TableName,
		query,
		body,
		returnColumn...,
	)
	shouldShowQuery(s.debugLevel, queryString, args...)

	if len(returnOption) > 0 && returnOption[0].Destination != nil {
		err := s.SelectMany(returnOption[0].Destination, ctx, queryString, args...)
		val := reflect.ValueOf(returnOption[0].Destination)

		return int64(val.Elem().Len()), err
	}
	return s.UpdateMany(ctx, queryString, args...)
}

// UpdateEachWithData performs a bulk update per row (row-specific values)
// using rowIdentifier (typically a primary key column or unique index).
func (s *BasePostgreSqlService) UpdateEachWithData(ctx context.Context, rowIdentifier string, query map[string]sql_query.SQLCondition, body interface{}) (int64, error) {
	queryString, args := common_builders.UpdateEachBuilder(s.TableName,
		rowIdentifier,
		query,
		body,
	)
	shouldShowQuery(s.debugLevel, queryString, args...)

	return s.UpdateMany(ctx, queryString, args...)
}

// SoftDeleteOne builds and executes a UPDATE soft delete query for a single row
// using SQLCondition filters and returns the deleted row ID.
func (s *BasePostgreSqlService) SoftDeleteOne(ctx context.Context, filter map[string]sql_query.SQLCondition, returnOption ...ReturningConfig) (interface{}, error) {
	returnColumn := []string{}

	if len(returnOption) > 0 {
		returnColumn = append(returnColumn, returnOption[0].Column...)
	}
	queryString, args := common_builders.UpdateBuilder(s.TableName, filter, dto.SetSoftDelete{
		IsDeleted: true,
		DeletedAt: "NOW()",
	}, returnColumn...)
	shouldShowQuery(s.debugLevel, queryString, args...)

	if len(returnOption) > 0 && returnOption[0].Destination != nil {
		return nil, s.SelectOne(returnOption[0].Destination, ctx, queryString, args...)
	}
	return s.DeleteOne(ctx, queryString, args...)
}

// SoftDeleteMany builds and executes a UPDATE soft delete query for multiple rows
// using SQLCondition filters and returns the number of rows deleted.
func (s *BasePostgreSqlService) SoftDeleteMany(ctx context.Context, filter map[string]sql_query.SQLCondition, returnOption ...ReturningConfig) (int64, error) {
	returnColumn := []string{}

	if len(returnOption) > 0 {
		returnColumn = append(returnColumn, returnOption[0].Column...)
	}
	queryString, args := common_builders.UpdateBuilder(s.TableName,
		filter,
		dto.SetSoftDelete{
			IsDeleted: true,
			DeletedAt: "NOW()",
		},
		returnColumn...,
	)
	shouldShowQuery(s.debugLevel, queryString, args...)

	if len(returnOption) > 0 && returnOption[0].Destination != nil {
		err := s.SelectMany(returnOption[0].Destination, ctx, queryString, args...)
		val := reflect.ValueOf(returnOption[0].Destination)

		return int64(val.Elem().Len()), err
	}
	return s.DeleteMany(ctx, queryString, args...)
}

// DeleteOne executes a DELETE ... RETURNING id query
// and returns the deleted row ID.
func (s *BasePostgreSqlService) DeleteOne(ctx context.Context, queryString string, args ...any) (interface{}, error) {
	shouldShowQuery(s.debugLevel, queryString, args...)

	var resultId string
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

// DeleteOneWithFilter builds and executes a DELETE query for a single row
// using SQLCondition filters and returns the deleted row ID.
func (s *BasePostgreSqlService) DeleteOneWithFilter(ctx context.Context, filter map[string]sql_query.SQLCondition) (interface{}, error) {
	queryString, args := common_builders.DeleteBuilder(s.TableName, filter)
	shouldShowQuery(s.debugLevel, queryString, args...)

	return s.DeleteOne(ctx, queryString, args...)
}

// DeleteMany executes a DELETE query that may affect multiple rows
// and returns the number of rows deleted.
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

// DeleteManyWithFilter builds and executes a DELETE query for multiple rows
// using SQLCondition filters and returns the number of rows deleted.
func (s *BasePostgreSqlService) DeleteManyWithFilter(ctx context.Context, filter map[string]sql_query.SQLCondition) (int64, error) {
	queryString, args := common_builders.DeleteBuilder(s.TableName, filter)
	shouldShowQuery(s.debugLevel, queryString, args...)

	return s.DeleteMany(ctx, queryString, args...)
}

// UseTransactions executes fn within a transaction.
// If fn returns an error, the transaction is rolled back.
// If fn succeeds, the transaction is committed.
// The returned value is whatever fn returns, or an error if commit/rollback fails.
func UseTransactions(ctx context.Context, pool PgxPoolInterface, fn func(tx pgx.Tx) (any, error)) (any, error) {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return nil, errors.New("Can't start transactions: " + err.Error())
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback(ctx)
			err = fmt.Errorf("transaction panicked: %v", r)
		} else {
			tx.Rollback(ctx)
		}
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
