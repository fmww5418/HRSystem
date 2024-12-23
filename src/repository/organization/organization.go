package organization

import (
	dorg "HRSystem/src/domain/organization"
	"HRSystem/src/entity"
	"gorm.io/gorm"
)

type organizationRepository struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) dorg.OrganizationRepository {
	return &organizationRepository{db: db}
}

func (r *organizationRepository) Create(org entity.Organization) error {
	return r.db.Create(&org).Error
}

func (r *organizationRepository) FindByID(id uint) (entity.Organization, error) {
	var org entity.Organization
	err := r.db.First(&org, id).Error
	return org, err
}

func (r *organizationRepository) GetOrganizationIDByUserID(userID uint) (uint, error) {
	var org entity.Organization
	err := r.db.Model(&org).
		Select("organizations.id").
		Joins("LEFT JOIN departments ON departments.organization_id = organizations.id").
		Joins("LEFT JOIN employees ON employees.department_id = departments.id").
		Where("employees.user_id = ? AND employees.deleted_at IS NULL AND departments.deleted_at IS NULL", userID).
		Or("organizations.admin_id = ?", userID).First(&org).Error
	return org.ID, err
}

func (r *organizationRepository) WithTransaction(fn func(tx *gorm.DB) error) error {
	return r.db.Transaction(fn)
}

func (r *organizationRepository) CreateWithTx(tx *gorm.DB, org *entity.Organization) error {
	return tx.Create(org).Error
}
