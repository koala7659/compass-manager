// Code generated by mockery v2.36.1. DO NOT EDIT.

package mocks

import (
	graphql "github.com/kyma-project/compass-manager/third_party/machinebox/graphql"

	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Do provides a mock function with given fields: req, res, gracefulUnregistration
func (_m *Client) Do(req *graphql.Request, res interface{}, gracefulUnregistration bool) error {
	ret := _m.Called(req, res, gracefulUnregistration)

	var r0 error
	if rf, ok := ret.Get(0).(func(*graphql.Request, interface{}, bool) error); ok {
		r0 = rf(req, res, gracefulUnregistration)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
