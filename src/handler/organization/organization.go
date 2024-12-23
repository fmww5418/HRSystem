package organization

import (
	"HRSystem/src/domain/organization"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type (
	OrganizationHandler struct {
		usecase organization.OrganizationUsecase
	}

	organizationResp struct {
		ID        uint      `json:"id"`
		Name      string    `json:"name"`
		AdminID   uint      `json:"admin_id"`
		CreatedAt time.Time `json:"created_at"`
	}
)

func NewOrganizationHandler(usecase organization.OrganizationUsecase) *OrganizationHandler {
	return &OrganizationHandler{usecase: usecase}
}

// CreateOrganization
// @Summary Create a new organization
// @Description Create a new organization with the provided details
// @Tags Organization
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token (Format: Bearer 'token')"
// @Param request body organization.OrganizationCreateReq true "Organization create request"
// @Success 201 {object} map[string]string "Organization created successfully"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Failed to create organization"
// @Router /organizations [post]
// @Security ApiKeyAuth
func (h *OrganizationHandler) CreateOrganization(c *gin.Context) {
	var req organization.OrganizationCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userID := c.GetString("user_id")
	operatorUserID, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}

	if err := h.usecase.CreateOrganization(uint(operatorUserID), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create organization"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Organization created successfully"})
}

// GetOrganizationByMe
// @Summary Get organization by myself
// @Description Retrieve an organization by myself
// @Tags Organization
// @Produce json
// @Param Authorization header string true "Bearer token (Format: Bearer 'token')"
// @Success 200 {object} organizationResp "Organization details"
// @Failure 400 {object} map[string]string "Invalid organization ID"
// @Failure 404 {object} map[string]string "Organization not found"
// @Failure 500 {object} map[string]string "Failed to fetch organization"
// @Router /organizations/me [get]
// @Security ApiKeyAuth
func (h *OrganizationHandler) GetOrganizationByMe(c *gin.Context) {
	uID := c.GetString("user_id")
	userID, err := strconv.Atoi(uID)
	org, err := h.usecase.GetOrganizationByID(uint(userID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Organization not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch organization"})
		}
		return
	}

	c.JSON(http.StatusOK, organizationResp{
		ID:        org.ID,
		Name:      org.Name,
		AdminID:   org.AdminID,
		CreatedAt: org.CreatedAt,
	})
}
