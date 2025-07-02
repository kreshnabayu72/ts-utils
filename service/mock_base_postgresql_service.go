package service

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/mock"
)

type MockBasePostgreSqlService struct {
	mock.Mock
}

func (m *MockBasePostgreSqlService) BaseService() BasePostgreSqlService {
	arg := m.Called()
	return arg.Get(0).(BasePostgreSqlService)
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

func (m *MockBasePostgreSqlService) InsertOne(
	ctx context.Context,
	queryString string,
	args ...any,
) (interface{}, error) {
	arg := m.Called(ctx, queryString, args)
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

func (m *MockBasePostgreSqlService) DeleteOne(
	ctx context.Context,
	queryString string,
	args ...any,
) (interface{}, error) {
	arg := m.Called(ctx, queryString, args)
	return arg.Get(0), arg.Error(1)
}
func (m *MockBasePostgreSqlService) SetTransaction(tx pgx.Tx) {
	m.Called(tx)
}
