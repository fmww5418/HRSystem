package auth

import (
	"HRSystem/src/lib/middleware"
	rauth "HRSystem/src/repository/auth"
	remployee "HRSystem/src/repository/employee"
	uauth "HRSystem/src/usecase/auth"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB, redisClient *redis.Client) {
	authRepo := rauth.NewAuthRepository(db)
	employeeRepo := remployee.NewEmployeeRepository(db)
	usecase := uauth.NewAuthUsecase(authRepo, employeeRepo, redisClient)
	handler := NewAuthHandler(usecase)

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", handler.Register)
		authGroup.POST("/login", handler.Login)

		authGroup.Use(middleware.AuthMiddleware(redisClient))
		authGroup.POST("/logout", handler.Logout)
	}
}
