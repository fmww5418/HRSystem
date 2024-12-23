package department

import (
	"gorm.io/gorm"

	ddep "HRSystem/src/domain/department"
	"HRSystem/src/entity"
)

type departmentRepository struct {
	db *gorm.DB
}

var _ ddep.DepartmentRepository = (*departmentRepository)(nil)

func NewDepartmentRepository(db *gorm.DB) ddep.DepartmentRepository {
	return &departmentRepository{db: db}
}

func (r *departmentRepository) Create(dept entity.Department) error {
	return r.db.Create(&dept).Error
}

func (r *departmentRepository) FindByID(id uint) (entity.Department, error) {
	var dep entity.Department
	err := r.db.First(&dep, id).Error
	return dep, err
}

func (r *departmentRepository) FindAllByOrgID(orgID uint) ([]entity.Department, error) {
	var departments []entity.Department
	err := r.db.Where("organization_id = ?", orgID).Find(&departments).Error
	return departments, err
}

func (r *departmentRepository) CreateWithTx(tx *gorm.DB, dept *entity.Department) error {
	return tx.Create(dept).Error
}
