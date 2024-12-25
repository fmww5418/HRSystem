package employee

import (
	"HRSystem/src/lib/middleware"
	rdept "HRSystem/src/repository/department"
	remployee "HRSystem/src/repository/employee"
	uemployee "HRSystem/src/usecase/employee"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB, redisClient *redis.Client) {
	employeeRepo := remployee.NewEmployeeRepository(db)
	deptRepo := rdept.NewDepartmentRepository(db)
	usecase := uemployee.NewEmployeeUsecase(employeeRepo, deptRepo)
	handler := NewEmployeeHandler(usecase)

	employeeGroup := router.Group("/employees")
	{
		employeeGroup.Use(middleware.AuthMiddleware(redisClient))
		employeeGroup.GET("", handler.GetAllEmployees)
		employeeGroup.GET("/:id", handler.GetEmployeeByID)
		employeeGroup.POST("", handler.CreateEmployee)
		employeeGroup.PUT("/:id", handler.UpdateEmployee)
		employeeGroup.DELETE("/:id", handler.DeleteEmployee)
	}
}
