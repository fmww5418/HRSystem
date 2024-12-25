package employee

import (
	"HRSystem/src/entity"
	"gorm.io/gorm"
)

type (
	EmployeeRepository interface {
		FindAllByUserID(operatorUserID uint) ([]entity.Employee, error)
		FindByID(id uint) (entity.Employee, error)
		FindByIDWithOrgCheck(userID uint, id uint) (entity.Employee, error)
		FindByUserID(id uint) (entity.Employee, error)
		Create(employee entity.Employee) error
		Update(id uint, employee entity.Employee) error
		Delete(id uint) error
		CreateWithTx(tx *gorm.DB, employee *entity.Employee) error
		UpdateDeptIDWithTxByUserID(tx *gorm.DB, id uint, deptID uint) error
	}

	EmployeeUsecase interface {
		GetAllEmployees(operatorUserID uint) ([]entity.Employee, error)
		GetEmployeeByID(operatorUserID uint, id uint) (entity.Employee, error)
		CreateEmployee(operatorUserID uint, req EmployeeRequest) error
		UpdateEmployee(operatorUserID uint, id uint, req EmployeeRequest) error
		DeleteEmployee(operatorUserID uint, id uint) error
	}
)

type EmployeeRequest struct {
	Name                 string   `json:"name" binding:"required"`
	Position             string   `json:"position" binding:"required"`
	ContactInfo          string   `json:"contact_info" binding:"required"`
	UserID               uint     `json:"user_id" binding:"required"`
	Salary               *float64 `json:"salary" binding:"required"`
	DepartmentID         *uint    `json:"department_id" `
	SupervisorEmployeeID *uint    `json:"supervisor_employee_id"`
	RemainedDayOff       *int     `json:"remained_day_off" binding:"required"`
}
