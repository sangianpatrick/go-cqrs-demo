package repository

import (
	"context"

	"github.com/sangianpatrick/go-cqrs-demo/helper/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongodbCollection struct {
	collection *mongo.Collection
}

func (mc *mongodbCollection) Clone() (mongodb.Collection, error) {
	cloneCollection, err := mc.collection.Clone()

	newmc := &mongodbCollection{
		collection: cloneCollection,
	}

	return newmc, err
}

func (mc *mongodbCollection) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	counted, err := mc.collection.CountDocuments(ctx, filter, opts...)
	return counted, err
}

func (mc *mongodbCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	singleResult := mc.collection.FindOne(ctx, filter, opts...)

	return singleResult
}

func (mc *mongodbCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	cursor, err := mc.collection.Find(ctx, filter, opts...)

	return cursor, err
}

func (mc *mongodbCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	insertOneResult, err := mc.collection.InsertOne(ctx, document, opts...)

	return insertOneResult, err
}

func (mc *mongodbCollection) InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	insertManyResult, err := mc.collection.InsertMany(ctx, documents, opts...)

	return insertManyResult, err
}

func (mc *mongodbCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	updateResult, err := mc.collection.UpdateOne(ctx, filter, update, opts...)

	return updateResult, err
}

func (mc *mongodbCollection) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	updateResult, err := mc.collection.UpdateMany(ctx, filter, update, opts...)

	return updateResult, err
}

func (mc *mongodbCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	deleteResult, err := mc.collection.DeleteOne(ctx, filter, opts...)

	return deleteResult, err
}

func (mc *mongodbCollection) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	deleteResult, err := mc.collection.DeleteMany(ctx, filter, opts...)

	return deleteResult, err
}

func (mc *mongodbCollection) Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (*mongo.Cursor, error) {
	aggsResult, err := mc.collection.Aggregate(ctx, pipeline, opts...)

	return aggsResult, err
}
