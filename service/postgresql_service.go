package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
	"runtime/debug"
	"slices"
	"strings"
	"time"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/db"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/sql_query"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/sql_query/common_builders"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type ReturningConfig struct {
	Column      []string
	Destination any
}

// PostgreSQLServicer provides an interface for implementing PostgreSQLService
type PostgreSQLServicer interface {
	Init()
	// Debug sets the debug level for printing executed SQL queries.
	// Level:
	//
	//	0 → no logs
	//	1 → log queries
	//	2 → log queries and arguments (separated)
	//  3 → log queries and arguments (combined in the string)
	//
	// Defaults to level 1 if an invalid level is passed.
	// Remove its usage in production
	Debug(level ...int)
	// GetPool returns the underlying connection pool (PgxPoolInterface)
	// used by this service.
	GetPool() PgxPoolInterface
	// GetTransaction returns the underlying connection transaction (pgx.Tx)
	// used by this service.
	GetTransaction() pgx.Tx
	// SetTransaction assigns an active transaction (pgx.Tx) to the service.
	// If a transaction is set, all queries will be executed within it.
	SetTransaction(tx pgx.Tx)
	CommitTransaction(ctx context.Context) error
	RollbackTransaction(ctx context.Context) error

	// Count executes a SELECT COUNT(*) query and returns the number of rows.
	Count(ctx context.Context, queryString string, args ...any) (int, error)
	// CountWithFilter builds a COUNT query using SQLCondition filters
	// and executes it.
	CountWithFilter(ctx context.Context, tableName string, filter sql_query.WhereQuery) (int, error)

	// Execute executes a query that wont return any value
	// Useful for defining temporary function and tables/views, so that the overall query is cleaner
	Execute(ctx context.Context, queryString string) error

	// SelectOne executes a SELECT query that returns a single row
	// and scans the result into the provided struct pointer v
	// (e.g., *dto.GetLoggedInUser).
	SelectOne(v any, ctx context.Context, queryString string, args ...any) error
	// SelectMany executes a SELECT query that returns multiple rows
	// and scans the results into the provided slice pointer v
	// (e.g., *[]dto.GetCustomFieldsResponse).
	SelectMany(v any, ctx context.Context, queryString string, args ...any) error

	// InsertOne executes an INSERT ... RETURNING id query
	// and returns the inserted row ID.
	InsertOne(ctx context.Context, queryString string, args ...any) (any, error)
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
	//   - any → Inserted ID (string) or nil (if Destination is provided).
	//   - error       → Any error encountered during query building, execution, or scanning.
	InsertOneWithData(
		ctx context.Context,
		tableName string,
		body any,
		returnOption ...ReturningConfig,
	) (any, error)
	// InsertMany executes an INSERT query for multiple rows
	// and returns the number of rows affected.
	InsertMany(ctx context.Context, queryString string, args ...any) (int64, error)
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
	//   - any → Either int64 (rows affected) or int64 (len of scanned rows).
	//   - error       → Any error encountered during query building, execution, or scanning.
	//
	// Notes:
	//   - Destination must be a pointer to a slice (e.g., *[]YourDTO), otherwise it returns an error.
	InsertManyWithData(
		ctx context.Context,
		tableName string,
		body any,
		returnOption ...ReturningConfig,
	) (any, error)

	// UpdateOne executes an UPDATE ... RETURNING id query
	// and returns the updated row ID.
	UpdateOne(ctx context.Context, queryString string, args ...any) (any, error)
	// UpdateOneWithData builds and executes an UPDATE query using a filter map (`query`)
	// and an UpdateQuery (map[string]any) as the update body (`body`).
	//
	// Parameters:
	//
	//	ctx         - The request context for cancellation and deadlines.
	//	query       - A map of column names to SQLCondition objects used in the WHERE clause.
	//	body        - An UpdateQuery (map[string]any) representing the columns to update.
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
	//   - any: The updated row's ID (string) if no Destination is provided.
	//     nil if Destination is provided (caller reads from Destination).
	//   - error:       Any error encountered during query building, execution, or scanning.
	UpdateOneWithData(
		ctx context.Context,
		tableName string,
		query sql_query.WhereQuery,
		body sql_query.UpdateQuery,
		returnOption ...ReturningConfig,
	) (any, error)
	// UpdateMany executes an UPDATE query that may affect multiple rows
	// and returns the number of rows updated.
	UpdateMany(ctx context.Context, queryString string, args ...any) (int64, error)
	// UpdateManyWithData builds and executes an UPDATE query that may affect multiple rows.
	// It supports returning and scanning the updated rows into a destination slice.
	//
	// Parameters:
	//
	//	ctx         - The request context for cancellation and deadlines.
	//	query       - A map of column names to SQLCondition objects used in the WHERE clause.
	//	body        - An UpdateQuery (map[string]any) representing the columns to update.
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
	UpdateManyWithData(
		ctx context.Context,
		tableName string,
		query sql_query.WhereQuery,
		body sql_query.UpdateQuery,
		returnOption ...ReturningConfig,
	) (int64, error)
	// UpdateEachWithData performs a bulk update per row (row-specific values)
	// using rowIdentifier (typically a primary key column or unique index).
	UpdateEachWithData(
		ctx context.Context,
		tableName string,
		rowIdentifier string,
		query sql_query.WhereQuery,
		body any,
	) (int64, error)

	// SoftDeleteOne builds and executes a UPDATE soft delete query for a single row
	// using SQLCondition filters and returns the deleted row ID.
	// Table must have column is_deleted and deleted_at.
	SoftDeleteOne(
		ctx context.Context,
		tableName string,
		filter sql_query.WhereQuery,
		returnOption ...ReturningConfig,
	) (any, error)
	// SoftDeleteMany builds and executes a UPDATE soft delete query for multiple rows
	// using SQLCondition filters and returns the number of rows deleted.
	// Table must has column is_deleted and deleted_at.
	SoftDeleteMany(
		ctx context.Context,
		tableName string,
		filter sql_query.WhereQuery,
		returnOption ...ReturningConfig,
	) (int64, error)
	// RestoreSoftDeleteOne builds and executes a UPDATE soft delete query for a single row
	// using SQLCondition filters and returns the deleted row ID.
	// Table must have column is_deleted and deleted_at.
	RestoreSoftDeleteOne(
		ctx context.Context,
		tableName string,
		filter sql_query.WhereQuery,
		returnOption ...ReturningConfig,
	) (any, error)
	// RestoreSoftDeleteMany builds and executes a UPDATE soft delete query for multiple rows
	// using SQLCondition filters and returns the number of rows deleted.
	// Table must has column is_deleted and deleted_at.
	RestoreSoftDeleteMany(
		ctx context.Context,
		tableName string,
		filter sql_query.WhereQuery,
		returnOption ...ReturningConfig,
	) (int64, error)

	// DeleteOne executes a DELETE ... RETURNING id query
	// and returns the deleted row ID.
	DeleteOne(ctx context.Context, queryString string, args ...any) (any, error)
	// DeleteOneWithFilter builds and executes a DELETE query for a single row
	// using SQLCondition filters and returns the deleted row ID.
	DeleteOneWithFilter(
		ctx context.Context,
		tableName string,
		filter sql_query.WhereQuery,
	) (any, error)
	// DeleteMany executes a DELETE query that may affect multiple rows
	// and returns the number of rows deleted.
	DeleteMany(ctx context.Context, queryString string, args ...any) (int64, error)
	// DeleteManyWithFilter builds and executes a DELETE query for multiple rows
	// using SQLCondition filters and returns the number of rows deleted.
	DeleteManyWithFilter(
		ctx context.Context,
		tableName string,
		filter sql_query.WhereQuery,
	) (int64, error)
	Close()
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
	CopyFrom(ctx context.Context, identifier pgx.Identifier, columns []string, source pgx.CopyFromSource) (int64, error)
	// Close closes the database pool connection.
	Close()
}

