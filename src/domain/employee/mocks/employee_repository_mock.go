// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	entity "HRSystem/src/entity"

	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"
)

// MockEmployeeRepository is an autogenerated mock type for the EmployeeRepository type
type MockEmployeeRepository struct {
	mock.Mock
}

type MockEmployeeRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEmployeeRepository) EXPECT() *MockEmployeeRepository_Expecter {
	return &MockEmployeeRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: _a0
func (_m *MockEmployeeRepository) Create(_a0 entity.Employee) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.Employee) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEmployeeRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockEmployeeRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 entity.Employee
func (_e *MockEmployeeRepository_Expecter) Create(_a0 interface{}) *MockEmployeeRepository_Create_Call {
	return &MockEmployeeRepository_Create_Call{Call: _e.mock.On("Create", _a0)}
}

func (_c *MockEmployeeRepository_Create_Call) Run(run func(_a0 entity.Employee)) *MockEmployeeRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(entity.Employee))
	})
	return _c
}

func (_c *MockEmployeeRepository_Create_Call) Return(_a0 error) *MockEmployeeRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEmployeeRepository_Create_Call) RunAndReturn(run func(entity.Employee) error) *MockEmployeeRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// CreateWithTx provides a mock function with given fields: tx, _a1
func (_m *MockEmployeeRepository) CreateWithTx(tx *gorm.DB, _a1 *entity.Employee) error {
	ret := _m.Called(tx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateWithTx")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *entity.Employee) error); ok {
		r0 = rf(tx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEmployeeRepository_CreateWithTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateWithTx'
type MockEmployeeRepository_CreateWithTx_Call struct {
	*mock.Call
}

// CreateWithTx is a helper method to define mock.On call
//   - tx *gorm.DB
//   - _a1 *entity.Employee
func (_e *MockEmployeeRepository_Expecter) CreateWithTx(tx interface{}, _a1 interface{}) *MockEmployeeRepository_CreateWithTx_Call {
	return &MockEmployeeRepository_CreateWithTx_Call{Call: _e.mock.On("CreateWithTx", tx, _a1)}
}

func (_c *MockEmployeeRepository_CreateWithTx_Call) Run(run func(tx *gorm.DB, _a1 *entity.Employee)) *MockEmployeeRepository_CreateWithTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(*entity.Employee))
	})
	return _c
}

func (_c *MockEmployeeRepository_CreateWithTx_Call) Return(_a0 error) *MockEmployeeRepository_CreateWithTx_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEmployeeRepository_CreateWithTx_Call) RunAndReturn(run func(*gorm.DB, *entity.Employee) error) *MockEmployeeRepository_CreateWithTx_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: id
func (_m *MockEmployeeRepository) Delete(id uint) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEmployeeRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockEmployeeRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - id uint
func (_e *MockEmployeeRepository_Expecter) Delete(id interface{}) *MockEmployeeRepository_Delete_Call {
	return &MockEmployeeRepository_Delete_Call{Call: _e.mock.On("Delete", id)}
}

func (_c *MockEmployeeRepository_Delete_Call) Run(run func(id uint)) *MockEmployeeRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *MockEmployeeRepository_Delete_Call) Return(_a0 error) *MockEmployeeRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEmployeeRepository_Delete_Call) RunAndReturn(run func(uint) error) *MockEmployeeRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// FindAllByUserID provides a mock function with given fields: operatorUserID
func (_m *MockEmployeeRepository) FindAllByUserID(operatorUserID uint) ([]entity.Employee, error) {
	ret := _m.Called(operatorUserID)

	if len(ret) == 0 {
		panic("no return value specified for FindAllByUserID")
	}

	var r0 []entity.Employee
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) ([]entity.Employee, error)); ok {
		return rf(operatorUserID)
	}
	if rf, ok := ret.Get(0).(func(uint) []entity.Employee); ok {
		r0 = rf(operatorUserID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Employee)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(operatorUserID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEmployeeRepository_FindAllByUserID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAllByUserID'
type MockEmployeeRepository_FindAllByUserID_Call struct {
	*mock.Call
}

// FindAllByUserID is a helper method to define mock.On call
//   - operatorUserID uint
func (_e *MockEmployeeRepository_Expecter) FindAllByUserID(operatorUserID interface{}) *MockEmployeeRepository_FindAllByUserID_Call {
	return &MockEmployeeRepository_FindAllByUserID_Call{Call: _e.mock.On("FindAllByUserID", operatorUserID)}
}

func (_c *MockEmployeeRepository_FindAllByUserID_Call) Run(run func(operatorUserID uint)) *MockEmployeeRepository_FindAllByUserID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *MockEmployeeRepository_FindAllByUserID_Call) Return(_a0 []entity.Employee, _a1 error) *MockEmployeeRepository_FindAllByUserID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEmployeeRepository_FindAllByUserID_Call) RunAndReturn(run func(uint) ([]entity.Employee, error)) *MockEmployeeRepository_FindAllByUserID_Call {
	_c.Call.Return(run)
	return _c
}

