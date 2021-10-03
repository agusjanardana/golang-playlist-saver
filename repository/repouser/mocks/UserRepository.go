// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	record "playlist-saver/model/record"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// GetAllUser provides a mock function with given fields: ctx
func (_m *UserRepository) GetAllUser(ctx context.Context) ([]record.User, error) {
	ret := _m.Called(ctx)

	var r0 []record.User
	if rf, ok := ret.Get(0).(func(context.Context) []record.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]record.User)
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

// Login provides a mock function with given fields: ctx, email
func (_m *UserRepository) Login(ctx context.Context, email string) (record.User, error) {
	ret := _m.Called(ctx, email)

	var r0 record.User
	if rf, ok := ret.Get(0).(func(context.Context, string) record.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(record.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: ctx, user
func (_m *UserRepository) Register(ctx context.Context, user record.User) (record.User, error) {
	ret := _m.Called(ctx, user)

	var r0 record.User
	if rf, ok := ret.Get(0).(func(context.Context, record.User) record.User); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(record.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, record.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: ctx, user, id
func (_m *UserRepository) UpdateUser(ctx context.Context, user record.User, id int) error {
	ret := _m.Called(ctx, user, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, record.User, int) error); ok {
		r0 = rf(ctx, user, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserAddToken provides a mock function with given fields: ctx, id, token
func (_m *UserRepository) UserAddToken(ctx context.Context, id int, token int) error {
	ret := _m.Called(ctx, id, token)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) error); ok {
		r0 = rf(ctx, id, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserFindById provides a mock function with given fields: ctx, id
func (_m *UserRepository) UserFindById(ctx context.Context, id int) (record.User, error) {
	ret := _m.Called(ctx, id)

	var r0 record.User
	if rf, ok := ret.Get(0).(func(context.Context, int) record.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(record.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
