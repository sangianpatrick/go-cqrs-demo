package repository

import (
	"github.com/sangianpatrick/go-cqrs-demo/helper/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongodbDatabase struct {
	database *mongo.Database
}

func (md *mongodbDatabase) Collection(collectionName string) mongodb.Collection {
	collection := md.database.Collection(collectionName)

	mc := &mongodbCollection{
		collection: collection,
	}

	return mc
}
