package department

import (
	"HRSystem/src/domain/department"
	ddep "HRSystem/src/domain/department"
	"HRSystem/src/domain/employee"
	"HRSystem/src/domain/organization"
	"HRSystem/src/entity"
	"errors"
	"fmt"
	"time"
)

type departmentUsecase struct {
	deptRepo     department.DepartmentRepository
	orgRepo      organization.OrganizationRepository
	employeeRepo employee.EmployeeRepository
}

var _ ddep.DepartmentUsecase = (*departmentUsecase)(nil)

func NewDepartmentUsecase(
	deptRepo department.DepartmentRepository,
	orgRepo organization.OrganizationRepository,
	employeeRepo employee.EmployeeRepository,
) department.DepartmentUsecase {
	return &departmentUsecase{deptRepo: deptRepo, employeeRepo: employeeRepo, orgRepo: orgRepo}
}

func (u *departmentUsecase) CreateDepartment(userID uint, req department.DepartmentCreateReq) error {
	org, err := u.orgRepo.FindByID(req.OrganizationID)
	if err != nil {
		return err
	}

	if org.AdminID != userID {
		return fmt.Errorf("you are not allowed to create a department")
	}

	return u.deptRepo.Create(entity.Department{
		Name:           req.Name,
		OrganizationID: req.OrganizationID,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	})
}

func (u *departmentUsecase) Invite(dept entity.Department, userID uint) error {
	receiver, err := u.employeeRepo.FindByUserID(userID)
	if err != nil {
		return fmt.Errorf("cannot find the employee by userID: %v", userID)
	}

	if receiver.DepartmentID != nil {
		return errors.New("department already exists")
	}

	if *receiver.DepartmentID == dept.ID {
		return errors.New("department already joined")
	}

	receiver.DepartmentID = &dept.ID
	if err = u.employeeRepo.Update(receiver.ID, receiver); err != nil {
		return err
	}

	return nil
}

func (u *departmentUsecase) GetAllDepartments(userID uint) ([]entity.Department, error) {
	employeeUser, err := u.employeeRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	if employeeUser.DepartmentID == nil {
		return nil, errors.New("can't get organization id because the user has no department")
	}

	return u.deptRepo.FindAllByOrgID(employeeUser.Department.OrganizationID)
}
