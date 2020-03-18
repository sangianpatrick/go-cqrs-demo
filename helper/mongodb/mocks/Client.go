// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mongodb "github.com/sangianpatrick/go-cqrs-demo/helper/mongodb"
	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Connect provides a mock function with given fields: ctx
func (_m *Client) Connect(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Database provides a mock function with given fields: dbName
func (_m *Client) Database(dbName string) mongodb.Database {
	ret := _m.Called(dbName)

	var r0 mongodb.Database
	if rf, ok := ret.Get(0).(func(string) mongodb.Database); ok {
		r0 = rf(dbName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongodb.Database)
		}
	}

	return r0
}