// PostgreSQLService implements PostgreSQLServicer
//
// Full implementation of CRUD operation for PostgreSQL.
type PostgreSQLService struct {
	Pool        PgxPoolInterface
	Transaction pgx.Tx
	schemaName  string
	// debugLevel debug level set for logging query.
	debugLevel int
}

// MakePostgreSQLService creates a new PostgreSQLServicer instance.
func MakePostgreSQLService(dbNameWithSchema db.DBName, options ...db.PostgreSQLServiceOptions) *PostgreSQLService {
	splitted := strings.Split(string(dbNameWithSchema), ":")
	if len(splitted) < 2 {
		// If schema is not specified, use public as default
		splitted = []string{"public", string(dbNameWithSchema)}
	}

	// Pass actual database name to ConnectPostgres
	pool := db.ConnectPostgres(db.DBName(splitted[1]), options...)

	return &PostgreSQLService{
		Pool:       pool,
		schemaName: splitted[0],
	}
}

func (s *PostgreSQLService) Init() {
	// TODO(refan): initialization can be moved here.
	// Pass
}

func (s *PostgreSQLService) Debug(level ...int) {
	if len(level) > 0 && slices.Contains([]int{1, 2, 3}, level[0]) {
		s.debugLevel = level[0]
		return
	}

	s.debugLevel = 1
}

