// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/silvergama/transations/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// AccountRepositoryInterface is an autogenerated mock type for the AccountRepositoryInterface type
type AccountRepositoryInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, account
func (_m *AccountRepositoryInterface) Create(_a0 context.Context, account *domain.Account) (int, error) {
	ret := _m.Called(_a0, account)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Account) (int, error)); ok {
		return rf(_a0, account)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Account) int); ok {
		r0 = rf(_a0, account)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *domain.Account) error); ok {
		r1 = rf(_a0, account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: _a0, id
func (_m *AccountRepositoryInterface) GetByID(_a0 context.Context, id int) (*domain.Account, error) {
	ret := _m.Called(_a0, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *domain.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*domain.Account, error)); ok {
		return rf(_a0, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *domain.Account); ok {
		r0 = rf(_a0, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(_a0, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAccountRepositoryInterface creates a new instance of AccountRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccountRepositoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *AccountRepositoryInterface {
	mock := &AccountRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}