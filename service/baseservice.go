package service

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Base Service Interface
type Service interface {
	Find(any, context.Context, interface{}, ...*options.FindOptions) error
	FindOne(any, context.Context, interface{}, ...*options.FindOneOptions) error
	InsertOne(context.Context, interface{}, ...*options.InsertOneOptions) (interface{}, error)
	InsertMany(context.Context, []interface{}, ...*options.InsertManyOptions) ([]interface{}, error)
	FindOneAndUpdate(any, context.Context, interface{}, interface{}, ...*options.FindOneAndUpdateOptions) error
	UpdateOne(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (int64, error)
	UpdateMany(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (int64, error)
	FindOneAndDelete(any, context.Context, interface{}, ...*options.FindOneAndDeleteOptions) error
	DeleteOne(context.Context, interface{}, ...*options.DeleteOptions) (int64, error)
	DeleteMany(context.Context, interface{}, ...*options.DeleteOptions) (int64, error)
	Aggregate(any, context.Context, mongo.Pipeline) error
}

type BaseService struct {
	collection *mongo.Collection
}

func NewBaseService(collection *mongo.Collection) *BaseService {
	return &BaseService{
		collection: collection,
	}
}

func (s *BaseService) Find(v any, ctx context.Context, filter interface{},
	opts ...*options.FindOptions) error {
	cursor, err := s.collection.Find(ctx, filter, opts...)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, v)
	if err != nil {
		return err
	}

	return nil
}

func (s *BaseService) FindOne(v any, ctx context.Context, filter interface{},
	opts ...*options.FindOneOptions) error {
	result := s.collection.FindOne(ctx, filter, opts...)
	if result.Err() != nil {
		return result.Err()
	}

	err := result.Decode(v)
	if err != nil {
		return err
	}

	return nil
}

func (s *BaseService) InsertOne(ctx context.Context, data interface{},
	opts ...*options.InsertOneOptions) (interface{}, error) {
	result, err := s.collection.InsertOne(ctx, data, opts...)
	if err != nil {
		return "", err
	}

	return result.InsertedID, nil
}

func (s *BaseService) InsertMany(ctx context.Context, data []interface{},
	opts ...*options.InsertManyOptions) ([]interface{}, error) {
	result, err := s.collection.InsertMany(ctx, data, opts...)

	if err != nil {
		return nil, err
	}

	return result.InsertedIDs, nil
}

func (s *BaseService) FindOneAndUpdate(v any, ctx context.Context, filter interface{}, data interface{},
	opts ...*options.FindOneAndUpdateOptions) error {
	result := s.collection.FindOneAndUpdate(ctx, filter, data, opts...)
	if result.Err() != nil {
		return result.Err()
	}

	err := result.Decode(v)
	if err != nil {
		return err
	}

	return nil
}

func (s *BaseService) UpdateOne(ctx context.Context, filter interface{}, data interface{},
	opts ...*options.UpdateOptions) (int64, error) {
	result, err := s.collection.UpdateOne(ctx, filter, data, opts...)
	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}

func (s *BaseService) UpdateMany(ctx context.Context, filter interface{}, data interface{},
	opts ...*options.UpdateOptions) (int64, error) {
	result, err := s.collection.UpdateMany(ctx, filter, data, opts...)
	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}

func (s *BaseService) FindOneAndDelete(v any, ctx context.Context, filter interface{},
	opts ...*options.FindOneAndDeleteOptions) error {
	result := s.collection.FindOneAndDelete(ctx, filter, opts...)
	if result.Err() != nil {
		return result.Err()
	}

	err := result.Decode(v)
	if err != nil {
		return err
	}

	return nil
}

func (s *BaseService) DeleteOne(ctx context.Context, filter interface{},
	opts ...*options.DeleteOptions) (int64, error) {
	result, err := s.collection.DeleteOne(ctx, filter, opts...)
	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}

func (s *BaseService) DeleteMany(ctx context.Context, filter interface{},
	opts ...*options.DeleteOptions) (int64, error) {
	result, err := s.collection.DeleteMany(ctx, filter, opts...)
	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}

func (s *BaseService) Aggregate(v any, ctx context.Context, pipeline mongo.Pipeline) error {
	cursor, err := s.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, v)
	if err != nil {
		return err
	}

	return nil
}
