package leave

import (
	"HRSystem/src/entity"
)

type LeaveRepository interface {
	FindAll() ([]entity.Request, error)
	FindByID(id uint) (entity.Request, error)
	Create(leaveRequest entity.Request) error
	UpdateStatus(id uint, status entity.RequestStatus) error
}

type (
	LeaveRequestInput struct {
		StartDate   string `json:"start_date" binding:"required"`
		EndDate     string `json:"end_date" binding:"required"`
		Description string `json:"description"`
	}

	LeaveUsecase interface {
		GetAllLeaveRequests() ([]entity.Request, error)
		GetLeaveRequestByID(id uint) (entity.Request, error)
		CreateLeaveRequest(userID uint, req LeaveRequestInput) error
		UpdateLeaveRequestStatus(id uint, operatorUserID uint, req UpdateLeaveStatusInput) error
	}
)

type UpdateLeaveStatusInput struct {
	Status *entity.RequestStatus `json:"status" binding:"required"`
}
