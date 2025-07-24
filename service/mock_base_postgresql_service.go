package service

import (
	"context"

	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/mock"
)

type MockBasePostgreSqlService struct {
	mock.Mock
}

func (m *MockBasePostgreSqlService) Debug(level ...int) {
	m.Called(level)
	return
}

func (m *MockBasePostgreSqlService) GetPool() PgxPoolInterface {
	arg := m.Called()
	return arg.Get(0).(PgxPoolInterface)
}

func (m *MockBasePostgreSqlService) GetTableName() string {
	arg := m.Called()
	return arg.Get(0).(string)
}

func (m *MockBasePostgreSqlService) SelectMany(
	v any,
	ctx context.Context,
	queryString string,
	args ...any,
) error {
	arg := m.Called(v, ctx, queryString, args)
	return arg.Error(0)
}

func (m *MockBasePostgreSqlService) SelectOne(
	v any,
	ctx context.Context,
	queryString string,
	args ...any,
) error {
	arg := m.Called(v, ctx, queryString, args)
	return arg.Error(0)
}

func (m *MockBasePostgreSqlService) Count(ctx context.Context, queryString string, args ...any) (int, error) {
	arg := m.Called(ctx, queryString, args)
	return arg.Get(0).(int), arg.Error(1)
}

func (m *MockBasePostgreSqlService) CountWithFilter(ctx context.Context, filter map[string]sql_query.SQLCondition) (int, error) {
	arg := m.Called(ctx, filter)
	return arg.Get(0).(int), arg.Error(1)
}

func (m *MockBasePostgreSqlService) InsertOne(
	ctx context.Context,
	queryString string,
	args ...any,
) (interface{}, error) {
	arg := m.Called(ctx, queryString, args)
	return arg.Get(0), arg.Error(1)
}

func (m *MockBasePostgreSqlService) InsertMany(
	ctx context.Context,
	queryString string,
	args ...any,
) (int64, error) {
	arg := m.Called(ctx, queryString, args)
	return arg.Get(0).(int64), arg.Error(1)
}

func (m *MockBasePostgreSqlService) InsertOneWithData(ctx context.Context, body interface{}) (interface{}, error) {
	arg := m.Called(ctx, body)
	return arg.Get(0), arg.Error(1)
}

func (m *MockBasePostgreSqlService) InsertManyWithData(ctx context.Context, body []interface{}) (interface{}, error) {
	arg := m.Called(ctx, body)
	return arg.Get(0), arg.Error(1)
}

func (m *MockBasePostgreSqlService) UpdateOne(
	ctx context.Context,
	queryString string,
	args ...any,
) (interface{}, error) {
	arg := m.Called(ctx, queryString, args)
	return arg.Get(0), arg.Error(1)
}

func (m *MockBasePostgreSqlService) UpdateOneWithData(ctx context.Context, query map[string]sql_query.SQLCondition, updateData interface{}) (interface{}, error) {
	arg := m.Called(ctx, query, updateData)
	return arg.Get(0), arg.Error(1)
}

func (m *MockBasePostgreSqlService) UpdateManyWithData(ctx context.Context, query map[string]sql_query.SQLCondition, updateData interface{}) (interface{}, error) {
	arg := m.Called(ctx, query, updateData)
	return arg.Get(0).(int64), arg.Error(1)
}

func (m *MockBasePostgreSqlService) UpdateOneWithDataAndReturning(
	v any,
	ctx context.Context,
	query map[string]sql_query.SQLCondition,
	body interface{},
	returning ...string,
) error {
	arg := m.Called(v, ctx, query, body, returning)
	return arg.Error(0)
}

func (m *MockBasePostgreSqlService) UpdateMany(
	ctx context.Context,
	queryString string,
	args ...any,
) (int64, error) {
	arg := m.Called(ctx, queryString, args)
	return arg.Get(0).(int64), arg.Error(1)
}

func (m *MockBasePostgreSqlService) UpdateEachWithData(ctx context.Context, rowIdentifier string, query map[string]sql_query.SQLCondition, updateData []interface{}) (int64, error) {
	arg := m.Called(ctx, rowIdentifier, query, updateData)
	return arg.Get(0).(int64), arg.Error(1)
}

func (m *MockBasePostgreSqlService) DeleteOne(
	ctx context.Context,
	queryString string,
	args ...any,
) (interface{}, error) {
	arg := m.Called(ctx, queryString, args)
	return arg.Get(0), arg.Error(1)
}

func (m *MockBasePostgreSqlService) DeleteOneWithFilter(ctx context.Context, filter map[string]sql_query.SQLCondition) (interface{}, error) {
	arg := m.Called(ctx, filter)
	return arg.Get(0), arg.Error(1)
}

func (m *MockBasePostgreSqlService) DeleteMany(
	ctx context.Context,
	queryString string,
	args ...any,
) (int64, error) {
	arg := m.Called(ctx, queryString, args)
	return arg.Get(0).(int64), arg.Error(1)
}

func (m *MockBasePostgreSqlService) SetTransaction(tx pgx.Tx) {
	m.Called(tx)
}

type MockPgxPool struct {
	mock.Mock
}

func (m *MockPgxPool) Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	called := m.Called(append([]interface{}{ctx, sql}, args...)...)
	return called.Get(0).(pgconn.CommandTag), called.Error(1)
}

func (m *MockPgxPool) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	called := m.Called(append([]interface{}{ctx, sql}, args...)...)
	return called.Get(0).(pgx.Rows), called.Error(1)
}

func (m *MockPgxPool) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	called := m.Called(append([]interface{}{ctx, sql}, args...)...)
	return called.Get(0).(pgx.Row)
}

// Mock Begin(ctx)
func (m *MockPgxPool) Begin(ctx context.Context) (pgx.Tx, error) {
	called := m.Called(ctx)
	return called.Get(0).(pgx.Tx), called.Error(1)
}

// Mock Close()
func (m *MockPgxPool) Close() {
	return
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

func (m *MockPgxTx) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
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

func (m *MockPgxTx) Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error) {
	args := m.Called(ctx, name, sql)
	return args.Get(0).(*pgconn.StatementDescription), args.Error(1)
}

func (m *MockPgxTx) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
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
