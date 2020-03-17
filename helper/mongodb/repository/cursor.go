package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type mongodbCursor struct {
	cursor *mongo.Cursor
}

func (mc *mongodbCursor) All(ctx context.Context, results interface{}) error {
	err := mc.cursor.All(ctx, results)
	return err
}

func (mc *mongodbCursor) Decode(val interface{}) error {
	err := mc.cursor.Decode(val)
	return err
}

func (mc *mongodbCursor) Close(ctx context.Context) error {
	err := mc.cursor.Close(ctx)
	return err
}

func (mc *mongodbCursor) Err(ctx context.Context) error {
	err := mc.cursor.Err()
	return err
}
