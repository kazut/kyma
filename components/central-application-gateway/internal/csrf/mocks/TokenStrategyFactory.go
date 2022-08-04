// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	csrf "github.com/kyma-project/kyma/components/central-application-gateway/internal/csrf"
	authorization "github.com/kyma-project/kyma/components/central-application-gateway/pkg/authorization"

	mock "github.com/stretchr/testify/mock"
)

// TokenStrategyFactory is an autogenerated mock type for the TokenStrategyFactory type
type TokenStrategyFactory struct {
	mock.Mock
}

// Create provides a mock function with given fields: authorizationStrategy, csrfTokenEndpointURL
func (_m *TokenStrategyFactory) Create(authorizationStrategy authorization.Strategy, csrfTokenEndpointURL string) csrf.TokenStrategy {
	ret := _m.Called(authorizationStrategy, csrfTokenEndpointURL)

	var r0 csrf.TokenStrategy
	if rf, ok := ret.Get(0).(func(authorization.Strategy, string) csrf.TokenStrategy); ok {
		r0 = rf(authorizationStrategy, csrfTokenEndpointURL)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(csrf.TokenStrategy)
		}
	}

	return r0
}

type mockConstructorTestingTNewTokenStrategyFactory interface {
	mock.TestingT
	Cleanup(func())
}

// NewTokenStrategyFactory creates a new instance of TokenStrategyFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTokenStrategyFactory(t mockConstructorTestingTNewTokenStrategyFactory) *TokenStrategyFactory {
	mock := &TokenStrategyFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
