package organization

import (
	"HRSystem/src/lib/middleware"
	rdept "HRSystem/src/repository/department"
	remployee "HRSystem/src/repository/employee"
	rorg "HRSystem/src/repository/organization"
	uorg "HRSystem/src/usecase/organization"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func RegisterOrganizationRoutes(router *gin.Engine, db *gorm.DB, redisClient *redis.Client) {
	orgRepo := rorg.NewOrganizationRepository(db)
	deptRepo := rdept.NewDepartmentRepository(db)
	employeeRepo := remployee.NewEmployeeRepository(db)
	usecase := uorg.NewOrganizationUsecase(orgRepo, deptRepo, employeeRepo)
	handler := NewOrganizationHandler(usecase)

	orgGroup := router.Group("/organizations")
	{
		orgGroup.Use(middleware.AuthMiddleware(redisClient))
		orgGroup.POST("", handler.CreateOrganization)
		orgGroup.GET("/me", handler.GetOrganizationByMe)
	}
}