func (s *PostgreSQLService) GetPool() PgxPoolInterface {
	return s.Pool
}

func (s *PostgreSQLService) GetTransaction() pgx.Tx {
	return s.Transaction
}

func (s *PostgreSQLService) SetTransaction(tx pgx.Tx) {
	s.Transaction = tx
}

func (s *PostgreSQLService) CommitTransaction(ctx context.Context) error {
	if s.Transaction == nil {
		return errors.New("no active transaction to commit")
	}

	if err := s.Transaction.Commit(ctx); err != nil {
		_ = s.Transaction.Rollback(ctx)
		return err
	}

	return nil
}

func (s *PostgreSQLService) RollbackTransaction(ctx context.Context) error {
	if s.Transaction == nil {
		return errors.New("no active transaction to rollback")
	}

	_ = s.Transaction.Rollback(ctx)
	return nil
}

func (s *PostgreSQLService) Count(
	ctx context.Context,
	queryString string,
	args ...any,
) (int, error) {
	var count int
	if err := s.executeQueryRow(ctx, queryString, args, &count, true); err != nil {
		log.Println("Count query failed:", err)
		return 0, err
	}
	return count, nil
}

func (s *PostgreSQLService) Execute(
	ctx context.Context,
	queryString string,
) error {
	rows, err := s.executeQuery(ctx, queryString, nil)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (s *PostgreSQLService) CountWithFilter(
	ctx context.Context,
	tableName string,
	filter sql_query.WhereQuery,
) (int, error) {
	queryString, args := common_builders.CountBuilder(tableName, filter)

	return s.Count(ctx, queryString, args...)
}

func (s *PostgreSQLService) SelectOne(
	v any,
	ctx context.Context,
	queryString string,
	args ...any,
) error {
	rows, err := s.executeQuery(ctx, queryString, args)
	if err != nil {
		return err
	}
	defer rows.Close()

	err = sql_query.ScanRowObject(v, rows.Rows)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *PostgreSQLService) SelectMany(
	v any,
	ctx context.Context,
	queryString string,
	args ...any,
) error {
	rows, err := s.executeQuery(ctx, queryString, args)
	if err != nil {
		return err
	}
	defer rows.Close()

	err = sql_query.ScanRowsArray(v, rows.Rows)
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

func (s *PostgreSQLService) InsertOne(
	ctx context.Context,
	queryString string,
	args ...any,
) (any, error) {
	var resultId int
	if err := s.executeQueryRow(ctx, queryString, args, &resultId); err != nil {
		return nil, err
	}
	return resultId, nil
}

func (s *PostgreSQLService) InsertOneWithData(
	ctx context.Context,
	tableName string,
	body any,
	returnOption ...ReturningConfig,
) (any, error) {
	returnColumn := []string{}

	if len(returnOption) > 0 {
		returnColumn = append(returnColumn, returnOption[0].Column...)
	}
	queryString, args := common_builders.InsertBuilder(tableName, body, returnColumn...)

	if len(returnOption) > 0 && returnOption[0].Destination != nil {
		return nil, s.SelectOne(returnOption[0].Destination, ctx, queryString, args...)
	}
	return s.InsertOne(ctx, queryString, args...)
}

func (s *PostgreSQLService) InsertMany(
	ctx context.Context,
	queryString string,
	args ...any,
) (int64, error) {
	return s.executeExec(ctx, queryString, args)
}

func (s *PostgreSQLService) InsertManyWithData(
	ctx context.Context,
	tableName string,
	body any,
	returnOption ...ReturningConfig,
) (any, error) {
	returnColumn := []string{}

	if len(returnOption) > 0 {
		returnColumn = append(returnColumn, returnOption[0].Column...)
	}
	queryString, args := common_builders.InsertBuilder(tableName, body, returnColumn...)

	if len(returnOption) > 0 && returnOption[0].Destination != nil {
		err := s.SelectMany(returnOption[0].Destination, ctx, queryString, args...)
		val := reflect.ValueOf(returnOption[0].Destination)

		return int64(val.Elem().Len()), err
	}
	return s.InsertMany(ctx, queryString, args...)
}

// Still in experimental stage, recommended to use InsertManyWithData until this function stable
func (s *PostgreSQLService) InsertBatch(
	ctx context.Context,
	tableName string,
	body any,
) (int64, error) {
	v := reflect.ValueOf(body)
	if v.Kind() != reflect.Slice {
		return 0, errors.New("body must be a slice")
	}
	if v.Len() == 0 {
		return 0, nil
	}

	firstElem := v.Index(0).Interface()
	t := reflect.TypeOf(firstElem)
	typeName := t.PkgPath() + "." + t.Name()

	// Ambil atau buat template
	insertTemplate, ok := sql_query.InsertCache[typeName]
	if !ok {
		insertTemplate = sql_query.BuildInsertTemplate(t) // disini tentukan useNow & useID
		sql_query.InsertCache[typeName] = insertTemplate
	}

	rows := make([][]any, v.Len())
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		row := make([]any, len(insertTemplate.InsertColumn))

		for j := range insertTemplate.InsertColumn {
			if insertTemplate.UseID[j] {
				row[j] = db.Node.Generate().Int64()
			} else if insertTemplate.UseNow[j] {
				row[j] = time.Now() // CopyFrom tidak bisa pakai literal NOW(), harus diisi value
			} else {
				row[j] = elem.FieldByIndex(insertTemplate.FieldIndexes[j]).Interface()
			}
		}
		rows[i] = row
	}

	if s.Transaction != nil {
		return s.Transaction.CopyFrom(
			ctx,
			pgx.Identifier{tableName},
			insertTemplate.InsertColumn,
			pgx.CopyFromRows(rows),
		)
	}

	return s.Pool.CopyFrom(
		ctx,
		pgx.Identifier{tableName},
		insertTemplate.InsertColumn,
		pgx.CopyFromRows(rows),
	)
}

func (s *PostgreSQLService) UpdateOne(
	ctx context.Context,
	queryString string,
	args ...any,
) (any, error) {
	var resultId int
	if err := s.executeQueryRow(ctx, queryString, args, &resultId); err != nil {
		return nil, err
	}
	return resultId, nil
}

func (s *PostgreSQLService) UpdateOneWithData(
	ctx context.Context,
	tableName string,
	query sql_query.WhereQuery,
	body sql_query.UpdateQuery,
	returnOption ...ReturningConfig,
) (any, error) {
	returnColumn := []string{}

	if len(returnOption) > 0 {
		returnColumn = append(returnColumn, returnOption[0].Column...)
	}
	queryString, args := common_builders.UpdateBuilder(tableName,
		query,
		body,
		returnColumn...,
	)

	if len(returnOption) > 0 && returnOption[0].Destination != nil {
		return nil, s.SelectOne(returnOption[0].Destination, ctx, queryString, args...)
	}
	return s.UpdateOne(ctx, queryString, args...)
}

func (s *PostgreSQLService) UpdateMany(
	ctx context.Context,
	queryString string,
	args ...any,
) (int64, error) {
	return s.executeExec(ctx, queryString, args)
}

func (s *PostgreSQLService) UpdateManyWithData(
	ctx context.Context,
	tableName string,
	query sql_query.WhereQuery,
	body sql_query.UpdateQuery,
	returnOption ...ReturningConfig,
) (int64, error) {
	returnColumn := []string{}

	if len(returnOption) > 0 {
		returnColumn = append(returnColumn, returnOption[0].Column...)
	}
	queryString, args := common_builders.UpdateBuilder(tableName,
		query,
		body,
		returnColumn...,
	)

	if len(returnOption) > 0 && returnOption[0].Destination != nil {
		err := s.SelectMany(returnOption[0].Destination, ctx, queryString, args...)
		val := reflect.ValueOf(returnOption[0].Destination)

		return int64(val.Elem().Len()), err
	}
	return s.UpdateMany(ctx, queryString, args...)
}

func (s *PostgreSQLService) UpdateEachWithData(
	ctx context.Context,
	tableName string,
	rowIdentifier string,
	query sql_query.WhereQuery,
	body any,
) (int64, error) {
	queryString, args := common_builders.UpdateEachBuilder(tableName,
		rowIdentifier,
		query,
		body,
	)

	return s.UpdateMany(ctx, queryString, args...)
}

func (s *PostgreSQLService) SoftDeleteOne(
	ctx context.Context,
	tableName string,
	filter sql_query.WhereQuery,
	returnOption ...ReturningConfig,
) (any, error) {
	returnColumn := []string{}

	if len(returnOption) > 0 {
		returnColumn = append(returnColumn, returnOption[0].Column...)
	}
	queryString, args, err := sql_query.NewSQLUpdateBuilder(tableName).
		Update(sql_query.UpdateQuery{
			"is_deleted": true,
			"deleted_at": "NOW()",
		}, true).
		Return(returnColumn...).
		Where(filter).
		Build()
	if err != nil {
		log.Println(err)
	}
	shouldShowQuery(s.debugLevel, queryString, args...)

	if len(returnOption) > 0 && returnOption[0].Destination != nil {
		return nil, s.SelectOne(returnOption[0].Destination, ctx, queryString, args...)
	}
	return s.DeleteOne(ctx, queryString, args...)
}

func (s *PostgreSQLService) SoftDeleteMany(
	ctx context.Context,
	tableName string,
	filter sql_query.WhereQuery,
	returnOption ...ReturningConfig,
) (int64, error) {
	returnColumn := []string{}

	if len(returnOption) > 0 {
		returnColumn = append(returnColumn, returnOption[0].Column...)
	}
	queryString, args, err := sql_query.NewSQLUpdateBuilder(tableName).
		Update(sql_query.UpdateQuery{
			"is_deleted": true,
			"deleted_at": "NOW()",
		}, true).
		Return(returnColumn...).
		Where(filter).
		Build()
	if err != nil {
		log.Println(err)
	}
	shouldShowQuery(s.debugLevel, queryString, args...)

	if len(returnOption) > 0 && returnOption[0].Destination != nil {
		err := s.SelectMany(returnOption[0].Destination, ctx, queryString, args...)
		val := reflect.ValueOf(returnOption[0].Destination)

		return int64(val.Elem().Len()), err
	}
	return s.DeleteMany(ctx, queryString, args...)
}

func (s *PostgreSQLService) RestoreSoftDeleteOne(
	ctx context.Context,
	tableName string,
	filter sql_query.WhereQuery,
	returnOption ...ReturningConfig,
) (any, error) {
	returnColumn := []string{}

	if len(returnOption) > 0 {
		returnColumn = append(returnColumn, returnOption[0].Column...)
	}
	queryString, args, err := sql_query.NewSQLUpdateBuilder(tableName).Update(sql_query.UpdateQuery{
		"is_deleted": false,
		"deleted_at": nil,
	}, true).
		Return(returnColumn...).
		Where(filter).
		Build()
	if err != nil {
		log.Println(err)
	}
	shouldShowQuery(s.debugLevel, queryString, args...)

	if len(returnOption) > 0 && returnOption[0].Destination != nil {
		return nil, s.SelectOne(returnOption[0].Destination, ctx, queryString, args...)
	}
	return s.DeleteOne(ctx, queryString, args...)
}

func (s *PostgreSQLService) RestoreSoftDeleteMany(
	ctx context.Context,
	tableName string,
	filter sql_query.WhereQuery,
	returnOption ...ReturningConfig,
) (int64, error) {
	returnColumn := []string{}

	if len(returnOption) > 0 {
		returnColumn = append(returnColumn, returnOption[0].Column...)
	}
	queryString, args, err := sql_query.NewSQLUpdateBuilder(tableName).Update(sql_query.UpdateQuery{
		"is_deleted": false,
		"deleted_at": nil,
	}, true).
		Return(returnColumn...).
		Where(filter).
		Build()
	if err != nil {
		log.Println(err)
	}
	shouldShowQuery(s.debugLevel, queryString, args...)

	if len(returnOption) > 0 && returnOption[0].Destination != nil {
		err := s.SelectMany(returnOption[0].Destination, ctx, queryString, args...)
		val := reflect.ValueOf(returnOption[0].Destination)

		return int64(val.Elem().Len()), err
	}
	return s.DeleteMany(ctx, queryString, args...)
}

func (s *PostgreSQLService) DeleteOne(
	ctx context.Context,
	queryString string,
	args ...any,
) (any, error) {
	var resultId int
	if err := s.executeQueryRow(ctx, queryString, args, &resultId); err != nil {
		return nil, err
	}
	return resultId, nil
}

func (s *PostgreSQLService) DeleteOneWithFilter(
	ctx context.Context,
	tableName string,
	filter sql_query.WhereQuery,
) (any, error) {
	queryString, args := common_builders.DeleteBuilder(tableName, filter)

	return s.DeleteOne(ctx, queryString, args...)
}

func (s *PostgreSQLService) DeleteMany(
	ctx context.Context,
	queryString string,
	args ...any,
) (int64, error) {
	return s.executeExec(ctx, queryString, args)
}

func (s *PostgreSQLService) DeleteManyWithFilter(
	ctx context.Context,
	tableName string,
	filter sql_query.WhereQuery,
) (int64, error) {
	queryString, args := common_builders.DeleteBuilder(tableName, filter)

	return s.DeleteMany(ctx, queryString, args...)
}

func (s *PostgreSQLService) Close() {
	s.Pool.Close()
}

// applySchema executes SET LOCAL search_path on the given transaction.
func (s *PostgreSQLService) applySchema(ctx context.Context, tx pgx.Tx) error {
	_, err := tx.Exec(ctx, fmt.Sprintf("SET LOCAL search_path TO '%s'", s.schemaName))
	return err
}

// executeQueryRow handles transaction/pool selection for QueryRow operations.
// It automatically calls shouldShowQuery for debug logging.
//
// Parameters:
//   - ctx: context for the operation
//   - queryString: the SQL query to execute
//   - args: query arguments
//   - dest: pointer to destination variable for Scan (e.g., &resultId)
//
// Returns:
//   - error: any error from QueryRow or Scan
func (s *PostgreSQLService) executeQueryRow(
	ctx context.Context,
	queryString string,
	args []any,
	dest any,
	isCount ...bool,
) error {
	shouldShowQuery(s.debugLevel, queryString, args...)

	if len(isCount) == 0 || !isCount[0] {
		// Check for replicated table restrictions
		if err := s.validateTableOperation(queryString); err != nil {
			return err
		}
	}

	if s.Transaction != nil {
		if err := s.applySchema(ctx, s.Transaction); err != nil {
			return err
		}
		return s.Transaction.QueryRow(ctx, queryString, args...).Scan(dest)
	}

	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if err := s.applySchema(ctx, tx); err != nil {
		return err
	}
	err = tx.QueryRow(ctx, queryString, args...).Scan(dest)
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}

// executeExec handles transaction/pool selection for Exec operations.
// It automatically calls shouldShowQuery for debug logging.
//
// Parameters:
//   - ctx: context for the operation
//   - queryString: the SQL query to execute
//   - args: query arguments
//
// Returns:
//   - int64: number of rows affected
//   - error: any error from Exec
func (s *PostgreSQLService) executeExec(
	ctx context.Context,
	queryString string,
	args []any,
) (int64, error) {
	shouldShowQuery(s.debugLevel, queryString, args...)

	// Check for replicated table restrictions
	if err := s.validateTableOperation(queryString); err != nil {
		return 0, err
	}

	if s.Transaction != nil {
		if err := s.applySchema(ctx, s.Transaction); err != nil {
			return 0, err
		}
		commandTag, err := s.Transaction.Exec(ctx, queryString, args...)
		if err != nil {
			return 0, err
		}
		return commandTag.RowsAffected(), nil
	}

	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	if err := s.applySchema(ctx, tx); err != nil {
		return 0, err
	}
	commandTag, err := tx.Exec(ctx, queryString, args...)
	if err != nil {
		return 0, err
	}
	return commandTag.RowsAffected(), tx.Commit(ctx)
}

type txRows struct {
	pgx.Rows
	tx pgx.Tx
}

func (t *txRows) Close() {
	t.Rows.Close()
	if t.tx != nil {
		t.tx.Rollback(context.Background())
	}
}

// executeQuery handles transaction/pool selection for Query operations.
// It automatically calls shouldShowQuery for debug logging.
// Caller is responsible for closing the returned rows.
//
// Parameters:
//   - ctx: context for the operation
//   - queryString: the SQL query to execute
//   - args: query arguments
//
// Returns:
//   - txRows: the result rows (must be closed by caller)
//   - error: any error from Query
func (s *PostgreSQLService) executeQuery(
	ctx context.Context,
	queryString string,
	args []any,
) (txRows, error) {
	shouldShowQuery(s.debugLevel, queryString, args...)

	if s.Transaction != nil {
		if err := s.applySchema(ctx, s.Transaction); err != nil {
			return txRows{}, err
		}
		rows, err := s.Transaction.Query(ctx, queryString, args...)
		if err != nil {
			return txRows{}, err
		}
		return txRows{Rows: rows, tx: nil}, nil
	}

	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return txRows{}, err
	}

	if err := s.applySchema(ctx, tx); err != nil {
		tx.Rollback(ctx)
		return txRows{}, err
	}
	rows, err := tx.Query(ctx, queryString, args...)
	if err != nil {
		tx.Rollback(ctx)
		return txRows{}, err
	}
	// NOTE: tx is committed after rows are fully consumed.
	// The caller closes rows via defer rows.Close(), and pgx auto-releases
	// the tx connection when rows are closed. For read queries this is safe
	// because SET LOCAL auto-resets when the tx ends.
	return txRows{Rows: rows, tx: tx}, nil
}

