package employee

import (
	"errors"
	"fmt"
	"time"

	ddept "HRSystem/src/domain/department"
	demployee "HRSystem/src/domain/employee"
	"HRSystem/src/entity"
)

type (
	employeeUsecase struct {
		employeeRepo demployee.EmployeeRepository
		deptRepo     ddept.DepartmentRepository
	}
)

var _ demployee.EmployeeUsecase = (*employeeUsecase)(nil)

func NewEmployeeUsecase(employeeRepo demployee.EmployeeRepository, deptRepo ddept.DepartmentRepository) demployee.EmployeeUsecase {
	return &employeeUsecase{employeeRepo: employeeRepo, deptRepo: deptRepo}
}

func (s *employeeUsecase) GetAllEmployees(operatorUserID uint) ([]entity.Employee, error) {
	return s.employeeRepo.FindAllByUserID(operatorUserID)
}

func (s *employeeUsecase) GetEmployeeByID(operatorUserID uint, id uint) (entity.Employee, error) {
	return s.employeeRepo.FindByIDWithOrgCheck(operatorUserID, id)
}

func (s *employeeUsecase) CreateEmployee(operatorUserID uint, req demployee.EmployeeRequest) error {
	if _, err := s.employeeRepo.FindByUserID(req.UserID); err == nil {
		return fmt.Errorf("employee with UserID %d already exists", req.UserID)
	}

	if req.DepartmentID != nil {
		operator, err := s.employeeRepo.FindByUserID(operatorUserID)
		if err != nil {
			return err
		}

		dept, err := s.deptRepo.FindByID(*req.DepartmentID)
		if err != nil {
			return err
		}

		if operator.Department == nil || dept.OrganizationID != operator.Department.OrganizationID {
			return fmt.Errorf("you are not allowed to create an employee with an outside organization. departmentID: %v", *req.DepartmentID)
		}

		if req.SupervisorEmployeeID != nil {
			supervisor, err := s.employeeRepo.FindByID(*req.SupervisorEmployeeID)
			if err != nil {
				return err
			}

			if supervisor.Department == nil || supervisor.Department.OrganizationID != dept.OrganizationID {
				return fmt.Errorf("you are not allowed to create an supervisor with an outside organization")
			}
		}
	}

	if req.DepartmentID == nil && req.SupervisorEmployeeID != nil {
		return fmt.Errorf("you can't assign a supervisor to the employee without department id")
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
		Salary:               *req.Salary,
		DepartmentID:         req.DepartmentID,
		SupervisorEmployeeID: req.SupervisorEmployeeID,
		RemainedDayOff:       remainedDayOff,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}

	return s.employeeRepo.Create(employee)
}

func (s *employeeUsecase) UpdateEmployee(operatorUserID uint, id uint, req demployee.EmployeeRequest) error {
	if _, err := s.employeeRepo.FindByIDWithOrgCheck(operatorUserID, id); err != nil {
		return err
	}

	if req.SupervisorEmployeeID != nil {
		employee, err := s.employeeRepo.FindByID(id)
		if err != nil {
			return err
		}

		supervisor, err := s.employeeRepo.FindByID(*req.SupervisorEmployeeID)
		if err != nil {
			return fmt.Errorf("can't find the supervisor: %v", err)
		}

		if employee.Department == nil || employee.Department.OrganizationID != supervisor.Department.OrganizationID {
			return errors.New("OrganizationID is not equal to SupervisorOrganizationID")
		}
	}

	return s.employeeRepo.Update(id, entity.Employee{
		Name:                 req.Name,
		Position:             req.Position,
		ContactInfo:          req.ContactInfo,
		Salary:               *req.Salary,
		RemainedDayOff:       *req.RemainedDayOff,
		SupervisorEmployeeID: req.SupervisorEmployeeID,
		UpdatedAt:            time.Now(),
	})
}

func (s *employeeUsecase) DeleteEmployee(operatorUserID uint, id uint) error {
	if _, err := s.employeeRepo.FindByIDWithOrgCheck(operatorUserID, id); err != nil {
		return err
	}

	return s.employeeRepo.Delete(id)
}
