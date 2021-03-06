// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mongodb "github.com/sangianpatrick/go-cqrs-demo/helper/mongodb"
	mock "github.com/stretchr/testify/mock"
)

// Database is an autogenerated mock type for the Database type
type Database struct {
	mock.Mock
}

// Collection provides a mock function with given fields: collectionName
func (_m *Database) Collection(collectionName string) mongodb.Collection {
	ret := _m.Called(collectionName)

	var r0 mongodb.Collection
	if rf, ok := ret.Get(0).(func(string) mongodb.Collection); ok {
		r0 = rf(collectionName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongodb.Collection)
		}
	}

	return r0
}
