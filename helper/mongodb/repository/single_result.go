package repository

import "go.mongodb.org/mongo-driver/mongo"

type mongodbSingleResult struct {
	singleResult *mongo.SingleResult
}

func (msr *mongodbSingleResult) Decode(v interface{}) error {
	err := msr.singleResult.Decode(v)
	if err != nil {
		return err
	}

	return nil
}

func (msr *mongodbSingleResult) Err() error {
	return msr.singleResult.Err()
}
