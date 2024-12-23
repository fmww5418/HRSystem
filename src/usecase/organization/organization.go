package organization

import (
	"HRSystem/src/domain/department"
	"HRSystem/src/domain/employee"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"

	"HRSystem/src/domain/organization"
	dorg "HRSystem/src/domain/organization"
	"HRSystem/src/entity"
)

type organizationUsecase struct {
	orgRepo      organization.OrganizationRepository
	deptRepo     department.DepartmentRepository
	employeeRepo employee.EmployeeRepository
}

var _ dorg.OrganizationUsecase = (*organizationUsecase)(nil)

func NewOrganizationUsecase(
	orgRepo organization.OrganizationRepository,
	deptRepo department.DepartmentRepository,
	employeeRepo employee.EmployeeRepository,
) organization.OrganizationUsecase {
	return &organizationUsecase{orgRepo: orgRepo, deptRepo: deptRepo, employeeRepo: employeeRepo}
}

func (u *organizationUsecase) CreateOrganization(userID uint, req dorg.OrganizationCreateReq) error {
	orgID, err := u.orgRepo.GetOrganizationIDByUserID(userID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if orgID != 0 {
		return fmt.Errorf("you are not allowed to create a department")
	}

	org := entity.Organization{
		Name:      req.Name,
		AdminID:   req.AdminID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = u.orgRepo.WithTransaction(func(tx *gorm.DB) error {
		if err = u.orgRepo.CreateWithTx(tx, &org); err != nil {
			return err
		}

		dept := entity.Department{
			OrganizationID: org.ID,
			Name:           "Admin Department",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}
		if err = u.deptRepo.CreateWithTx(tx, &dept); err != nil {
			return err
		}

		if err = u.employeeRepo.UpdateDeptIDWithTxByUserID(tx, userID, dept.ID); err != nil {
			return err
		}

		return nil
	})

	return err
}

func (u *organizationUsecase) GetOrganizationByID(userID uint) (entity.Organization, error) {
	employeeUser, err := u.employeeRepo.FindByUserID(userID)
	if err != nil {
		return entity.Organization{}, err
	}

	if employeeUser.DepartmentID == nil {
		return entity.Organization{}, errors.New("can't get organization id because the user has no department")
	}

	return u.orgRepo.FindByID(employeeUser.Department.OrganizationID)
}
