package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sangianpatrick/go-cqrs-demo/helper/mongodb/repository"
)

func TestNewClient(t *testing.T) {
	t.Run("should return error while get new instance of mongodb client", func(t *testing.T) {
		var uri = "xxxx"
		_, err := repository.NewClient(uri)
		assert.Error(t, err)
	})

	t.Run("should return success", func(t *testing.T) {
		var uri = "mongodb://localhost:27018/test"
		_, err := repository.NewClient(uri)
		assert.NoError(t, err)
	})
}

func TestConnect(t *testing.T) {
	t.Run("should return success while connect to mongodb server", func(t *testing.T) {
		var uri = "mongodb://localhost:27018/test"
		client, _ := repository.NewClient(uri)
		err := client.Connect(context.Background())
		assert.NoError(t, err)
	})
}
