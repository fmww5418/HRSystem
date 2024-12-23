//go:build unit

package leave

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mEmployee "HRSystem/src/domain/employee/mocks"
	dleave "HRSystem/src/domain/leave"
	"HRSystem/src/domain/leave/mocks"
	"HRSystem/src/entity"
)

func TestCreateLeaveRequest(t *testing.T) {
	type testCase struct {
		name      string
		userID    uint
		req       dleave.LeaveRequestInput
		mockFn    func(t *testing.T, userID uint, req dleave.LeaveRequestInput) dleave.LeaveUsecase
		expectErr error
	}

	var testCases = []testCase{
		{
			name:   "Success",
			userID: 1,
			req: dleave.LeaveRequestInput{
				StartDate:   "2023-10-01 11:00:00",
				EndDate:     "2023-10-10 14:00:00",
				Description: "Vacation",
			},
			mockFn: func(t *testing.T, userID uint, req dleave.LeaveRequestInput) dleave.LeaveUsecase {
				mockRepo := mocks.NewMockLeaveRepository(t)
				mockEmployeeRepo := mEmployee.NewMockEmployeeRepository(t)
				usecase := NewLeaveUsecase(mockRepo, mockEmployeeRepo)

				mockEmployeeRepo.EXPECT().FindByUserID(userID).Return(entity.Employee{ID: userID}, nil).Once()
				mockRepo.EXPECT().Create(mock.AnythingOfType("entity.Request")).Return(nil).Once()

				return usecase
			},
			expectErr: nil,
		},
		{
			name:   "Employee Not Found",
			userID: 2,
			req: dleave.LeaveRequestInput{
				StartDate:   "2023-10-01 12:00:00",
				EndDate:     "2023-10-10 13:00:00",
				Description: "Vacation",
			},
			mockFn: func(t *testing.T, userID uint, req dleave.LeaveRequestInput) dleave.LeaveUsecase {
				mockRepo := mocks.NewMockLeaveRepository(t)
				mockEmployeeRepo := mEmployee.NewMockEmployeeRepository(t)
				usecase := NewLeaveUsecase(mockRepo, mockEmployeeRepo)

				mockEmployeeRepo.EXPECT().FindByUserID(userID).Return(entity.Employee{}, errors.New("not found")).Once()

				return usecase
			},
			expectErr: errors.New("not found"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			uLeave := tc.mockFn(tt, tc.userID, tc.req)

			err := uLeave.CreateLeaveRequest(tc.userID, tc.req)

			assert.Equal(tt, tc.expectErr, err)
		})
	}
}
