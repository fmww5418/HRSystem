package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	dauth "HRSystem/src/domain/auth"
)

type AuthHandler struct {
	usecase dauth.AuthUsecase
}

func NewAuthHandler(usecase dauth.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		usecase: usecase,
	}
}

// Register
// @Summary User Registration
// @Description Register a new user with username, password, and role
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body dauth.RegisterRequest true "Registration request body"
// @Success 200 {object} map[string]string "Registration successful"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req dauth.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.usecase.Register(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

// Login
// @Summary User Login
// @Description Login with username and password to receive a JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body dauth.LoginRequest true "Login request body"
// @Success 200 {object} map[string]interface{} "Login successful, returns user ID and JWT token"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req dauth.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	token, userID, err := h.usecase.Login(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": userID, "token": token})
}

// Logout
// @Summary User Logout
// @Description Logout by invalidating the JWT token from the server-side
// @Tags Authentication
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token (Format: Bearer 'token')"
// @Success 200 {object} map[string]string "Logout successful"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Router /auth/logout [post]
// @Security ApiKeyAuth
func (h *AuthHandler) Logout(c *gin.Context) {
	userID := c.GetString("user_id")

	if err := h.usecase.Logout(c.Request.Context(), userID); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
