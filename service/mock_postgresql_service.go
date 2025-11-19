package service

import (
	"context"

	"gitea.qwertysystem.net/BETS/ts-utils/v2/sql_query"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/mock"
)

// MockPostgreSQLService implements PostgreSQLServicer
//
// Used to mock the PostgreSQLService implementation.
type MockPostgreSQLService struct {
	mock.Mock
}

func (m *MockPostgreSQLService) Init() {
	// Pass
}

func (m *MockPostgreSQLService) Debug(level ...int) {
	m.Called(level)
}

func (m *MockPostgreSQLService) GetPool() PgxPoolInterface {
	arg := m.Called()
	return arg.Get(0).(PgxPoolInterface)
}

func (m *MockPostgreSQLService) GetTransaction() pgx.Tx {
	arg := m.Called()
	return arg.Get(0).(pgx.Tx)
}

func (m *MockPostgreSQLService) SetTransaction(tx pgx.Tx) {
	m.Called(tx)
}

func (m *MockPostgreSQLService) CommitTransaction(ctx context.Context) error {
	arg := m.Called(ctx)
	return arg.Error(0)
}

func (m *MockPostgreSQLService) RollbackTransaction(ctx context.Context) error {
	arg := m.Called(ctx)
	return arg.Error(0)
}

func (m *MockPostgreSQLService) Count(
	ctx context.Context,
	queryString string,
	args ...any,
) (int, error) {
	arg := m.Called(ctx, queryString, args)
	return arg.Get(0).(int), arg.Error(1)
}

func (m *MockPostgreSQLService) CountWithFilter(
	ctx context.Context,
	tableName string,
	filter map[string]sql_query.SQLCondition,
) (int, error) {
	arg := m.Called(ctx, tableName, filter)
	return arg.Get(0).(int), arg.Error(1)
}

func (m *MockPostgreSQLService) Execute(ctx context.Context, queryString string) error {
	arg := m.Called(ctx, queryString)
	return arg.Error(0)
}

func (m *MockPostgreSQLService) SelectOne(
	v any,
	ctx context.Context,
	queryString string,
	args ...any,
) error {
	arg := m.Called(v, ctx, queryString, args)
	return arg.Error(0)
}

func (m *MockPostgreSQLService) SelectMany(
	v any,
	ctx context.Context,
	queryString string,
	args ...any,
) error {
	arg := m.Called(v, ctx, queryString, args)
	return arg.Error(0)
}

func (m *MockPostgreSQLService) InsertOne(
	ctx context.Context,
	queryString string,
	args ...any,
) (any, error) {
	arg := m.Called(ctx, queryString, args)
	return arg.Get(0), arg.Error(1)
}

func (m *MockPostgreSQLService) InsertOneWithData(
	ctx context.Context,
	tableName string,
	body any,
	returnOption ...ReturningConfig,
) (any, error) {
	var arg mock.Arguments

	if len(returnOption) > 0 {
		arg = m.Called(ctx, tableName, body, returnOption)
	} else {
		arg = m.Called(ctx, tableName, body)
	}

	return arg.Get(0), arg.Error(1)
}

func (m *MockPostgreSQLService) InsertMany(
	ctx context.Context,
	queryString string,
	args ...any,
) (int64, error) {
	arg := m.Called(ctx, queryString, args)
	return arg.Get(0).(int64), arg.Error(1)
}

func (m *MockPostgreSQLService) InsertManyWithData(
	ctx context.Context,
	tableName string,
	body any,
	returnOption ...ReturningConfig,
) (any, error) {
	var arg mock.Arguments

	if len(returnOption) > 0 {
		arg = m.Called(ctx, tableName, body, returnOption)
	} else {
		arg = m.Called(ctx, tableName, body)
	}
	return arg.Get(0), arg.Error(1)
}

func (m *MockPostgreSQLService) UpdateOne(
	ctx context.Context,
	queryString string,
	args ...any,
) (any, error) {
	arg := m.Called(ctx, queryString, args)
	return arg.Get(0), arg.Error(1)
}

func (m *MockPostgreSQLService) UpdateOneWithData(
	ctx context.Context,
	tableName string,
	query map[string]sql_query.SQLCondition,
	body any,
	returnOption ...ReturningConfig,
) (any, error) {
	var arg mock.Arguments
	if len(returnOption) > 0 {
		arg = m.Called(ctx, tableName, query, body, returnOption)
	} else {
		arg = m.Called(ctx, tableName, query, body)
	}
	return arg.Get(0), arg.Error(1)
}

func (m *MockPostgreSQLService) UpdateMany(
	ctx context.Context,
	queryString string,
	args ...any,
) (int64, error) {
	arg := m.Called(ctx, queryString, args)
	return arg.Get(0).(int64), arg.Error(1)
}

func (m *MockPostgreSQLService) UpdateManyWithData(
	ctx context.Context,
	tableName string,
	query map[string]sql_query.SQLCondition,
	body any,
	returnOption ...ReturningConfig,
) (int64, error) {
	var arg mock.Arguments

	if len(returnOption) > 0 {
		arg = m.Called(ctx, tableName, query, body, returnOption)
	} else {
		arg = m.Called(ctx, tableName, query, body)
	}
	return arg.Get(0).(int64), arg.Error(1)
}

