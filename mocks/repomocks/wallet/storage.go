// Code generated by mockery v2.40.1. DO NOT EDIT.

package repomocks

import (
	sql "database/sql"

	mock "github.com/stretchr/testify/mock"

	wallet "wallet/storage/wallet"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: w
func (_m *Repository) Create(w *wallet.Wallet) error {
	ret := _m.Called(w)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*wallet.Wallet) error); ok {
		r0 = rf(w)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *Repository) Delete(id int64) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByMemberID provides a mock function with given fields: memberID
func (_m *Repository) DeleteByMemberID(memberID int64) error {
	ret := _m.Called(memberID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByMemberID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(memberID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: id
func (_m *Repository) GetByID(id int64) (*wallet.Wallet, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *wallet.Wallet
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (*wallet.Wallet, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int64) *wallet.Wallet); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*wallet.Wallet)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByMemberID provides a mock function with given fields: memberID
func (_m *Repository) GetByMemberID(memberID int64) ([]*wallet.Wallet, error) {
	ret := _m.Called(memberID)

	if len(ret) == 0 {
		panic("no return value specified for GetByMemberID")
	}

	var r0 []*wallet.Wallet
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) ([]*wallet.Wallet, error)); ok {
		return rf(memberID)
	}
	if rf, ok := ret.Get(0).(func(int64) []*wallet.Wallet); ok {
		r0 = rf(memberID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*wallet.Wallet)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(memberID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBalance provides a mock function with given fields: id, balance
func (_m *Repository) UpdateBalance(id int64, balance int64) error {
	ret := _m.Called(id, balance)

	if len(ret) == 0 {
		panic("no return value specified for UpdateBalance")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, int64) error); ok {
		r0 = rf(id, balance)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WithTX provides a mock function with given fields: tx
func (_m *Repository) WithTX(tx *sql.Tx) (wallet.Repository, error) {
	ret := _m.Called(tx)

	if len(ret) == 0 {
		panic("no return value specified for WithTX")
	}

	var r0 wallet.Repository
	var r1 error
	if rf, ok := ret.Get(0).(func(*sql.Tx) (wallet.Repository, error)); ok {
		return rf(tx)
	}
	if rf, ok := ret.Get(0).(func(*sql.Tx) wallet.Repository); ok {
		r0 = rf(tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(wallet.Repository)
		}
	}

	if rf, ok := ret.Get(1).(func(*sql.Tx) error); ok {
		r1 = rf(tx)
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
