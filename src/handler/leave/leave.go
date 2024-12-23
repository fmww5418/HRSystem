package leave

import (
	dleave "HRSystem/src/domain/leave"
	"HRSystem/src/entity"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	LeaveHandler struct {
		leaveUsecase dleave.LeaveUsecase
	}

	LeaveRequestResp struct {
		ID          uint                 `json:"id"`
		EmployeeID  uint                 `json:"employee_id"`
		StartDate   time.Time            `json:"start_date"`
		EndDate     time.Time            `json:"end_date"`
		Description string               `json:"description"`
		Status      entity.RequestStatus `json:"status"`
	}
)

func NewLeaveHandler(leaveUsecase dleave.LeaveUsecase) *LeaveHandler {
	return &LeaveHandler{
		leaveUsecase: leaveUsecase,
	}
}

// GetAllLeaveRequests
// @Summary Get all leave requests
// @Description Retrieve a list of all leave requests
// @Tags Leave
// @Produce json
// @Param Authorization header string true "Bearer token (Format: Bearer 'token')"
// @Success 200 {array} LeaveRequestResp "List of leave requests"
// @Failure 500 {object} map[string]string "Failed to fetch leave requests"
// @Router /leave_requests [get]
// @Security ApiKeyAuth
func (h *LeaveHandler) GetAllLeaveRequests(c *gin.Context) {
	leaveRequests, err := h.leaveUsecase.GetAllLeaveRequests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch leave requests"})
		return
	}

	leaveRequestResp := make([]LeaveRequestResp, 0, len(leaveRequests))
	for _, leaveRequest := range leaveRequests {
		leaveRequestResp = append(leaveRequestResp, LeaveRequestResp{
			ID:          leaveRequest.ID,
			EmployeeID:  leaveRequest.EmployeeID,
			StartDate:   leaveRequest.StartDate,
			EndDate:     leaveRequest.EndDate,
			Description: leaveRequest.Description,
			Status:      leaveRequest.Status,
		})
	}

	c.JSON(http.StatusOK, leaveRequestResp)
}

// GetLeaveRequestByID
// @Summary Get leave request by ID
// @Description Retrieve a leave request by its ID
// @Tags Leave
// @Produce json
// @Param Authorization header string true "Bearer token (Format: Bearer 'token')"
// @Param id path int true "Leave request ID"
// @Success 200 {object} entity.Request "Leave request details"
// @Failure 400 {object} map[string]string "Invalid leave request ID"
// @Failure 404 {object} map[string]string "Leave request not found"
// @Router /leave_requests/{id} [get]
// @Security ApiKeyAuth
func (h *LeaveHandler) GetLeaveRequestByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid leave request ID"})
		return
	}

	leaveRequest, err := h.leaveUsecase.GetLeaveRequestByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Leave request not found"})
		return
	}

	c.JSON(http.StatusOK, LeaveRequestResp{
		ID:          leaveRequest.ID,
		EmployeeID:  leaveRequest.EmployeeID,
		StartDate:   leaveRequest.StartDate,
		EndDate:     leaveRequest.EndDate,
		Description: leaveRequest.Description,
		Status:      leaveRequest.Status,
	})
}

// CreateLeaveRequest
// @Summary Create a new leave request
// @Description Create a new leave request with the provided details
// @Tags Leave
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token (Format: Bearer 'token')"
// @Param request body dleave.LeaveRequestInput true "Leave request create request"
// @Success 201 {object} map[string]string "Leave request created successfully"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Failed to create leave request"
// @Router /leave_requests [post]
// @Security ApiKeyAuth
func (h *LeaveHandler) CreateLeaveRequest(c *gin.Context) {
	var req dleave.LeaveRequestInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	uID := c.GetString("user_id")
	userID, err := strconv.Atoi(uID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := h.leaveUsecase.CreateLeaveRequest(uint(userID), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create leave request. error: %v", err)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Leave request created successfully"})
}

// UpdateLeaveRequestStatus
// @Summary Update leave request status
// @Description Update the status of a leave request
// @Tags Leave
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token (Format: Bearer 'token')"
// @Param id path int true "Leave request ID"
// @Param request body dleave.UpdateLeaveStatusInput true "Leave request status update request"
// @Success 200 {object} map[string]string "Leave request status updated successfully"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Failed to update leave request status"
// @Router /leave_requests/{id}/status [put]
// @Security ApiKeyAuth
func (h *LeaveHandler) UpdateLeaveRequestStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid leave request ID"})
		return
	}

	var req dleave.UpdateLeaveStatusInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userID := c.GetString("user_id")
	operatorUserID, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
	}

	if err := h.leaveUsecase.UpdateLeaveRequestStatus(uint(id), uint(operatorUserID), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to update leave request status. err: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Leave request status updated successfully"})
}
