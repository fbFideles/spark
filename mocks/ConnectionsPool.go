// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"
	database "spark/database"

	mock "github.com/stretchr/testify/mock"
)

// ConnectionsPool is an autogenerated mock type for the ConnectionsPool type
type ConnectionsPool struct {
	mock.Mock
}

// NewTransaction provides a mock function with given fields: ctx, readOnly
func (_m *ConnectionsPool) NewTransaction(ctx context.Context, readOnly bool) (database.Transaction, error) {
	ret := _m.Called(ctx, readOnly)

	var r0 database.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, bool) (database.Transaction, error)); ok {
		return rf(ctx, readOnly)
	}
	if rf, ok := ret.Get(0).(func(context.Context, bool) database.Transaction); ok {
		r0 = rf(ctx, readOnly)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, bool) error); ok {
		r1 = rf(ctx, readOnly)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewConnectionsPool interface {
	mock.TestingT
	Cleanup(func())
}

// NewConnectionsPool creates a new instance of ConnectionsPool. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConnectionsPool(t mockConstructorTestingTNewConnectionsPool) *ConnectionsPool {
	mock := &ConnectionsPool{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
