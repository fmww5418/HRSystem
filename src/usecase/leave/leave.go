package leave

import (
	"HRSystem/src/domain/employee"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"

	dleave "HRSystem/src/domain/leave"
	"HRSystem/src/entity"
	lerrors "HRSystem/src/lib/errors"
	ltime "HRSystem/src/lib/time"
)

type (
	leaveUsecase struct {
		leaveRepo    dleave.LeaveRepository
		employeeRepo employee.EmployeeRepository
	}
)

var _ dleave.LeaveUsecase = (*leaveUsecase)(nil)

func NewLeaveUsecase(
	leaveRepo dleave.LeaveRepository,
	employeeRepo employee.EmployeeRepository,
) dleave.LeaveUsecase {
	return &leaveUsecase{leaveRepo: leaveRepo, employeeRepo: employeeRepo}
}

func (s *leaveUsecase) GetAllLeaveRequests() ([]entity.Request, error) {
	return s.leaveRepo.FindAll()
}

func (s *leaveUsecase) GetLeaveRequestByID(id uint) (entity.Request, error) {
	return s.leaveRepo.FindByID(id)
}

func (s *leaveUsecase) CreateLeaveRequest(userID uint, req dleave.LeaveRequestInput) error {
	// TODO: must verify if it is overlapping with others?

	startTime := ltime.ParseTime(time.DateTime, req.StartDate, nil)
	endTime := ltime.ParseTime(time.DateTime, req.EndDate, nil)
	if endTime.Before(endTime) || startTime.After(endTime) {
		return fmt.Errorf("start date time can't be before end date")
	}

	employeeUser, err := s.employeeRepo.FindByUserID(userID)
	if err != nil {
		return err
	}

	leaveRequest := entity.Request{
		EmployeeID:  employeeUser.ID,
		StartDate:   startTime,
		EndDate:     endTime,
		Status:      entity.RequestStatusPending,
		Description: req.Description,
	}
	return s.leaveRepo.Create(leaveRequest)
}

func (s *leaveUsecase) UpdateLeaveRequestStatus(id uint, operatorUserID uint, req dleave.UpdateLeaveStatusInput) error {
	operatorEmployee, err := s.employeeRepo.FindByUserID(operatorUserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("cannot find the employee by userID: %v", operatorUserID)
		}
		return fmt.Errorf("%w: %v", lerrors.ErrForbidden, err)
	}

	leaveReq, err := s.leaveRepo.FindByID(id)
	if err != nil {
		return err
	}

	if leaveReq.Employee.SupervisorEmployeeID != nil && *leaveReq.Employee.SupervisorEmployeeID != operatorEmployee.ID {
		return fmt.Errorf("%w: is not allowed to update request status", lerrors.ErrForbidden)
	}

	if *req.Status < entity.RequestStatusPending || *req.Status >= entity.RequestStatusMax {
		return fmt.Errorf("%w: target status is illega", lerrors.ErrForbidden)
	}

	return s.leaveRepo.UpdateStatus(id, *req.Status)
}
