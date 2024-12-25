//go:build unit

package employee

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mDept "HRSystem/src/domain/department/mocks"
	demployee "HRSystem/src/domain/employee"
	"HRSystem/src/domain/employee/mocks"
	"HRSystem/src/entity"
)

func TestCreateEmployee(t *testing.T) {
	type testCase struct {
		name           string
		operatorUserID uint
		req            demployee.EmployeeRequest
		mockFn         func(t *testing.T, operatorUserID uint, req demployee.EmployeeRequest) demployee.EmployeeUsecase
		expectErr      error
	}

	var testCases = []testCase{
		{
			name:           "Success",
			operatorUserID: 1,
			req: demployee.EmployeeRequest{
				Name:           "John Doe",
				Position:       "Developer",
				ContactInfo:    "john.doe@example.com",
				UserID:         2,
				Salary:         float64Ptr(50000),
				RemainedDayOff: intPtr(10),
			},
			mockFn: func(t *testing.T, operatorUserID uint, req demployee.EmployeeRequest) demployee.EmployeeUsecase {
				mockRepo := mocks.NewMockEmployeeRepository(t)
				usecase := NewEmployeeUsecase(mockRepo, nil)

				mockRepo.EXPECT().FindByUserID(req.UserID).Return(entity.Employee{}, errors.New("not found")).Once()
				mockRepo.EXPECT().Create(mock.AnythingOfType("entity.Employee")).Return(nil).Once()

				return usecase
			},
			expectErr: nil,
		},
		{
			name:           "Employee Already Exists",
			operatorUserID: 1,
			req: demployee.EmployeeRequest{
				Name:           "Jane Doe",
				Position:       "Manager",
				ContactInfo:    "jane.doe@example.com",
				UserID:         3,
				Salary:         float64Ptr(60000),
				RemainedDayOff: intPtr(15),
			},
			mockFn: func(t *testing.T, operatorUserID uint, req demployee.EmployeeRequest) demployee.EmployeeUsecase {
				mockRepo := mocks.NewMockEmployeeRepository(t)
				usecase := NewEmployeeUsecase(mockRepo, nil)

				mockRepo.EXPECT().FindByUserID(req.UserID).Return(entity.Employee{ID: req.UserID}, nil).Once()

				return usecase
			},
			expectErr: errors.New("employee with UserID 3 already exists"),
		},
		{
			name:           "Department ID Provided",
			operatorUserID: 1,
			req: demployee.EmployeeRequest{
				Name:           "Alice Smith",
				Position:       "Analyst",
				ContactInfo:    "alice.smith@example.com",
				UserID:         4,
				Salary:         float64Ptr(55000),
				RemainedDayOff: intPtr(12),
				DepartmentID:   uintPtr(1),
			},
			mockFn: func(t *testing.T, operatorUserID uint, req demployee.EmployeeRequest) demployee.EmployeeUsecase {
				mockRepo := mocks.NewMockEmployeeRepository(t)
				mockDeptRepo := mDept.NewMockDepartmentRepository(t)
				usecase := NewEmployeeUsecase(mockRepo, mockDeptRepo)

				mockRepo.EXPECT().FindByUserID(req.UserID).Return(entity.Employee{}, errors.New("not found")).Once()
				mockRepo.EXPECT().FindByUserID(operatorUserID).Return(entity.Employee{Department: &entity.Department{OrganizationID: 1}}, nil).Once()
				mockDeptRepo.EXPECT().FindByID(*req.DepartmentID).Return(entity.Department{OrganizationID: 1}, nil).Once()
				mockRepo.EXPECT().Create(mock.AnythingOfType("entity.Employee")).Return(nil).Once()

				return usecase
			},
			expectErr: nil,
		},
		{
			name:           "Department and Supervisor ID Provided",
			operatorUserID: 1,
			req: demployee.EmployeeRequest{
				Name:                 "Bob Johnson",
				Position:             "Team Lead",
				ContactInfo:          "bob.johnson@example.com",
				UserID:               5,
				Salary:               float64Ptr(70000),
				RemainedDayOff:       intPtr(20),
				DepartmentID:         uintPtr(1),
				SupervisorEmployeeID: uintPtr(6),
			},
			mockFn: func(t *testing.T, operatorUserID uint, req demployee.EmployeeRequest) demployee.EmployeeUsecase {
				mockRepo := mocks.NewMockEmployeeRepository(t)
				mockDeptRepo := mDept.NewMockDepartmentRepository(t)
				usecase := NewEmployeeUsecase(mockRepo, mockDeptRepo)

				mockRepo.EXPECT().FindByUserID(req.UserID).Return(entity.Employee{}, errors.New("not found")).Once()
				mockRepo.EXPECT().FindByUserID(operatorUserID).Return(entity.Employee{Department: &entity.Department{OrganizationID: 1}}, nil).Once()
				mockDeptRepo.EXPECT().FindByID(*req.DepartmentID).Return(entity.Department{OrganizationID: 1}, nil).Once()
				mockRepo.EXPECT().FindByID(*req.SupervisorEmployeeID).Return(entity.Employee{Department: &entity.Department{OrganizationID: 1}}, nil).Once()
				mockRepo.EXPECT().Create(mock.AnythingOfType("entity.Employee")).Return(nil).Once()

				return usecase
			},
			expectErr: nil,
		},
		{
			name:           "Department Organization Mismatch",
			operatorUserID: 1,
			req: demployee.EmployeeRequest{
				Name:           "David Green",
				Position:       "Architect",
				ContactInfo:    "david.green@example.com",
				UserID:         8,
				Salary:         float64Ptr(75000),
				RemainedDayOff: intPtr(25),
				DepartmentID:   uintPtr(3),
			},
			mockFn: func(t *testing.T, operatorUserID uint, req demployee.EmployeeRequest) demployee.EmployeeUsecase {
				mockRepo := mocks.NewMockEmployeeRepository(t)
				mockDeptRepo := mDept.NewMockDepartmentRepository(t)
				usecase := NewEmployeeUsecase(mockRepo, mockDeptRepo)

				mockRepo.EXPECT().FindByUserID(req.UserID).Return(entity.Employee{}, errors.New("not found")).Once()
				mockRepo.EXPECT().FindByUserID(operatorUserID).Return(entity.Employee{Department: &entity.Department{OrganizationID: 1}}, nil).Once()
				mockDeptRepo.EXPECT().FindByID(*req.DepartmentID).Return(entity.Department{OrganizationID: 2}, nil).Once()

				return usecase
			},
			expectErr: errors.New("you are not allowed to create an employee with an outside organization. departmentID: 3"),
		},
		{
			name:           "Supervisor Department Nil",
			operatorUserID: 1,
			req: demployee.EmployeeRequest{
				Name:                 "Eve White",
				Position:             "Director",
				ContactInfo:          "eve.white@example.com",
				UserID:               9,
				Salary:               float64Ptr(90000),
				RemainedDayOff:       intPtr(30),
				DepartmentID:         uintPtr(4),
				SupervisorEmployeeID: uintPtr(10),
			},
			mockFn: func(t *testing.T, operatorUserID uint, req demployee.EmployeeRequest) demployee.EmployeeUsecase {
				mockRepo := mocks.NewMockEmployeeRepository(t)
				mockDeptRepo := mDept.NewMockDepartmentRepository(t)
				usecase := NewEmployeeUsecase(mockRepo, mockDeptRepo)

				mockRepo.EXPECT().FindByUserID(req.UserID).Return(entity.Employee{}, errors.New("not found")).Once()
				mockRepo.EXPECT().FindByUserID(operatorUserID).Return(entity.Employee{Department: &entity.Department{OrganizationID: 1}}, nil).Once()
				mockDeptRepo.EXPECT().FindByID(*req.DepartmentID).Return(entity.Department{OrganizationID: 1}, nil).Once()
				mockRepo.EXPECT().FindByID(*req.SupervisorEmployeeID).Return(entity.Employee{Department: nil}, nil).Once()

				return usecase
			},
			expectErr: errors.New("you are not allowed to create an supervisor with an outside organization"),
		},
		{
			name:           "Supervisor Organization Mismatch",
			operatorUserID: 1,
			req: demployee.EmployeeRequest{
				Name:                 "Frank Black",
				Position:             "VP",
				ContactInfo:          "frank.black@example.com",
				UserID:               11,
				Salary:               float64Ptr(95000),
				RemainedDayOff:       intPtr(35),
				DepartmentID:         uintPtr(5),
				SupervisorEmployeeID: uintPtr(12),
			},
			mockFn: func(t *testing.T, operatorUserID uint, req demployee.EmployeeRequest) demployee.EmployeeUsecase {
				mockRepo := mocks.NewMockEmployeeRepository(t)
				mockDeptRepo := mDept.NewMockDepartmentRepository(t)
				usecase := NewEmployeeUsecase(mockRepo, mockDeptRepo)

				mockRepo.EXPECT().FindByUserID(req.UserID).Return(entity.Employee{}, errors.New("not found")).Once()
				mockRepo.EXPECT().FindByUserID(operatorUserID).Return(entity.Employee{Department: &entity.Department{OrganizationID: 1}}, nil).Once()
				mockDeptRepo.EXPECT().FindByID(*req.DepartmentID).Return(entity.Department{OrganizationID: 1}, nil).Once()
				mockRepo.EXPECT().FindByID(*req.SupervisorEmployeeID).Return(entity.Employee{Department: &entity.Department{OrganizationID: 2}}, nil).Once()

				return usecase
			},
			expectErr: errors.New("you are not allowed to create an supervisor with an outside organization"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			uEmployee := tc.mockFn(tt, tc.operatorUserID, tc.req)

			err := uEmployee.CreateEmployee(tc.operatorUserID, tc.req)

			assert.Equal(tt, tc.expectErr, err)
		})
	}
}

func float64Ptr(f float64) *float64 {
	return &f
}

func intPtr(i int) *int {
	return &i
}

func uintPtr(u uint) *uint {
	return &u
}
