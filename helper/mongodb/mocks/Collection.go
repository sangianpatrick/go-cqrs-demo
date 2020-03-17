// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mongodb "github.com/sangianpatrick/go-cqrs-demo/helper/mongodb"
	mock "github.com/stretchr/testify/mock"

	options "go.mongodb.org/mongo-driver/mongo/options"
)

// Collection is an autogenerated mock type for the Collection type
type Collection struct {
	mock.Mock
}

// Clone provides a mock function with given fields:
func (_m *Collection) Clone() (mongodb.Collection, error) {
	ret := _m.Called()

	var r0 mongodb.Collection
	if rf, ok := ret.Get(0).(func() mongodb.Collection); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongodb.Collection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountDocuments provides a mock function with given fields: ctx, filter, opts
func (_m *Collection) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.CountOptions) int64); ok {
		r0 = rf(ctx, filter, opts...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...*options.CountOptions) error); ok {
		r1 = rf(ctx, filter, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOne provides a mock function with given fields: ctx, filter, opts
func (_m *Collection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) mongodb.SingleResult {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 mongodb.SingleResult
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.FindOneOptions) mongodb.SingleResult); ok {
		r0 = rf(ctx, filter, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongodb.SingleResult)
		}
	}

	return r0
}