func (m *MockPostgreSQLService) UpdateEachWithData(
	ctx context.Context,
	tableName string,
	rowIdentifier string,
	query map[string]sql_query.SQLCondition,
	body any,
) (int64, error) {
	arg := m.Called(ctx, tableName, rowIdentifier, query, body)
	return arg.Get(0).(int64), arg.Error(1)
}

func (m *MockPostgreSQLService) SoftDeleteOne(
	ctx context.Context,
	tableName string,
	filter map[string]sql_query.SQLCondition,
	returnOption ...ReturningConfig,
) (any, error) {
	var arg mock.Arguments

	if len(returnOption) > 0 {
		arg = m.Called(ctx, tableName, filter, returnOption)
	} else {
		arg = m.Called(ctx, tableName, filter)
	}

	return arg.Get(0), arg.Error(1)
}

func (m *MockPostgreSQLService) SoftDeleteMany(
	ctx context.Context,
	tableName string,
	filter map[string]sql_query.SQLCondition,
	returnOption ...ReturningConfig,
) (int64, error) {
	var arg mock.Arguments

	if len(returnOption) > 0 {
		arg = m.Called(ctx, tableName, filter, returnOption)
	} else {
		arg = m.Called(ctx, tableName, filter)
	}
	return arg.Get(0).(int64), arg.Error(1)
}

func (m *MockPostgreSQLService) DeleteOne(
	ctx context.Context,
	queryString string,
	args ...any,
) (any, error) {
	arg := m.Called(ctx, queryString, args)
	return arg.Get(0), arg.Error(1)
}

func (m *MockPostgreSQLService) DeleteOneWithFilter(
	ctx context.Context,
	tableName string,
	filter map[string]sql_query.SQLCondition,
) (any, error) {
	arg := m.Called(ctx, tableName, filter)
	return arg.Get(0), arg.Error(1)
}

func (m *MockPostgreSQLService) DeleteMany(
	ctx context.Context,
	queryString string,
	args ...any,
) (int64, error) {
	arg := m.Called(ctx, queryString, args)
	return arg.Get(0).(int64), arg.Error(1)
}

func (m *MockPostgreSQLService) DeleteManyWithFilter(
	ctx context.Context,
	tableName string,
	filter map[string]sql_query.SQLCondition,
) (int64, error) {
	arg := m.Called(ctx, tableName, filter)
	return arg.Get(0).(int64), arg.Error(1)
}

func (m *MockPostgreSQLService) Close() {
	// Pass
}

type MockPgxPool struct {
	mock.Mock
}

func (m *MockPgxPool) CopyFrom(
	ctx context.Context,
	identifier pgx.Identifier,
	columns []string,
	source pgx.CopyFromSource,
) (int64, error) {
	args := m.Called(ctx, identifier, columns, source)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockPgxPool) Exec(
	ctx context.Context,
	sql string,
	args ...any,
) (pgconn.CommandTag, error) {
	called := m.Called(append([]any{ctx, sql}, args...)...)
	return called.Get(0).(pgconn.CommandTag), called.Error(1)
}

func (m *MockPgxPool) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	called := m.Called(append([]any{ctx, sql}, args...)...)
	return called.Get(0).(pgx.Rows), called.Error(1)
}

func (m *MockPgxPool) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	called := m.Called(append([]any{ctx, sql}, args...)...)
	return called.Get(0).(pgx.Row)
}

// Mock Begin(ctx)
func (m *MockPgxPool) Begin(ctx context.Context) (pgx.Tx, error) {
	called := m.Called(ctx)
	return called.Get(0).(pgx.Tx), called.Error(1)
}

// Mock Close()
func (m *MockPgxPool) Close() {
	m.Called()
}

type MockPgxTx struct {
	mock.Mock
}

func (m *MockPgxTx) Begin(ctx context.Context) (pgx.Tx, error) {
	args := m.Called(ctx)
	return args.Get(0).(pgx.Tx), args.Error(1)
}

func (m *MockPgxTx) Commit(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockPgxTx) Rollback(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockPgxTx) CopyFrom(
	ctx context.Context,
	tableName pgx.Identifier,
	columnNames []string,
	rowSrc pgx.CopyFromSource,
) (int64, error) {
	args := m.Called(ctx, tableName, columnNames, rowSrc)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockPgxTx) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	args := m.Called(ctx, b)
	return args.Get(0).(pgx.BatchResults)
}

func (m *MockPgxTx) LargeObjects() pgx.LargeObjects {
	args := m.Called()
	return args.Get(0).(pgx.LargeObjects)
}

func (m *MockPgxTx) Prepare(
	ctx context.Context,
	name, sql string,
) (*pgconn.StatementDescription, error) {
	args := m.Called(ctx, name, sql)
	return args.Get(0).(*pgconn.StatementDescription), args.Error(1)
}

func (m *MockPgxTx) Exec(
	ctx context.Context,
	sql string,
	arguments ...any,
) (pgconn.CommandTag, error) {
	args := m.Called(ctx, sql, arguments)
	return args.Get(0).(pgconn.CommandTag), args.Error(1)
}

func (m *MockPgxTx) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	callArgs := m.Called(ctx, sql, args)
	return callArgs.Get(0).(pgx.Rows), callArgs.Error(1)
}

func (m *MockPgxTx) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	callArgs := m.Called(ctx, sql, args)
	return callArgs.Get(0).(pgx.Row)
}

func (m *MockPgxTx) Conn() *pgx.Conn {
	args := m.Called()
	return args.Get(0).(*pgx.Conn)
}
