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
		assert.Nil(t, cur)
	})
}

func TestInsertOne(t *testing.T) {
	var uri string = "mongodb://localhost:27017/test"
	client, _ := repository.NewClient(uri)
	client.Connect(context.Background())
	db := client.Database("test")
	collection := db.Collection("haha")
	data := make(map[string]interface{})
	data["name"] = "test"

	t.Run("should return InsertResult and error", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*1)
		defer cancel()
		ioResult, err := collection.InsertOne(ctx, data)
		assert.Error(t, err)
		assert.Nil(t, ioResult)
	})
}

func TestInsertMany(t *testing.T) {
	var uri string = "mongodb://localhost:27017/test"
	client, _ := repository.NewClient(uri)
	client.Connect(context.Background())
	db := client.Database("test")
	collection := db.Collection("haha")
	list := make([]interface{}, 1)
	data := map[string]interface{}{
		"name": "test",
	}

	list[0] = data

	t.Run("should return InsertResult and error", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*1)
		defer cancel()
		ioResult, err := collection.InsertMany(ctx, list)
		assert.Error(t, err)
		assert.Nil(t, ioResult)
	})
}

func TestUpdateOne(t *testing.T) {
	var uri string = "mongodb://localhost:27017/test"
	client, _ := repository.NewClient(uri)
	client.Connect(context.Background())
	db := client.Database("test")
	collection := db.Collection("haha")

	update := map[string]interface{}{
		"name": "test",
	}

	t.Run("should return UpdateResult and error", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*1)
		defer cancel()
		uResult, err := collection.UpdateOne(ctx, bson.M{}, update)
		assert.Error(t, err)
		assert.Nil(t, uResult)
	})
}

func TestUpdateMany(t *testing.T) {
	var uri string = "mongodb://localhost:27017/test"
	client, _ := repository.NewClient(uri)
	client.Connect(context.Background())
	db := client.Database("test")
	collection := db.Collection("haha")

	update := map[string]interface{}{
		"name": "test",
	}

	t.Run("should return UpdateResult and error", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*1)
		defer cancel()
		uResult, err := collection.UpdateMany(ctx, bson.M{}, update)
		assert.Error(t, err)
		assert.Nil(t, uResult)
	})
}

func TestDeleteOne(t *testing.T) {
	var uri string = "mongodb://localhost:27017/test"
	client, _ := repository.NewClient(uri)
	client.Connect(context.Background())
	db := client.Database("test")
	collection := db.Collection("haha")

	t.Run("should return DeleteResult and error", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*1)
		defer cancel()
		dResult, err := collection.DeleteOne(ctx, bson.M{})
		assert.Error(t, err)
		assert.Nil(t, dResult)
	})
}

func TestDeleteMany(t *testing.T) {
	var uri string = "mongodb://localhost:27017/test"
	client, _ := repository.NewClient(uri)
	client.Connect(context.Background())
	db := client.Database("test")
	collection := db.Collection("haha")

	t.Run("should return DeleteResult and error", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*1)
		defer cancel()
		dResult, err := collection.DeleteMany(ctx, bson.M{})
		assert.Error(t, err)
		assert.Nil(t, dResult)
	})
}

func TestAggregate(t *testing.T) {
	var uri string = "mongodb://localhost:27017/test"
	client, _ := repository.NewClient(uri)
	client.Connect(context.Background())
	db := client.Database("test")
	collection := db.Collection("haha")

	t.Run("should return cursor", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*1)
		defer cancel()
		cur, err := collection.Aggregate(ctx, bson.M{})
		assert.Error(t, err)
		assert.Nil(t, cur)
	})
}
