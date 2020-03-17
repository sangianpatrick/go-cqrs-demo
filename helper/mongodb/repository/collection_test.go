package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/sangianpatrick/go-cqrs-demo/helper/mongodb/repository"
)

func TestClone(t *testing.T) {
	var uri string = "mongodb://localhost:27017/test"
	client, _ := repository.NewClient(uri)
	client.Connect(context.Background())
	db := client.Database("test")
	collection := db.Collection("haha")

	t.Run("should return new clone collection and error", func(t *testing.T) {
		ncol, err := collection.Clone()
		assert.NotNil(t, ncol)
		assert.NoError(t, err)
	})
}

func TestCountDocuments(t *testing.T) {
	var uri string = "mongodb://localhost:27017/test"
	client, _ := repository.NewClient(uri)
	client.Connect(context.Background())
	db := client.Database("test")
	collection := db.Collection("haha")

	t.Run("should return result of count documents", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*1)
		defer cancel()
		counted, err := collection.CountDocuments(ctx, bson.M{})
		assert.Zero(t, counted)
		assert.Error(t, err)
	})
}

func TestFindOne(t *testing.T) {
	var uri string = "mongodb://localhost:27017/test"
	client, _ := repository.NewClient(uri)
	client.Connect(context.Background())
	db := client.Database("test")
	collection := db.Collection("haha")

	t.Run("should return single result", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*1)
		defer cancel()
		sr := collection.FindOne(ctx, bson.M{})
		assert.NotNil(t, sr)
	})
}

func TestFind(t *testing.T) {
	var uri string = "mongodb://localhost:27017/test"
	client, _ := repository.NewClient(uri)
	client.Connect(context.Background())
	db := client.Database("test")
	collection := db.Collection("haha")

	t.Run("should return cursor", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*1)
		defer cancel()
		cur, err := collection.Find(ctx, bson.M{})
		assert.Error(t, err)
		assert.NotNil(t, cur)
	})
}
