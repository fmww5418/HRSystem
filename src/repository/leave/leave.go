package leave

import (
	dleave "HRSystem/src/domain/leave"
	"HRSystem/src/entity"
	"gorm.io/gorm"
)

type leaveRepository struct {
	db *gorm.DB
}

var _ dleave.LeaveRepository = (*leaveRepository)(nil)

func NewLeaveRepository(db *gorm.DB) dleave.LeaveRepository {
	return &leaveRepository{db: db}
}

func (r *leaveRepository) FindAll() ([]entity.Request, error) {
	var leaveRequests []entity.Request
	err := r.db.Find(&leaveRequests).Error
	return leaveRequests, err
}

func (r *leaveRepository) FindByID(id uint) (entity.Request, error) {
	var leaveRequest entity.Request
	err := r.db.Preload("Employee").First(&leaveRequest, id).Error
	return leaveRequest, err
}

func (r *leaveRepository) Create(leaveRequest entity.Request) error {
	return r.db.Create(&leaveRequest).Error
}

func (r *leaveRepository) UpdateStatus(id uint, status entity.RequestStatus) error {
	return r.db.Model(&entity.Request{}).
		Where("id = ?", id).
		Update("status", status).Error
}
