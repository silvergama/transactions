// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"

	account "github.com/silvergama/transations/account"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *Repository) Create(_a0 context.Context, _a1 *account.Account) (int, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *account.Account) (int, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *account.Account) int); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *account.Account) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: _a0, id
func (_m *Repository) GetByID(_a0 context.Context, id int) (*account.Account, error) {
	ret := _m.Called(_a0, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *account.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*account.Account, error)); ok {
		return rf(_a0, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *account.Account); ok {
		r0 = rf(_a0, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*account.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(_a0, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
