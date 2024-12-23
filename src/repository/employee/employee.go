package employee

import (
	"gorm.io/gorm"

	demployee "HRSystem/src/domain/employee"
	"HRSystem/src/entity"
)

type employeeRepository struct {
	db *gorm.DB
}

var _ demployee.EmployeeRepository = (*employeeRepository)(nil)

func NewEmployeeRepository(db *gorm.DB) demployee.EmployeeRepository {
	return &employeeRepository{db: db}
}

func (r *employeeRepository) FindAll() ([]entity.Employee, error) {
	var employees []entity.Employee
	err := r.db.Find(&employees).Error
	return employees, err
}

func (r *employeeRepository) FindByID(id uint) (entity.Employee, error) {
	var employee entity.Employee
	err := r.db.Preload("User").
		Preload("Department").
		Preload("SupervisorEmployee").
		Preload("Department.Organization").
		First(&employee, id).Error
	return employee, err
}

func (r *employeeRepository) FindByUserID(userID uint) (entity.Employee, error) {
	var employee entity.Employee
	err := r.db.Model(&employee).Preload("Department").Where("user_id = ?", userID).First(&employee).Error
	return employee, err
}

func (r *employeeRepository) Create(employee entity.Employee) error {
	return r.db.Create(&employee).Error
}

func (r *employeeRepository) Update(id uint, employee entity.Employee) error {
	return r.db.Model(&entity.Employee{}).Where("id = ?", id).Updates(employee).Error
}

func (r *employeeRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Employee{}, id).Error
}

func (r *employeeRepository) CreateWithTx(tx *gorm.DB, employee *entity.Employee) error {
	return tx.Create(employee).Error
}

func (r *employeeRepository) UpdateDeptIDWithTxByUserID(tx *gorm.DB, userID uint, deptID uint) error {
	return tx.Model(&entity.Employee{}).Where("user_id = ?", userID).Update("department_id", deptID).Error
}
