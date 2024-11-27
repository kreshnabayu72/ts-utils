package service

import (
	"context"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MockBaseService struct {
	mock.Mock
}

func (m *MockBaseService) Find(v any, ctx context.Context, filter interface{}, opts ...*options.FindOptions) error {
	args := m.Called(v, ctx, filter, opts)
	return args.Error(0)
}

func (m *MockBaseService) FindOne(v any, ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) error {
	args := m.Called(v, ctx, filter, opts)
	return args.Error(0)
}

func (m *MockBaseService) InsertOne(ctx context.Context, data interface{}, opts ...*options.InsertOneOptions) (interface{}, error) {
	args := m.Called(ctx, data, opts)
	return args.Get(0), args.Error(1)
}

func (m *MockBaseService) InsertMany(ctx context.Context, data []interface{}, opts ...*options.InsertManyOptions) ([]interface{}, error) {
	args := m.Called(ctx, data, opts)
	return args.Get(0).([]interface{}), args.Error(1)
}

func (m *MockBaseService) FindOneAndUpdate(v any, ctx context.Context, filter interface{}, updateData interface{}, opts ...*options.FindOneAndUpdateOptions) error {
	args := m.Called(v, ctx, filter, updateData, opts)
	return args.Error(0)
}

func (m *MockBaseService) UpdateOne(ctx context.Context, filter interface{}, updateData interface{}, opts ...*options.UpdateOptions) (int64, error) {
	args := m.Called(ctx, filter, updateData, opts)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockBaseService) UpdateMany(ctx context.Context, filter interface{}, updateData interface{}, opts ...*options.UpdateOptions) (int64, error) {
	args := m.Called(ctx, filter, updateData, opts)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockBaseService) FindOneAndDelete(v any, ctx context.Context, filter interface{}, opts ...*options.FindOneAndDeleteOptions) error {
	args := m.Called(v, ctx, filter, opts)
	return args.Error(0)
}

func (m *MockBaseService) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (int64, error) {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockBaseService) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (int64, error) {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockBaseService) Aggregate(v any, ctx context.Context, pipeline mongo.Pipeline, opts ...*options.AggregateOptions) error {
	args := m.Called(v, ctx, pipeline, opts)
	return args.Error(0)
}

func (m *MockBaseService) MakeUnique(ctx context.Context, filter interface{}) error {
	args := m.Called(ctx, filter)
	return args.Error(0)
}
