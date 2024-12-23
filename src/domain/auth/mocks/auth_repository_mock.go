// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	entity "HRSystem/src/entity"

	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"
)

// MockAuthRepository is an autogenerated mock type for the AuthRepository type
type MockAuthRepository struct {
	mock.Mock
}

type MockAuthRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAuthRepository) EXPECT() *MockAuthRepository_Expecter {
	return &MockAuthRepository_Expecter{mock: &_m.Mock}
}

// CreateUser provides a mock function with given fields: user
func (_m *MockAuthRepository) CreateUser(user entity.User) error {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAuthRepository_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type MockAuthRepository_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - user entity.User
func (_e *MockAuthRepository_Expecter) CreateUser(user interface{}) *MockAuthRepository_CreateUser_Call {
	return &MockAuthRepository_CreateUser_Call{Call: _e.mock.On("CreateUser", user)}
}

func (_c *MockAuthRepository_CreateUser_Call) Run(run func(user entity.User)) *MockAuthRepository_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(entity.User))
	})
	return _c
}

func (_c *MockAuthRepository_CreateUser_Call) Return(_a0 error) *MockAuthRepository_CreateUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAuthRepository_CreateUser_Call) RunAndReturn(run func(entity.User) error) *MockAuthRepository_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// CreateUserWithTx provides a mock function with given fields: tx, user
func (_m *MockAuthRepository) CreateUserWithTx(tx *gorm.DB, user *entity.User) error {
	ret := _m.Called(tx, user)

	if len(ret) == 0 {
		panic("no return value specified for CreateUserWithTx")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *entity.User) error); ok {
		r0 = rf(tx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAuthRepository_CreateUserWithTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUserWithTx'
type MockAuthRepository_CreateUserWithTx_Call struct {
	*mock.Call
}

// CreateUserWithTx is a helper method to define mock.On call
//   - tx *gorm.DB
//   - user *entity.User
func (_e *MockAuthRepository_Expecter) CreateUserWithTx(tx interface{}, user interface{}) *MockAuthRepository_CreateUserWithTx_Call {
	return &MockAuthRepository_CreateUserWithTx_Call{Call: _e.mock.On("CreateUserWithTx", tx, user)}
}

func (_c *MockAuthRepository_CreateUserWithTx_Call) Run(run func(tx *gorm.DB, user *entity.User)) *MockAuthRepository_CreateUserWithTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(*entity.User))
	})
	return _c
}

func (_c *MockAuthRepository_CreateUserWithTx_Call) Return(_a0 error) *MockAuthRepository_CreateUserWithTx_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAuthRepository_CreateUserWithTx_Call) RunAndReturn(run func(*gorm.DB, *entity.User) error) *MockAuthRepository_CreateUserWithTx_Call {
	_c.Call.Return(run)
	return _c
}

// FindByUserID provides a mock function with given fields: userID
func (_m *MockAuthRepository) FindByUserID(userID uint) (entity.User, error) {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for FindByUserID")
	}

	var r0 entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (entity.User, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uint) entity.User); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAuthRepository_FindByUserID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByUserID'
type MockAuthRepository_FindByUserID_Call struct {
	*mock.Call
}

// FindByUserID is a helper method to define mock.On call
//   - userID uint
func (_e *MockAuthRepository_Expecter) FindByUserID(userID interface{}) *MockAuthRepository_FindByUserID_Call {
	return &MockAuthRepository_FindByUserID_Call{Call: _e.mock.On("FindByUserID", userID)}
}

func (_c *MockAuthRepository_FindByUserID_Call) Run(run func(userID uint)) *MockAuthRepository_FindByUserID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *MockAuthRepository_FindByUserID_Call) Return(_a0 entity.User, _a1 error) *MockAuthRepository_FindByUserID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAuthRepository_FindByUserID_Call) RunAndReturn(run func(uint) (entity.User, error)) *MockAuthRepository_FindByUserID_Call {
	_c.Call.Return(run)
	return _c
}

// FindByUsername provides a mock function with given fields: username
func (_m *MockAuthRepository) FindByUsername(username string) (entity.User, error) {
	ret := _m.Called(username)

	if len(ret) == 0 {
		panic("no return value specified for FindByUsername")
	}

	var r0 entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (entity.User, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) entity.User); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAuthRepository_FindByUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByUsername'
type MockAuthRepository_FindByUsername_Call struct {
	*mock.Call
}

// FindByUsername is a helper method to define mock.On call
//   - username string
func (_e *MockAuthRepository_Expecter) FindByUsername(username interface{}) *MockAuthRepository_FindByUsername_Call {
	return &MockAuthRepository_FindByUsername_Call{Call: _e.mock.On("FindByUsername", username)}
}

func (_c *MockAuthRepository_FindByUsername_Call) Run(run func(username string)) *MockAuthRepository_FindByUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockAuthRepository_FindByUsername_Call) Return(_a0 entity.User, _a1 error) *MockAuthRepository_FindByUsername_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAuthRepository_FindByUsername_Call) RunAndReturn(run func(string) (entity.User, error)) *MockAuthRepository_FindByUsername_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUser provides a mock function with given fields: user
func (_m *MockAuthRepository) UpdateUser(user entity.User) error {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAuthRepository_UpdateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUser'
type MockAuthRepository_UpdateUser_Call struct {
	*mock.Call
}

// UpdateUser is a helper method to define mock.On call
//   - user entity.User
func (_e *MockAuthRepository_Expecter) UpdateUser(user interface{}) *MockAuthRepository_UpdateUser_Call {
	return &MockAuthRepository_UpdateUser_Call{Call: _e.mock.On("UpdateUser", user)}
}

func (_c *MockAuthRepository_UpdateUser_Call) Run(run func(user entity.User)) *MockAuthRepository_UpdateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(entity.User))
	})
	return _c
}

func (_c *MockAuthRepository_UpdateUser_Call) Return(_a0 error) *MockAuthRepository_UpdateUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAuthRepository_UpdateUser_Call) RunAndReturn(run func(entity.User) error) *MockAuthRepository_UpdateUser_Call {
	_c.Call.Return(run)
	return _c
}

// WithTransaction provides a mock function with given fields: fn
func (_m *MockAuthRepository) WithTransaction(fn func(*gorm.DB) error) error {
	ret := _m.Called(fn)

	if len(ret) == 0 {
		panic("no return value specified for WithTransaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(func(*gorm.DB) error) error); ok {
		r0 = rf(fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAuthRepository_WithTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithTransaction'
type MockAuthRepository_WithTransaction_Call struct {
	*mock.Call
}

// WithTransaction is a helper method to define mock.On call
//   - fn func(*gorm.DB) error
func (_e *MockAuthRepository_Expecter) WithTransaction(fn interface{}) *MockAuthRepository_WithTransaction_Call {
	return &MockAuthRepository_WithTransaction_Call{Call: _e.mock.On("WithTransaction", fn)}
}

func (_c *MockAuthRepository_WithTransaction_Call) Run(run func(fn func(*gorm.DB) error)) *MockAuthRepository_WithTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(func(*gorm.DB) error))
	})
	return _c
}

func (_c *MockAuthRepository_WithTransaction_Call) Return(_a0 error) *MockAuthRepository_WithTransaction_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAuthRepository_WithTransaction_Call) RunAndReturn(run func(func(*gorm.DB) error) error) *MockAuthRepository_WithTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAuthRepository creates a new instance of MockAuthRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAuthRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAuthRepository {
	mock := &MockAuthRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
