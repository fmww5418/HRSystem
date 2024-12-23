package leave

import (
	"HRSystem/src/lib/middleware"
	remployee "HRSystem/src/repository/employee"
	rleave "HRSystem/src/repository/leave"
	uleave "HRSystem/src/usecase/leave"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB, redisClient *redis.Client) {
	leaveRepo := rleave.NewLeaveRepository(db)
	employeeRepo := remployee.NewEmployeeRepository(db)
	leaveUsecase := uleave.NewLeaveUsecase(leaveRepo, employeeRepo)
	handler := NewLeaveHandler(leaveUsecase)

	leaveGroup := router.Group("/leave_requests")
	{
		leaveGroup.Use(middleware.AuthMiddleware(redisClient))
		leaveGroup.GET("", handler.GetAllLeaveRequests)
		leaveGroup.GET("/:id", handler.GetLeaveRequestByID)
		leaveGroup.POST("", handler.CreateLeaveRequest)
		leaveGroup.PUT("/:id/status", handler.UpdateLeaveRequestStatus)
	}
}