// UseTransactions executes fn within a transaction.
// DEPRECATED!
// If fn returns an error, the transaction is rolled back.
// If fn succeeds, the transaction is committed.
// The returned value is whatever fn returns, or an error if commit/rollback fails.
//
// ⚠️ WARNING about holdCommit mode:
// When holdCommit is true, this function will NOT commit or rollback.
// Instead, the caller is fully responsible for calling tx.Rollback or tx.Commit.
// Forgetting to do so will:
//   - Leave the connection "dirty" and unusable for other queries
//   - Hold locks and temporary resources open in PostgreSQL
//   - Eventually cause connection pool exhaustion under load
//
// Although pgx will automatically rollback when the connection is returned
// to the pool, you must NOT rely on this. Always explicitly close the tx.
func UseTransactions[T any](
	ctx context.Context,
	pool PgxPoolInterface,
	fn func(tx pgx.Tx) (T, error),
	holdCommit ...bool,
) (result T, err error) {
	tx, err := pool.Begin(ctx)
	if err != nil {
		log.Printf("can't start transactions: %v", err)
		err = errors.New("something went wrong")
		return
	}

	// When holdCommit is true, caller is responsible for calling Rollback or Commit. Transaction is returned unclosed.
	if len(holdCommit) > 0 && holdCommit[0] {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("transaction panicked: %v\n%s", r, debug.Stack())
				err = errors.New("something went wrong")
			}
		}()

		result, err = fn(tx)
		if err != nil {
			log.Printf("Error %v\n", err.Error())
		}

		return result, err
	}

	defer func() {
		if r := recover(); r != nil {
			log.Printf("transaction panicked: %v\n%s", r, debug.Stack())
			_ = tx.Rollback(ctx)
			err = errors.New("something went wrong")
		} else {
			_ = tx.Rollback(ctx)
		}
	}()

	result, err = fn(tx)
	if err != nil {
		log.Printf("Error %v\n", err.Error())
		return
	}

	if commitErr := tx.Commit(ctx); commitErr != nil {
		log.Printf("failed to commit: %v", commitErr)
		err = errors.New("something went wrong")
		return
	}

	return result, nil
}

