package department

import (
	"HRSystem/src/lib/middleware"
	rdpet "HRSystem/src/repository/department"
	remployee "HRSystem/src/repository/employee"
	rorg "HRSystem/src/repository/organization"
	udept "HRSystem/src/usecase/department"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func RegisterDepartmentRoutes(router *gin.Engine, db *gorm.DB, redisClient *redis.Client) {
	deptRepo := rdpet.NewDepartmentRepository(db)
	employeeRepo := remployee.NewEmployeeRepository(db)
	orgRepo := rorg.NewOrganizationRepository(db)
	usecase := udept.NewDepartmentUsecase(deptRepo, orgRepo, employeeRepo)
	handler := NewDepartmentHandler(usecase)

	deptGroup := router.Group("/departments")
	{
		deptGroup.Use(middleware.AuthMiddleware(redisClient))
		deptGroup.POST("", handler.CreateDepartment)
		deptGroup.GET("", handler.GetAllDepartments)
		// 其他路由
	}
}
