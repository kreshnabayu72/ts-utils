package service

import (
	"context"

	"github.com/betsv2/tsv2-go-utils/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Base Service Interface
type Service interface {
	Find(v any, ctx context.Context, filter interface{}, options ...*options.FindOptions) error
	FindOne(v any, ctx context.Context, filter interface{}, options ...*options.FindOneOptions) error
	InsertOne(ctx context.Context, data interface{}, options ...*options.InsertOneOptions) (interface{}, error)
	InsertMany(ctx context.Context, data []interface{}, options ...*options.InsertManyOptions) ([]interface{}, error)
	FindOneAndUpdate(v any, ctx context.Context, filter interface{}, updateData interface{}, options ...*options.FindOneAndUpdateOptions) error
	UpdateOne(ctx context.Context, filter interface{}, updateData interface{}, options ...*options.UpdateOptions) (int64, error)
	UpdateMany(ctx context.Context, fitler interface{}, updateData interface{}, options ...*options.UpdateOptions) (int64, error)
	FindOneAndDelete(v any, ctx context.Context, filter interface{}, options ...*options.FindOneAndDeleteOptions) error
	DeleteOne(ctx context.Context, filter interface{}, options ...*options.DeleteOptions) (int64, error)
	DeleteMany(ctx context.Context, filter interface{}, options ...*options.DeleteOptions) (int64, error)
	Aggregate(v any, ctx context.Context, pipeline mongo.Pipeline, options ...*options.AggregateOptions) error
	MakeUnique(ctx context.Context, filter interface{}) error
}

type BaseService struct {
	collection *mongo.Collection
}

func MakeService(companyCode string, collection string) *BaseService {
	db.ConnectToCompanyDb(companyCode)
	coll := db.GetCollection(collection)

	return &BaseService{
		collection: coll,
	}
}

// Deprecated: Use `MakeService` instead with company code and collection name.
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

func (s *BaseService) Aggregate(v any, ctx context.Context, pipeline mongo.Pipeline, opts ...*options.AggregateOptions) error {
	cursor, err := s.collection.Aggregate(ctx, pipeline, opts...)
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

func (s *BaseService) MakeUnique(ctx context.Context, filter interface{}) error {
	indexModel := mongo.IndexModel{
		Keys:    filter,
		Options: options.Index().SetUnique(true),
	}

	_, err := s.collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return err
	}

	return nil
}
