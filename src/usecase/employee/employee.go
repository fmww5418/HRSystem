package employee

import (
	"errors"
	"fmt"
	"time"

	demployee "HRSystem/src/domain/employee"
	"HRSystem/src/entity"
)

type (
	employeeUsecase struct {
		repo demployee.EmployeeRepository
	}
)

var _ demployee.EmployeeUsecase = (*employeeUsecase)(nil)

func NewEmployeeService(repo demployee.EmployeeRepository) demployee.EmployeeUsecase {
	return &employeeUsecase{repo: repo}
}

func (s *employeeUsecase) GetAllEmployees() ([]entity.Employee, error) {
	return s.repo.FindAll()
}

func (s *employeeUsecase) GetEmployeeByID(id uint) (entity.Employee, error) {
	return s.repo.FindByID(id)
}

func (s *employeeUsecase) CreateEmployee(req demployee.EmployeeRequest) error {
	if _, err := s.repo.FindByUserID(req.UserID); err == nil {
		return fmt.Errorf("employee with UserID %d already exists", req.UserID)
	}

	remainedDayOff := 0
	if req.RemainedDayOff != nil {
		remainedDayOff = *req.RemainedDayOff
	}

	employee := entity.Employee{
		Name:                 req.Name,
		Position:             req.Position,
		ContactInfo:          req.ContactInfo,
		UserID:               req.UserID,
		Salary:               req.Salary,
		DepartmentID:         req.DepartmentID,
		SupervisorEmployeeID: req.SupervisorEmployeeID,
		RemainedDayOff:       remainedDayOff,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}

	return s.repo.Create(employee)
}

func (s *employeeUsecase) UpdateEmployee(id uint, req demployee.EmployeeRequest) error {
	if req.SupervisorEmployeeID != nil {
		employee, err := s.repo.FindByID(id)
		if err != nil {
			return err
		}

		supervisor, err := s.repo.FindByID(*req.SupervisorEmployeeID)
		if err != nil {
			return err
		}

		if employee.Department == nil || employee.Department.OrganizationID != supervisor.Department.OrganizationID {
			return errors.New("OrganizationID is not equal to SupervisorOrganizationID")
		}
	}

	return s.repo.Update(id, entity.Employee{
		Name:                 req.Name,
		Position:             req.Position,
		ContactInfo:          req.ContactInfo,
		Salary:               req.Salary,
		RemainedDayOff:       *req.RemainedDayOff,
		SupervisorEmployeeID: req.SupervisorEmployeeID,
		UpdatedAt:            time.Now(),
	})
}

func (s *employeeUsecase) DeleteEmployee(id uint) error {
	return s.repo.Delete(id)
}
