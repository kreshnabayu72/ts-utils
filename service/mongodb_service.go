package service

import (
	"context"
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/v2/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDBServicer provides an interface for implementing MongoDBService.
type MongoDBServicer interface {
	Init()
	Find(
		v any,
		ctx context.Context,
		filter any,
		options ...*options.FindOptions,
	) error
	FindOne(
		v any,
		ctx context.Context,
		filter any,
		options ...*options.FindOneOptions,
	) error
	InsertOne(
		ctx context.Context,
		data any,
		options ...*options.InsertOneOptions,
	) (any, error)
	InsertMany(
		ctx context.Context,
		data []any,
		options ...*options.InsertManyOptions,
	) ([]any, error)
	FindOneAndUpdate(
		v any,
		ctx context.Context,
		filter any,
		updateData any,
		options ...*options.FindOneAndUpdateOptions,
	) error
	UpdateOne(
		ctx context.Context,
		filter any,
		updateData any,
		options ...*options.UpdateOptions,
	) (int64, error)
	UpdateMany(
		ctx context.Context,
		filter any,
		updateData any,
		options ...*options.UpdateOptions,
	) (int64, error)
	FindOneAndDelete(
		v any,
		ctx context.Context,
		filter any,
		options ...*options.FindOneAndDeleteOptions,
	) error
	DeleteOne(
		ctx context.Context,
		filter any,
		options ...*options.DeleteOptions,
	) (int64, error)
	DeleteMany(
		ctx context.Context,
		filter any,
		options ...*options.DeleteOptions,
	) (int64, error)
	Aggregate(
		v any,
		ctx context.Context,
		pipeline mongo.Pipeline,
		options ...*options.AggregateOptions,
	) error
	AggregateExplain(
		ctx context.Context,
		pipeline mongo.Pipeline,
		opts ...*options.AggregateOptions,
	) (bson.M, error)
	MakeUnique(
		ctx context.Context,
		filter any,
	) error
	CountDocuments(
		ctx context.Context,
		filter any,
	) (int64, error)
	CreateIndex(
		ctx context.Context,
		fields bson.D,
		index *options.IndexOptions,
	) error
	Close()
}

// MongoDBService implements MongoDBServicer
//
// Full implementation of CRUD operation for MongoDB.
type MongoDBService struct {
	collection *mongo.Collection
}

// MakeMongoDBService creates a new MongoDBService instance.
func MakeMongoDBService(dbName db.DBName, collection string) *MongoDBService {
	db.ConnectToDB(dbName)
	coll := db.GetCollection(collection)

	return &MongoDBService{
		collection: coll,
	}
}

func (s *MongoDBService) Init() {
	// Pass
}

func (s *MongoDBService) Find(v any, ctx context.Context, filter any,
	opts ...*options.FindOptions,
) error {
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

func (s *MongoDBService) FindOne(v any, ctx context.Context, filter any,
	opts ...*options.FindOneOptions,
) error {
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

func (s *MongoDBService) InsertOne(ctx context.Context, data any,
	opts ...*options.InsertOneOptions,
) (any, error) {

	result, err := s.collection.InsertOne(ctx, data, opts...)
	if err != nil {
		return "", err
	}

	// Injecting updatedAt and createdAt on the insert function call to always gives timestamp to the field.
	updatedAt := bson.M{
		"$set": bson.M{
			"updatedAt": time.Now(),
			"createdAt": time.Now(),
		},
	}
	_, err = s.collection.UpdateOne(ctx, bson.M{
		"_id": result.InsertedID,
	}, updatedAt)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func (s *MongoDBService) InsertMany(ctx context.Context, data []any,
	opts ...*options.InsertManyOptions,
) ([]any, error) {
	result, err := s.collection.InsertMany(ctx, data, opts...)
	if err != nil {
		return nil, err
	}

	// Injecting updatedAt and createdAt on the insert function call to always gives timestamp to the field.
	updatedAt := bson.M{
		"$set": bson.M{
			"updatedAt": time.Now(),
			"createdAt": time.Now(),
		},
	}
	_, err = s.collection.UpdateMany(ctx, bson.M{
		"_id": bson.M{
			"$in": result.InsertedIDs,
		},
	}, updatedAt)
	if err != nil {
		return nil, err
	}

	return result.InsertedIDs, nil
}

func (s *MongoDBService) FindOneAndUpdate(
	v any,
	ctx context.Context,
	filter any,
	data any,
	opts ...*options.FindOneAndUpdateOptions,
) error {
	result := s.collection.FindOneAndUpdate(ctx, filter, data, opts...)
	if result.Err() != nil {
		return result.Err()
	}

	err := result.Decode(v)
	if err != nil {
		return err
	}

	// Injecting updatedAt on the update function call to always update the updatedAt field.
	updatedAt := bson.M{
		"$set": bson.M{
			"updatedAt": time.Now(),
		},
	}
	_, err = s.collection.UpdateOne(ctx, filter, updatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (s *MongoDBService) UpdateOne(ctx context.Context, filter any, data any,
	opts ...*options.UpdateOptions,
) (int64, error) {

	result, err := s.collection.UpdateOne(ctx, filter, data, opts...)
	if err != nil {
		return 0, err
	}

	// Injecting updatedAt on the update function call to always update the updatedAt field.
	updatedAt := bson.M{
		"$set": bson.M{
			"updatedAt": time.Now(),
		},
	}
	_, err = s.collection.UpdateOne(ctx, filter, updatedAt)
	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}

func (s *MongoDBService) UpdateMany(ctx context.Context, filter any, data any,
	opts ...*options.UpdateOptions,
) (int64, error) {
	result, err := s.collection.UpdateMany(ctx, filter, data, opts...)
	if err != nil {
		return 0, err
	}

	// Injecting updatedAt on the update function call to always update the updatedAt field.
	updatedAt := bson.M{
		"$set": bson.M{
			"updatedAt": time.Now(),
		},
	}
	_, err = s.collection.UpdateMany(ctx, filter, updatedAt)
	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}

func (s *MongoDBService) FindOneAndDelete(v any, ctx context.Context, filter any,
	opts ...*options.FindOneAndDeleteOptions,
) error {
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

func (s *MongoDBService) DeleteOne(ctx context.Context, filter any,
	opts ...*options.DeleteOptions,
) (int64, error) {
	result, err := s.collection.DeleteOne(ctx, filter, opts...)
	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}

func (s *MongoDBService) DeleteMany(ctx context.Context, filter any,
	opts ...*options.DeleteOptions,
) (int64, error) {
	result, err := s.collection.DeleteMany(ctx, filter, opts...)
	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}

func (s *MongoDBService) Aggregate(
	v any,
	ctx context.Context,
	pipeline mongo.Pipeline,
	opts ...*options.AggregateOptions,
) error {
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

func (s *MongoDBService) AggregateExplain(
	ctx context.Context,
	pipeline mongo.Pipeline,
	opts ...*options.AggregateOptions,
) (bson.M, error) {
	// Build the aggregate command
	command := bson.D{
		{Key: "explain", Value: bson.D{
			{Key: "aggregate", Value: s.collection.Name()},
			{Key: "pipeline", Value: pipeline},
			{Key: "cursor", Value: bson.D{}},
		}},
	}

	// Add options if needed (e.g., allowDiskUse)
	if len(opts) > 0 && opts[0] != nil && opts[0].AllowDiskUse != nil {
		command[0].Value = append(
			command[0].Value.(bson.D),
			bson.E{Key: "allowDiskUse", Value: *opts[0].AllowDiskUse},
		)
	}

	// Run the explain command
	var result bson.M
	err := s.collection.Database().RunCommand(ctx, command).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *MongoDBService) MakeUnique(ctx context.Context, filter any) error {
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

func (s *MongoDBService) CountDocuments(ctx context.Context, filter any) (int64, error) {
	documentCounts, err := s.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return documentCounts, nil
}

func (s *MongoDBService) CreateIndex(
	ctx context.Context,
	fields bson.D,
	index *options.IndexOptions,
) error {
	indexModel := mongo.IndexModel{
		Keys:    fields,
		Options: index,
	}

	_, err := s.collection.Indexes().CreateOne(ctx, indexModel)

	return err
}

func (s *MongoDBService) Close() {
	// Pass
}
