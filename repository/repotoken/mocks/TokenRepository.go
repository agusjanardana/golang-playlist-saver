// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	record "playlist-saver/model/record"

	mock "github.com/stretchr/testify/mock"
)

// TokenRepository is an autogenerated mock type for the TokenRepository type
type TokenRepository struct {
	mock.Mock
}

// CheckToken provides a mock function with given fields: ctx, token
func (_m *TokenRepository) CheckToken(ctx context.Context, token string) (bool, error) {
	ret := _m.Called(ctx, token)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetToken provides a mock function with given fields: ctx
func (_m *TokenRepository) GetToken(ctx context.Context) ([]record.Token, error) {
	ret := _m.Called(ctx)

	var r0 []record.Token
	if rf, ok := ret.Get(0).(func(context.Context) []record.Token); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]record.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostToken provides a mock function with given fields: ctx, token
func (_m *TokenRepository) PostToken(ctx context.Context, token record.Token) error {
	ret := _m.Called(ctx, token)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, record.Token) error); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