// shouldShowQuery prints SQL queries for debugging purposes.
// Uses fmt.Printf instead of slog because:
// - slog escapes \n and \t as literal strings, making SQL unreadable
// - fmt.Printf preserves actual newlines and formatting for better readability
// - This is development debug logging, not production structured logging
func shouldShowQuery(level int, query string, args ...any) {
	switch level {
	case 1:
		fmt.Printf("\n=== SQL Query ===\n%s\n", query)
	case 2:
		fmt.Printf("\n=== SQL Query ===\n%s\nArgs: %v\n", query, args)
	case 3:
		fmt.Printf("\n=== SQL Query ===\n%s\n", sql_query.InterpolateQuery(query, args...))
	}
}

// validateTableOperation checks if the query corresponds to a CUD operation on a replicated table
func (s *PostgreSQLService) validateTableOperation(query string) error {
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		return nil
	}

	// Check if service has any replicated tables
	// ReplicatedTables is defined in replication_manager.go
	replicatedSources, ok := ReplicatedTables[serviceName]
	if !ok || len(replicatedSources) == 0 {
		return nil
	}

	tableName := extractTableName(query)
	if tableName == "" {
		return nil
	}

	for _, tables := range replicatedSources {
		for _, table := range tables {
			if table.GetTableName() == tableName {
				return fmt.Errorf("CUD operation on replicated table '%s' is prohibited for service '%s'", tableName, serviceName)
			}
		}
	}

	return nil
}

// extractTableName attempts to extract the table name from INSERT, UPDATE, DELETE queries
func extractTableName(query string) string {
	// Pattern to match:
	// - Command: INSERT INTO, UPDATE, DELETE FROM (case insensitive)
	// - Optional schema and quotes
	// - Capture table name
	const pattern = `(?i)\b(?:INSERT\s+INTO|UPDATE|DELETE\s+FROM)\s+(?:["` + "`" + `]?(?:[\w]+\.)?([\w]+)["` + "`" + `]?)`
	re := regexp.MustCompile(pattern)

	matches := re.FindStringSubmatch(query)
	if len(matches) > 1 {
		return matches[1]
	}

	return ""
}
