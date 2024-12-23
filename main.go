package main

import (
	"HRSystem/src/handler/department"
	"HRSystem/src/handler/organization"
	"fmt"

	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	"HRSystem/config"
	_ "HRSystem/docs"
	"HRSystem/src/handler/auth"
	"HRSystem/src/handler/employee"
	"HRSystem/src/handler/leave"
	"HRSystem/src/lib/db"
	rlib "HRSystem/src/lib/redis"
)

var (
	gormDB *gorm.DB
)

func main() {
	// Load configuration
	config.LoadConfig(nil)

	// Initialize MySQL database
	var err error
	gormDB, err = db.Init()
	if err != nil {
		log.Fatalf("Failed to initialize MySQL: %v", err)
	}

	// Initialize Redis client
	redisClient := rlib.Init()

	// Setup Gin router
	router := gin.Default()

	// Swagger documentation route
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Apply middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Register routes
	auth.RegisterRoutes(router, gormDB, redisClient)
	employee.RegisterRoutes(router, gormDB, redisClient)
	leave.RegisterRoutes(router, gormDB, redisClient)
	organization.RegisterOrganizationRoutes(router, gormDB, redisClient)
	department.RegisterDepartmentRoutes(router, gormDB, redisClient)

	// Start the server
	serverAddr := fmt.Sprintf(":%d", config.Config.ServerPort)
	log.Printf("Server running on %s", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
