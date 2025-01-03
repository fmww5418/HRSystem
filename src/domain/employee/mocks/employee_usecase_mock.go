// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	employee "HRSystem/src/domain/employee"
	entity "HRSystem/src/entity"

	mock "github.com/stretchr/testify/mock"
)

// MockEmployeeUsecase is an autogenerated mock type for the EmployeeUsecase type
type MockEmployeeUsecase struct {
	mock.Mock
}

type MockEmployeeUsecase_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEmployeeUsecase) EXPECT() *MockEmployeeUsecase_Expecter {
	return &MockEmployeeUsecase_Expecter{mock: &_m.Mock}
}

// CreateEmployee provides a mock function with given fields: operatorUserID, req
func (_m *MockEmployeeUsecase) CreateEmployee(operatorUserID uint, req employee.EmployeeRequest) error {
	ret := _m.Called(operatorUserID, req)

	if len(ret) == 0 {
		panic("no return value specified for CreateEmployee")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, employee.EmployeeRequest) error); ok {
		r0 = rf(operatorUserID, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEmployeeUsecase_CreateEmployee_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateEmployee'
type MockEmployeeUsecase_CreateEmployee_Call struct {
	*mock.Call
}

// CreateEmployee is a helper method to define mock.On call
//   - operatorUserID uint
//   - req employee.EmployeeRequest
func (_e *MockEmployeeUsecase_Expecter) CreateEmployee(operatorUserID interface{}, req interface{}) *MockEmployeeUsecase_CreateEmployee_Call {
	return &MockEmployeeUsecase_CreateEmployee_Call{Call: _e.mock.On("CreateEmployee", operatorUserID, req)}
}

func (_c *MockEmployeeUsecase_CreateEmployee_Call) Run(run func(operatorUserID uint, req employee.EmployeeRequest)) *MockEmployeeUsecase_CreateEmployee_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint), args[1].(employee.EmployeeRequest))
	})
	return _c
}

func (_c *MockEmployeeUsecase_CreateEmployee_Call) Return(_a0 error) *MockEmployeeUsecase_CreateEmployee_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEmployeeUsecase_CreateEmployee_Call) RunAndReturn(run func(uint, employee.EmployeeRequest) error) *MockEmployeeUsecase_CreateEmployee_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteEmployee provides a mock function with given fields: operatorUserID, id
func (_m *MockEmployeeUsecase) DeleteEmployee(operatorUserID uint, id uint) error {
	ret := _m.Called(operatorUserID, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteEmployee")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(operatorUserID, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEmployeeUsecase_DeleteEmployee_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteEmployee'
type MockEmployeeUsecase_DeleteEmployee_Call struct {
	*mock.Call
}

// DeleteEmployee is a helper method to define mock.On call
//   - operatorUserID uint
//   - id uint
func (_e *MockEmployeeUsecase_Expecter) DeleteEmployee(operatorUserID interface{}, id interface{}) *MockEmployeeUsecase_DeleteEmployee_Call {
	return &MockEmployeeUsecase_DeleteEmployee_Call{Call: _e.mock.On("DeleteEmployee", operatorUserID, id)}
}

func (_c *MockEmployeeUsecase_DeleteEmployee_Call) Run(run func(operatorUserID uint, id uint)) *MockEmployeeUsecase_DeleteEmployee_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint), args[1].(uint))
	})
	return _c
}

func (_c *MockEmployeeUsecase_DeleteEmployee_Call) Return(_a0 error) *MockEmployeeUsecase_DeleteEmployee_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEmployeeUsecase_DeleteEmployee_Call) RunAndReturn(run func(uint, uint) error) *MockEmployeeUsecase_DeleteEmployee_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllEmployees provides a mock function with given fields: operatorUserID
func (_m *MockEmployeeUsecase) GetAllEmployees(operatorUserID uint) ([]entity.Employee, error) {
	ret := _m.Called(operatorUserID)

	if len(ret) == 0 {
		panic("no return value specified for GetAllEmployees")
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

// MockEmployeeUsecase_GetAllEmployees_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllEmployees'
type MockEmployeeUsecase_GetAllEmployees_Call struct {
	*mock.Call
}

// GetAllEmployees is a helper method to define mock.On call
//   - operatorUserID uint
func (_e *MockEmployeeUsecase_Expecter) GetAllEmployees(operatorUserID interface{}) *MockEmployeeUsecase_GetAllEmployees_Call {
	return &MockEmployeeUsecase_GetAllEmployees_Call{Call: _e.mock.On("GetAllEmployees", operatorUserID)}
}

func (_c *MockEmployeeUsecase_GetAllEmployees_Call) Run(run func(operatorUserID uint)) *MockEmployeeUsecase_GetAllEmployees_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *MockEmployeeUsecase_GetAllEmployees_Call) Return(_a0 []entity.Employee, _a1 error) *MockEmployeeUsecase_GetAllEmployees_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEmployeeUsecase_GetAllEmployees_Call) RunAndReturn(run func(uint) ([]entity.Employee, error)) *MockEmployeeUsecase_GetAllEmployees_Call {
	_c.Call.Return(run)
	return _c
}