// FindByID provides a mock function with given fields: id
func (_m *MockEmployeeRepository) FindByID(id uint) (entity.Employee, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindByID")
	}

	var r0 entity.Employee
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (entity.Employee, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) entity.Employee); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entity.Employee)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEmployeeRepository_FindByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByID'
type MockEmployeeRepository_FindByID_Call struct {
	*mock.Call
}

// FindByID is a helper method to define mock.On call
//   - id uint
func (_e *MockEmployeeRepository_Expecter) FindByID(id interface{}) *MockEmployeeRepository_FindByID_Call {
	return &MockEmployeeRepository_FindByID_Call{Call: _e.mock.On("FindByID", id)}
}

func (_c *MockEmployeeRepository_FindByID_Call) Run(run func(id uint)) *MockEmployeeRepository_FindByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *MockEmployeeRepository_FindByID_Call) Return(_a0 entity.Employee, _a1 error) *MockEmployeeRepository_FindByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEmployeeRepository_FindByID_Call) RunAndReturn(run func(uint) (entity.Employee, error)) *MockEmployeeRepository_FindByID_Call {
	_c.Call.Return(run)
	return _c
}

// FindByIDWithOrgCheck provides a mock function with given fields: userID, id
func (_m *MockEmployeeRepository) FindByIDWithOrgCheck(userID uint, id uint) (entity.Employee, error) {
	ret := _m.Called(userID, id)

	if len(ret) == 0 {
		panic("no return value specified for FindByIDWithOrgCheck")
	}

	var r0 entity.Employee
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, uint) (entity.Employee, error)); ok {
		return rf(userID, id)
	}
	if rf, ok := ret.Get(0).(func(uint, uint) entity.Employee); ok {
		r0 = rf(userID, id)
	} else {
		r0 = ret.Get(0).(entity.Employee)
	}

	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(userID, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEmployeeRepository_FindByIDWithOrgCheck_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByIDWithOrgCheck'
type MockEmployeeRepository_FindByIDWithOrgCheck_Call struct {
	*mock.Call
}

// FindByIDWithOrgCheck is a helper method to define mock.On call
//   - userID uint
//   - id uint
func (_e *MockEmployeeRepository_Expecter) FindByIDWithOrgCheck(userID interface{}, id interface{}) *MockEmployeeRepository_FindByIDWithOrgCheck_Call {
	return &MockEmployeeRepository_FindByIDWithOrgCheck_Call{Call: _e.mock.On("FindByIDWithOrgCheck", userID, id)}
}

func (_c *MockEmployeeRepository_FindByIDWithOrgCheck_Call) Run(run func(userID uint, id uint)) *MockEmployeeRepository_FindByIDWithOrgCheck_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint), args[1].(uint))
	})
	return _c
}

func (_c *MockEmployeeRepository_FindByIDWithOrgCheck_Call) Return(_a0 entity.Employee, _a1 error) *MockEmployeeRepository_FindByIDWithOrgCheck_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEmployeeRepository_FindByIDWithOrgCheck_Call) RunAndReturn(run func(uint, uint) (entity.Employee, error)) *MockEmployeeRepository_FindByIDWithOrgCheck_Call {
	_c.Call.Return(run)
	return _c
}

