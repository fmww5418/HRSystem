package employee

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"

	demployee "HRSystem/src/domain/employee"
)

type (
	EmployeeHandler struct {
		usecase demployee.EmployeeUsecase
	}

	EmployeeResp struct {
		ID                   uint    `json:"id"`
		Name                 string  `json:"name"`
		Position             string  `json:"position"`
		ContactInfo          string  `json:"contact_info"`
		Salary               float64 `json:"salary"`
		UserID               uint    `json:"user_id"`
		DepartmentID         *uint   `json:"department_id"`
		SupervisorEmployeeID *uint   `json:"supervisor_employee_id"`
		RemainedDayOff       int     `json:"remained_day_off"`
	}
)

func NewEmployeeHandler(usecase demployee.EmployeeUsecase) *EmployeeHandler {
	return &EmployeeHandler{usecase: usecase}
}

// GetAllEmployees
// @Summary Get all employees
// @Description Retrieve a list of all employees
// @Tags Employee
// @Produce json
// @Param Authorization header string true "Bearer token (Format: Bearer 'token')"
// @Success 200 {array} EmployeeResp "List of employees"
// @Failure 500 {object} map[string]string "Failed to fetch employees"
// @Router /employees [get]
// @Security ApiKeyAuth
func (h *EmployeeHandler) GetAllEmployees(c *gin.Context) {
	userID := c.GetString("user_id")
	operatorUserID, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	employees, err := h.usecase.GetAllEmployees(uint(operatorUserID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch employees"})
		return
	}

	employeeResp := make([]EmployeeResp, 0, len(employees))
	for _, employee := range employees {
		employeeResp = append(employeeResp, EmployeeResp{
			ID:                   employee.ID,
			Name:                 employee.Name,
			Position:             employee.Position,
			ContactInfo:          employee.ContactInfo,
			Salary:               employee.Salary,
			UserID:               employee.UserID,
			DepartmentID:         employee.DepartmentID,
			SupervisorEmployeeID: employee.SupervisorEmployeeID,
			RemainedDayOff:       employee.RemainedDayOff,
		})
	}

	c.JSON(http.StatusOK, employeeResp)
}

// GetEmployeeByID
// @Summary Get employee by ID
// @Description Retrieve an employee by their ID
// @Tags Employee
// @Produce json
// @Param Authorization header string true "Bearer token (Format: Bearer 'token')"
// @Param id path int true "Employee ID"
// @Success 200 {object} EmployeeResp "Employee details"
// @Failure 400 {object} map[string]string "Invalid employee ID"
// @Failure 404 {object} map[string]string "Employee not found"
// @Failure 500 {object} map[string]string "Failed to fetch employee"
// @Router /employees/{id} [get]
// @Security ApiKeyAuth
func (h *EmployeeHandler) GetEmployeeByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	userID := c.GetString("user_id")
	operatorUserID, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	employee, err := h.usecase.GetEmployeeByID(uint(operatorUserID), uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Failed to fetch employee. error: %v", err)})
		}
		return
	}

	c.JSON(http.StatusOK, EmployeeResp{
		ID:                   employee.ID,
		Name:                 employee.Name,
		Position:             employee.Position,
		ContactInfo:          employee.ContactInfo,
		Salary:               employee.Salary,
		UserID:               employee.UserID,
		DepartmentID:         employee.DepartmentID,
		SupervisorEmployeeID: employee.SupervisorEmployeeID,
		RemainedDayOff:       employee.RemainedDayOff,
	})
}

// CreateEmployee
// @Summary Create a new employee
// @Description Create a new employee with the provided details
// @Tags Employee
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token (Format: Bearer 'token')"
// @Param request body demployee.EmployeeRequest true "Employee create request"
// @Success 201 {object} map[string]string "Employee created successfully"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Failed to create employee"
// @Router /employees [post]
// @Security ApiKeyAuth
func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
	var req demployee.EmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userID := c.GetString("user_id")
	operatorUserID, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	if err := h.usecase.CreateEmployee(uint(operatorUserID), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create employee"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Employee created successfully"})
}

// UpdateEmployee
// @Summary Update an employee
// @Description Update an employee's details
// @Tags Employee
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Param Authorization header string true "Bearer token (Format: Bearer 'token')"
// @Param request body demployee.EmployeeRequest true "Employee update request"
// @Success 200 {object} map[string]string "Employee updated successfully"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Failed to update employee"
// @Router /employees/{id} [put]
// @Security ApiKeyAuth
func (h *EmployeeHandler) UpdateEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	var req demployee.EmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userID := c.GetString("user_id")
	operatorUserID, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	if err := h.usecase.UpdateEmployee(uint(operatorUserID), uint(id), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to update employee. error: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee updated successfully"})
}

// DeleteEmployee
// @Summary Delete an employee
// @Description Delete an employee by their ID
// @Tags Employee
// @Produce json
// @Param Authorization header string true "Bearer token (Format: Bearer 'token')"
// @Param id path int true "Employee ID"
// @Success 200 {object} map[string]string "Employee deleted successfully"
// @Failure 400 {object} map[string]string "Invalid employee ID"
// @Failure 500 {object} map[string]string "Failed to delete employee"
// @Router /employees/{id} [delete]
// @Security ApiKeyAuth
func (h *EmployeeHandler) DeleteEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	userID := c.GetString("user_id")
	operatorUserID, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	if err := h.usecase.DeleteEmployee(uint(operatorUserID), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete employee"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}
