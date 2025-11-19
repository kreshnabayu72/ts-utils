package service

import (
	"context"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MockMongoDBService implements MongoDBServicer
//
// Used to mock the MongoDBService implementation.
type MockMongoDBService struct {
	mock.Mock
}

func (m *MockMongoDBService) Init() {
	// Pass
}

func (m *MockMongoDBService) Find(
	v any,
	ctx context.Context,
	filter any,
	opts ...*options.FindOptions,
) error {
	args := m.Called(v, ctx, filter, opts)
	return args.Error(0)
}

func (m *MockMongoDBService) FindOne(
	v any,
	ctx context.Context,
	filter any,
	opts ...*options.FindOneOptions,
) error {
	args := m.Called(v, ctx, filter, opts)
	return args.Error(0)
}

func (m *MockMongoDBService) InsertOne(
	ctx context.Context,
	data any,
	opts ...*options.InsertOneOptions,
) (any, error) {
	args := m.Called(ctx, data, opts)
	return args.Get(0), args.Error(1)
}

func (m *MockMongoDBService) InsertMany(
	ctx context.Context,
	data []any,
	opts ...*options.InsertManyOptions,
) ([]any, error) {
	args := m.Called(ctx, data, opts)
	return args.Get(0).([]any), args.Error(1)
}

func (m *MockMongoDBService) FindOneAndUpdate(
	v any,
	ctx context.Context,
	filter any,
	updateData any,
	opts ...*options.FindOneAndUpdateOptions,
) error {
	args := m.Called(v, ctx, filter, updateData, opts)
	return args.Error(0)
}

func (m *MockMongoDBService) UpdateOne(
	ctx context.Context,
	filter any,
	updateData any,
	opts ...*options.UpdateOptions,
) (int64, error) {
	args := m.Called(ctx, filter, updateData, opts)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockMongoDBService) UpdateMany(
	ctx context.Context,
	filter any,
	updateData any,
	opts ...*options.UpdateOptions,
) (int64, error) {
	args := m.Called(ctx, filter, updateData, opts)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockMongoDBService) FindOneAndDelete(
	v any,
	ctx context.Context,
	filter any,
	opts ...*options.FindOneAndDeleteOptions,
) error {
	args := m.Called(v, ctx, filter, opts)
	return args.Error(0)
}

func (m *MockMongoDBService) DeleteOne(
	ctx context.Context,
	filter any,
	opts ...*options.DeleteOptions,
) (int64, error) {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockMongoDBService) DeleteMany(
	ctx context.Context,
	filter any,
	opts ...*options.DeleteOptions,
) (int64, error) {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockMongoDBService) Aggregate(
	v any,
	ctx context.Context,
	pipeline mongo.Pipeline,
	opts ...*options.AggregateOptions,
) error {
	args := m.Called(v, ctx, pipeline, opts)
	return args.Error(0)
}

func (m *MockMongoDBService) MakeUnique(ctx context.Context, filter any) error {
	args := m.Called(ctx, filter)
	return args.Error(0)
}

func (m *MockMongoDBService) CountDocuments(ctx context.Context, filter any) (int64, error) {
	args := m.Mock.Called(ctx, filter)

	if args.Get(0) == nil {
		return 0, nil
	}

	return args.Get(0).(int64), args.Error(1)
}

func (m *MockMongoDBService) CreateIndex(
	ctx context.Context,
	fields bson.D,
	index *options.IndexOptions,
) error {
	args := m.Called(ctx, fields, index)
	return args.Error(0)
}

func (m *MockMongoDBService) AggregateExplain(
	ctx context.Context,
	pipeline mongo.Pipeline,
	opts ...*options.AggregateOptions,
) (bson.M, error) {
	args := m.Called(ctx, pipeline, opts)
	return args.Get(0).(bson.M), args.Error(1)
}

func (m *MockMongoDBService) Close() {
	// Pass
}
