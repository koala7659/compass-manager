// Code generated by mockery v2.36.1. DO NOT EDIT.

package mocks

import (
	apperrors "github.com/kyma-project/compass-manager/internal/apperrors"

	gqlschema "github.com/kyma-project/compass-manager/pkg/gqlschema"

	graphql "github.com/kyma-incubator/compass/components/director/pkg/graphql"

	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// CreateRuntime provides a mock function with given fields: config, globalAccount
func (_m *Client) CreateRuntime(config *gqlschema.RuntimeInput, globalAccount string) (string, apperrors.AppError) {
	ret := _m.Called(config, globalAccount)

	var r0 string
	var r1 apperrors.AppError
	if rf, ok := ret.Get(0).(func(*gqlschema.RuntimeInput, string) (string, apperrors.AppError)); ok {
		return rf(config, globalAccount)
	}
	if rf, ok := ret.Get(0).(func(*gqlschema.RuntimeInput, string) string); ok {
		r0 = rf(config, globalAccount)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*gqlschema.RuntimeInput, string) apperrors.AppError); ok {
		r1 = rf(config, globalAccount)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}

// DeleteRuntime provides a mock function with given fields: compassID, globalAccount
func (_m *Client) DeleteRuntime(compassID string, globalAccount string) apperrors.AppError {
	ret := _m.Called(compassID, globalAccount)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, string) apperrors.AppError); ok {
		r0 = rf(compassID, globalAccount)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// GetConnectionToken provides a mock function with given fields: compassID, globalAccount
func (_m *Client) GetConnectionToken(compassID string, globalAccount string) (graphql.OneTimeTokenForRuntimeExt, apperrors.AppError) {
	ret := _m.Called(compassID, globalAccount)

	var r0 graphql.OneTimeTokenForRuntimeExt
	var r1 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, string) (graphql.OneTimeTokenForRuntimeExt, apperrors.AppError)); ok {
		return rf(compassID, globalAccount)
	}
	if rf, ok := ret.Get(0).(func(string, string) graphql.OneTimeTokenForRuntimeExt); ok {
		r0 = rf(compassID, globalAccount)
	} else {
		r0 = ret.Get(0).(graphql.OneTimeTokenForRuntimeExt)
	}

	if rf, ok := ret.Get(1).(func(string, string) apperrors.AppError); ok {
		r1 = rf(compassID, globalAccount)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}

// GetRuntime provides a mock function with given fields: compassID, globalAccount
func (_m *Client) GetRuntime(compassID string, globalAccount string) (graphql.RuntimeExt, apperrors.AppError) {
	ret := _m.Called(compassID, globalAccount)

	var r0 graphql.RuntimeExt
	var r1 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, string) (graphql.RuntimeExt, apperrors.AppError)); ok {
		return rf(compassID, globalAccount)
	}
	if rf, ok := ret.Get(0).(func(string, string) graphql.RuntimeExt); ok {
		r0 = rf(compassID, globalAccount)
	} else {
		r0 = ret.Get(0).(graphql.RuntimeExt)
	}

	if rf, ok := ret.Get(1).(func(string, string) apperrors.AppError); ok {
		r1 = rf(compassID, globalAccount)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
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
