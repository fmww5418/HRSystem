// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	entity "HRSystem/src/entity"

	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"
)

// MockDepartmentRepository is an autogenerated mock type for the DepartmentRepository type
type MockDepartmentRepository struct {
	mock.Mock
}

type MockDepartmentRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDepartmentRepository) EXPECT() *MockDepartmentRepository_Expecter {
	return &MockDepartmentRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: _a0
func (_m *MockDepartmentRepository) Create(_a0 entity.Department) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.Department) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDepartmentRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockDepartmentRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 entity.Department
func (_e *MockDepartmentRepository_Expecter) Create(_a0 interface{}) *MockDepartmentRepository_Create_Call {
	return &MockDepartmentRepository_Create_Call{Call: _e.mock.On("Create", _a0)}
}

func (_c *MockDepartmentRepository_Create_Call) Run(run func(_a0 entity.Department)) *MockDepartmentRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(entity.Department))
	})
	return _c
}

func (_c *MockDepartmentRepository_Create_Call) Return(_a0 error) *MockDepartmentRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDepartmentRepository_Create_Call) RunAndReturn(run func(entity.Department) error) *MockDepartmentRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// CreateWithTx provides a mock function with given fields: tx, dept
func (_m *MockDepartmentRepository) CreateWithTx(tx *gorm.DB, dept *entity.Department) error {
	ret := _m.Called(tx, dept)

	if len(ret) == 0 {
		panic("no return value specified for CreateWithTx")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *entity.Department) error); ok {
		r0 = rf(tx, dept)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDepartmentRepository_CreateWithTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateWithTx'
type MockDepartmentRepository_CreateWithTx_Call struct {
	*mock.Call
}

// CreateWithTx is a helper method to define mock.On call
//   - tx *gorm.DB
//   - dept *entity.Department
func (_e *MockDepartmentRepository_Expecter) CreateWithTx(tx interface{}, dept interface{}) *MockDepartmentRepository_CreateWithTx_Call {
	return &MockDepartmentRepository_CreateWithTx_Call{Call: _e.mock.On("CreateWithTx", tx, dept)}
}

func (_c *MockDepartmentRepository_CreateWithTx_Call) Run(run func(tx *gorm.DB, dept *entity.Department)) *MockDepartmentRepository_CreateWithTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(*entity.Department))
	})
	return _c
}

func (_c *MockDepartmentRepository_CreateWithTx_Call) Return(_a0 error) *MockDepartmentRepository_CreateWithTx_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDepartmentRepository_CreateWithTx_Call) RunAndReturn(run func(*gorm.DB, *entity.Department) error) *MockDepartmentRepository_CreateWithTx_Call {
	_c.Call.Return(run)
	return _c
}

// FindAllByOrgID provides a mock function with given fields: _a0
func (_m *MockDepartmentRepository) FindAllByOrgID(_a0 uint) ([]entity.Department, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for FindAllByOrgID")
	}

	var r0 []entity.Department
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) ([]entity.Department, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(uint) []entity.Department); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Department)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDepartmentRepository_FindAllByOrgID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAllByOrgID'
type MockDepartmentRepository_FindAllByOrgID_Call struct {
	*mock.Call
}

// FindAllByOrgID is a helper method to define mock.On call
//   - _a0 uint
func (_e *MockDepartmentRepository_Expecter) FindAllByOrgID(_a0 interface{}) *MockDepartmentRepository_FindAllByOrgID_Call {
	return &MockDepartmentRepository_FindAllByOrgID_Call{Call: _e.mock.On("FindAllByOrgID", _a0)}
}

func (_c *MockDepartmentRepository_FindAllByOrgID_Call) Run(run func(_a0 uint)) *MockDepartmentRepository_FindAllByOrgID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *MockDepartmentRepository_FindAllByOrgID_Call) Return(_a0 []entity.Department, _a1 error) *MockDepartmentRepository_FindAllByOrgID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDepartmentRepository_FindAllByOrgID_Call) RunAndReturn(run func(uint) ([]entity.Department, error)) *MockDepartmentRepository_FindAllByOrgID_Call {
	_c.Call.Return(run)
	return _c
}

// FindByID provides a mock function with given fields: _a0
func (_m *MockDepartmentRepository) FindByID(_a0 uint) (entity.Department, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for FindByID")
	}

	var r0 entity.Department
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (entity.Department, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(uint) entity.Department); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(entity.Department)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDepartmentRepository_FindByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByID'
type MockDepartmentRepository_FindByID_Call struct {
	*mock.Call
}

// FindByID is a helper method to define mock.On call
//   - _a0 uint
func (_e *MockDepartmentRepository_Expecter) FindByID(_a0 interface{}) *MockDepartmentRepository_FindByID_Call {
	return &MockDepartmentRepository_FindByID_Call{Call: _e.mock.On("FindByID", _a0)}
}

func (_c *MockDepartmentRepository_FindByID_Call) Run(run func(_a0 uint)) *MockDepartmentRepository_FindByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *MockDepartmentRepository_FindByID_Call) Return(_a0 entity.Department, _a1 error) *MockDepartmentRepository_FindByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDepartmentRepository_FindByID_Call) RunAndReturn(run func(uint) (entity.Department, error)) *MockDepartmentRepository_FindByID_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDepartmentRepository creates a new instance of MockDepartmentRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDepartmentRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDepartmentRepository {
	mock := &MockDepartmentRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}