package repository

import (
	"context"

	"github.com/sangianpatrick/go-cqrs-demo/helper/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongodbClient struct {
	client *mongo.Client
}

// NewClient is mongodb Client constructor
func NewClient(uri string) (mongodb.Client, error) {

	opts := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(opts)

	if err != nil {
		return nil, err
	}

	return &mongodbClient{
		client: client,
	}, nil
}

func (mc *mongodbClient) Connect(ctx context.Context) error {
	err := mc.client.Connect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (mc *mongodbClient) Database(dbName string) mongodb.Database {
	d := mc.client.Database(dbName)
	md := &mongodbDatabase{
		database: d,
	}

	return md
}
