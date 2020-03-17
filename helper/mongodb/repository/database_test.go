package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sangianpatrick/go-cqrs-demo/helper/mongodb/repository"
)

func TestCollection(t *testing.T) {
	var uri string = "mongodb://localhost:27017/test"
	client, _ := repository.NewClient(uri)
	client.Connect(context.Background())
	db := client.Database("test")
	collection := db.Collection("haha")

	assert.NotNil(t, collection)
}