// GetEmployeeByID provides a mock function with given fields: operatorUserID, id
func (_m *MockEmployeeUsecase) GetEmployeeByID(operatorUserID uint, id uint) (entity.Employee, error) {
	ret := _m.Called(operatorUserID, id)

	if len(ret) == 0 {
		panic("no return value specified for GetEmployeeByID")
	}

	var r0 entity.Employee
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, uint) (entity.Employee, error)); ok {
		return rf(operatorUserID, id)
	}
	if rf, ok := ret.Get(0).(func(uint, uint) entity.Employee); ok {
		r0 = rf(operatorUserID, id)
	} else {
		r0 = ret.Get(0).(entity.Employee)
	}

	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(operatorUserID, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEmployeeUsecase_GetEmployeeByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEmployeeByID'
type MockEmployeeUsecase_GetEmployeeByID_Call struct {
	*mock.Call
}

// GetEmployeeByID is a helper method to define mock.On call
//   - operatorUserID uint
//   - id uint
func (_e *MockEmployeeUsecase_Expecter) GetEmployeeByID(operatorUserID interface{}, id interface{}) *MockEmployeeUsecase_GetEmployeeByID_Call {
	return &MockEmployeeUsecase_GetEmployeeByID_Call{Call: _e.mock.On("GetEmployeeByID", operatorUserID, id)}
}

func (_c *MockEmployeeUsecase_GetEmployeeByID_Call) Run(run func(operatorUserID uint, id uint)) *MockEmployeeUsecase_GetEmployeeByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint), args[1].(uint))
	})
	return _c
}

func (_c *MockEmployeeUsecase_GetEmployeeByID_Call) Return(_a0 entity.Employee, _a1 error) *MockEmployeeUsecase_GetEmployeeByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEmployeeUsecase_GetEmployeeByID_Call) RunAndReturn(run func(uint, uint) (entity.Employee, error)) *MockEmployeeUsecase_GetEmployeeByID_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateEmployee provides a mock function with given fields: operatorUserID, id, req
func (_m *MockEmployeeUsecase) UpdateEmployee(operatorUserID uint, id uint, req employee.EmployeeRequest) error {
	ret := _m.Called(operatorUserID, id, req)

	if len(ret) == 0 {
		panic("no return value specified for UpdateEmployee")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint, employee.EmployeeRequest) error); ok {
		r0 = rf(operatorUserID, id, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEmployeeUsecase_UpdateEmployee_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateEmployee'
type MockEmployeeUsecase_UpdateEmployee_Call struct {
	*mock.Call
}

// UpdateEmployee is a helper method to define mock.On call
//   - operatorUserID uint
//   - id uint
//   - req employee.EmployeeRequest
func (_e *MockEmployeeUsecase_Expecter) UpdateEmployee(operatorUserID interface{}, id interface{}, req interface{}) *MockEmployeeUsecase_UpdateEmployee_Call {
	return &MockEmployeeUsecase_UpdateEmployee_Call{Call: _e.mock.On("UpdateEmployee", operatorUserID, id, req)}
}

func (_c *MockEmployeeUsecase_UpdateEmployee_Call) Run(run func(operatorUserID uint, id uint, req employee.EmployeeRequest)) *MockEmployeeUsecase_UpdateEmployee_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint), args[1].(uint), args[2].(employee.EmployeeRequest))
	})
	return _c
}

func (_c *MockEmployeeUsecase_UpdateEmployee_Call) Return(_a0 error) *MockEmployeeUsecase_UpdateEmployee_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEmployeeUsecase_UpdateEmployee_Call) RunAndReturn(run func(uint, uint, employee.EmployeeRequest) error) *MockEmployeeUsecase_UpdateEmployee_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockEmployeeUsecase creates a new instance of MockEmployeeUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockEmployeeUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockEmployeeUsecase {
	mock := &MockEmployeeUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
