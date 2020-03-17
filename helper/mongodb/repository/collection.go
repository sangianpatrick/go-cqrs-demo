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
	if err != nil {
		return nil, err
	}

	newmc := &mongodbCollection{
		collection: cloneCollection,
	}

	return newmc, err
}

func (mc *mongodbCollection) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	counted, err := mc.collection.CountDocuments(ctx, filter, opts...)
	return counted, err
}

func (mc *mongodbCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) mongodb.SingleResult {
	singleResult := mc.collection.FindOne(ctx, filter, opts...)

	msr := &mongodbSingleResult{
		singleResult: singleResult,
	}

	return msr
}

func (mc *mongodbCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (mongodb.Cursor, error) {
	cursor, err := mc.collection.Find(ctx, filter, opts...)

	cur := &mongodbCursor{
		cursor: cursor,
	}

	return cur, err
}
