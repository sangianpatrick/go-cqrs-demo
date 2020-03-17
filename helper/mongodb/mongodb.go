package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client is mongodb's client collection of function
type Client interface {
	Connect(ctx context.Context) error
	Database(dbName string) Database
}

// Database is mongodb's database collection of function
type Database interface {
	Collection(collectionName string) Collection
}

// Collection is mongodb's collection collection of function
type Collection interface {
	Clone() (Collection, error)
	CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error)
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) SingleResult
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (Cursor, error)
}

// SingleResult is mongodb's single result collection of function
type SingleResult interface {
	Decode(v interface{}) error
	Err() error
}

// Cursor is mongodb's cursor collection of function
type Cursor interface {
	All(ctx context.Context, results interface{}) error
	Decode(val interface{}) error
	Close(ctx context.Context) error
	Err(ctx context.Context) error
}
