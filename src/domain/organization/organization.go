package organization

import (
	"HRSystem/src/entity"
	"gorm.io/gorm"
)

type (
	OrganizationRepository interface {
		Create(entity.Organization) error
		FindByID(uint) (entity.Organization, error)
		GetOrganizationIDByUserID(userID uint) (uint, error)
		WithTransaction(fn func(tx *gorm.DB) error) error
		CreateWithTx(tx *gorm.DB, org *entity.Organization) error
	}

	OrganizationUsecase interface {
		CreateOrganization(userID uint, req OrganizationCreateReq) error
		GetOrganizationByID(id uint) (entity.Organization, error)
		// 其他方法
	}

	OrganizationCreateReq struct {
		Name    string `json:"name" binding:"required"`
		AdminID uint   `json:"admin_id" binding:"required"`
	}
)
