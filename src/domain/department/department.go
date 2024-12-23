package department

import (
	"HRSystem/src/entity"
	"gorm.io/gorm"
)

type (
	DepartmentRepository interface {
		Create(entity.Department) error
		FindByID(uint) (entity.Department, error)
		FindAllByOrgID(uint) ([]entity.Department, error)
		CreateWithTx(tx *gorm.DB, dept *entity.Department) error
	}

	DepartmentUsecase interface {
		CreateDepartment(userID uint, req DepartmentCreateReq) error
		GetAllDepartments(userID uint) ([]entity.Department, error)
		Invite(dept entity.Department, userID uint) error
	}

	DepartmentCreateReq struct {
		Name           string `json:"name" binding:"required"`
		OrganizationID uint   `json:"organization_id" binding:"required"`
	}

	DepartmentGetReq struct {
		OrganizationID uint `json:"organization_id" binding:"required"`
	}
)
