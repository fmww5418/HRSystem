package employee

import (
	"fmt"
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

func (r *employeeRepository) FindAllByUserID(operatorUserID uint) ([]entity.Employee, error) {
	var employees []entity.Employee
	var operator entity.Employee

	if err := r.db.Preload("Department").
		Where("user_id = ?", operatorUserID).First(&operator).Error; err != nil {
		return employees, err
	}

	if operator.Department == nil {
		return employees, fmt.Errorf("can't find employees")
	}

	err := r.db.Model(&entity.Organization{}).
		Select("employees.*").
		Joins("INNER JOIN departments ON departments.organization_id = organizations.id").
		Joins("INNER JOIN employees ON employees.department_id = departments.id").
		Where("organizations.id = ? AND employees.deleted_at IS NULL AND departments.deleted_at IS NULL", operator.Department.OrganizationID).
		Find(&employees).Error

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

func (r *employeeRepository) FindByIDWithOrgCheck(operatorUserID uint, id uint) (entity.Employee, error) {
	operator, err := r.FindByUserID(operatorUserID)
	if err != nil {
		return entity.Employee{}, err
	}

	employeeUser, err := r.FindByID(id)
	if err != nil {
		return entity.Employee{}, err
	}

	if employeeUser.Department == nil ||
		operator.Department == nil ||
		employeeUser.Department.OrganizationID != operator.Department.OrganizationID {
		return entity.Employee{}, fmt.Errorf("you are not allowed to operate an employee with an outside organization")
	}

	return employeeUser, nil
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