// FindByUserID provides a mock function with given fields: id
func (_m *MockEmployeeRepository) FindByUserID(id uint) (entity.Employee, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindByUserID")
	}

	var r0 entity.Employee
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (entity.Employee, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) entity.Employee); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entity.Employee)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEmployeeRepository_FindByUserID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByUserID'
type MockEmployeeRepository_FindByUserID_Call struct {
	*mock.Call
}

// FindByUserID is a helper method to define mock.On call
//   - id uint
func (_e *MockEmployeeRepository_Expecter) FindByUserID(id interface{}) *MockEmployeeRepository_FindByUserID_Call {
	return &MockEmployeeRepository_FindByUserID_Call{Call: _e.mock.On("FindByUserID", id)}
}

func (_c *MockEmployeeRepository_FindByUserID_Call) Run(run func(id uint)) *MockEmployeeRepository_FindByUserID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *MockEmployeeRepository_FindByUserID_Call) Return(_a0 entity.Employee, _a1 error) *MockEmployeeRepository_FindByUserID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEmployeeRepository_FindByUserID_Call) RunAndReturn(run func(uint) (entity.Employee, error)) *MockEmployeeRepository_FindByUserID_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: id, _a1
func (_m *MockEmployeeRepository) Update(id uint, _a1 entity.Employee) error {
	ret := _m.Called(id, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, entity.Employee) error); ok {
		r0 = rf(id, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEmployeeRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockEmployeeRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - id uint
//   - _a1 entity.Employee
func (_e *MockEmployeeRepository_Expecter) Update(id interface{}, _a1 interface{}) *MockEmployeeRepository_Update_Call {
	return &MockEmployeeRepository_Update_Call{Call: _e.mock.On("Update", id, _a1)}
}

func (_c *MockEmployeeRepository_Update_Call) Run(run func(id uint, _a1 entity.Employee)) *MockEmployeeRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint), args[1].(entity.Employee))
	})
	return _c
}

func (_c *MockEmployeeRepository_Update_Call) Return(_a0 error) *MockEmployeeRepository_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEmployeeRepository_Update_Call) RunAndReturn(run func(uint, entity.Employee) error) *MockEmployeeRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateDeptIDWithTxByUserID provides a mock function with given fields: tx, id, deptID
func (_m *MockEmployeeRepository) UpdateDeptIDWithTxByUserID(tx *gorm.DB, id uint, deptID uint) error {
	ret := _m.Called(tx, id, deptID)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDeptIDWithTxByUserID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, uint, uint) error); ok {
		r0 = rf(tx, id, deptID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEmployeeRepository_UpdateDeptIDWithTxByUserID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateDeptIDWithTxByUserID'
type MockEmployeeRepository_UpdateDeptIDWithTxByUserID_Call struct {
	*mock.Call
}

// UpdateDeptIDWithTxByUserID is a helper method to define mock.On call
//   - tx *gorm.DB
//   - id uint
//   - deptID uint
func (_e *MockEmployeeRepository_Expecter) UpdateDeptIDWithTxByUserID(tx interface{}, id interface{}, deptID interface{}) *MockEmployeeRepository_UpdateDeptIDWithTxByUserID_Call {
	return &MockEmployeeRepository_UpdateDeptIDWithTxByUserID_Call{Call: _e.mock.On("UpdateDeptIDWithTxByUserID", tx, id, deptID)}
}

func (_c *MockEmployeeRepository_UpdateDeptIDWithTxByUserID_Call) Run(run func(tx *gorm.DB, id uint, deptID uint)) *MockEmployeeRepository_UpdateDeptIDWithTxByUserID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(uint), args[2].(uint))
	})
	return _c
}

func (_c *MockEmployeeRepository_UpdateDeptIDWithTxByUserID_Call) Return(_a0 error) *MockEmployeeRepository_UpdateDeptIDWithTxByUserID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEmployeeRepository_UpdateDeptIDWithTxByUserID_Call) RunAndReturn(run func(*gorm.DB, uint, uint) error) *MockEmployeeRepository_UpdateDeptIDWithTxByUserID_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockEmployeeRepository creates a new instance of MockEmployeeRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockEmployeeRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockEmployeeRepository {
	mock := &MockEmployeeRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
