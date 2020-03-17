package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sangianpatrick/go-cqrs-demo/helper/mongodb/repository"
)

func TestNewClient(t *testing.T) {

	t.Run("should return client and error", func(t *testing.T) {
		var uri = "mongodb://localhost:27018/test"
		_, err := repository.NewClient(uri)
		assert.NoError(t, err)
	})
}

func TestConnect(t *testing.T) {
	t.Run("should run connection and return nil error", func(t *testing.T) {
		var uri = "mongodb://localhost:27018/test"
		client, _ := repository.NewClient(uri)
		err := client.Connect(context.Background())
		assert.NoError(t, err)
	})
}
