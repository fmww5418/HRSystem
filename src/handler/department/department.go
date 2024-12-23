package department

import (
	"HRSystem/src/domain/department"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type (
	DepartmentHandler struct {
		usecase department.DepartmentUsecase
	}

	departmentResp struct {
		ID             uint      `json:"id"`
		Name           string    `json:"name"`
		OrganizationID uint      `json:"organization_id"`
		CreatedAt      time.Time `json:"created_at"`
	}
)

func NewDepartmentHandler(usecase department.DepartmentUsecase) *DepartmentHandler {
	return &DepartmentHandler{usecase: usecase}
}

// CreateDepartment
// @Summary Create a new department
// @Description Create a new department with the provided details
// @Tags Department
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token (Format: Bearer 'token')"
// @Param request body department.DepartmentCreateReq true "Department create request"
// @Success 201 {object} map[string]string "Department created successfully"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Failed to create department"
// @Router /departments [post]
// @Security ApiKeyAuth
func (h *DepartmentHandler) CreateDepartment(c *gin.Context) {
	var dept department.DepartmentCreateReq
	if err := c.ShouldBindJSON(&dept); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	uID := c.GetString("user_id")
	userID, err := strconv.Atoi(uID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	if err := h.usecase.CreateDepartment(uint(userID), dept); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create department. err: %v", err)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Department created successfully"})
}

// GetAllDepartments
// @Summary Get all departments
// @Description Retrieve a list of all departments
// @Tags Department
// @Produce json
// @Param Authorization header string true "Bearer token (Format: Bearer 'token')"
// @Success 200 {array} departmentResp "List of departments"
// @Failure 500 {object} map[string]string "Failed to fetch departments"
// @Router /departments [get]
// @Security ApiKeyAuth
func (h *DepartmentHandler) GetAllDepartments(c *gin.Context) {
	uID := c.GetString("user_id")
	userID, err := strconv.Atoi(uID)

	departments, err := h.usecase.GetAllDepartments(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch departments"})
		return
	}

	resp := make([]departmentResp, 0, len(departments))
	for _, dept := range departments {
		resp = append(resp, departmentResp{
			ID:             dept.ID,
			Name:           dept.Name,
			OrganizationID: dept.OrganizationID,
			CreatedAt:      dept.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, resp)
}
