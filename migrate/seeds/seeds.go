package seeds

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"HRSystem/src/entity"
	ltime "HRSystem/src/lib/time"
)

func Case1(db *gorm.DB) {
	log.Println("Seeding database with initial data...")

	// Create users
	password, _ := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)

	admin := entity.User{
		Username:  "admin",
		Password:  string(password),
		Role:      entity.RoleAdmin,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.Create(&admin).Error; err != nil {
		log.Fatalf("Failed to seeds admin user: %v", err)
	}

	org := entity.Organization{
		Name:      "Super Organization",
		AdminID:   admin.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.Create(&org).Error; err != nil {
		log.Fatalf("Failed to seeds org: %v", err)
	}

	deptAdmin := entity.Department{
		Name:           "Admin Department",
		OrganizationID: org.ID,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	deptTech := entity.Department{
		Name:           "Tech",
		OrganizationID: org.ID,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := db.Create(&deptAdmin).Error; err != nil {
		log.Fatalf("Failed to seeds department: %v", err)
	}
	if err := db.Create(&deptTech).Error; err != nil {
		log.Fatalf("Failed to seeds department: %v", err)
	}

	userJohn := entity.User{
		Username:  "john",
		Password:  string(password),
		Role:      entity.RoleEmployee,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	userPeter := entity.User{
		Username:  "peter",
		Password:  string(password),
		Role:      entity.RoleEmployee,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.Create(&userJohn).Error; err != nil {
		log.Fatalf("Failed to seeds employee user: %v", err)
	}

	if err := db.Create(&userPeter).Error; err != nil {
		log.Fatalf("Failed to seeds employee user: %v", err)
	}

	// Create employee data
	empAdmin := entity.Employee{
		Name:         "Boss",
		Position:     "CEO",
		ContactInfo:  "admin@example.com",
		Salary:       15,
		UserID:       admin.ID,
		DepartmentID: &deptAdmin.ID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	empJohn := entity.Employee{
		Name:                 "John Doe",
		Position:             "Engineer Manager",
		ContactInfo:          "john@example.com",
		Salary:               500000,
		UserID:               userJohn.ID,
		DepartmentID:         &deptTech.ID,
		SupervisorEmployeeID: &admin.ID,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}

	empPeter := entity.Employee{
		Name:                 "Peter Wu",
		Position:             "Engineer",
		ContactInfo:          "peter@example.com",
		Salary:               100000,
		UserID:               userPeter.ID,
		DepartmentID:         &deptTech.ID,
		SupervisorEmployeeID: &empJohn.ID,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}

	if err := db.Create(&empAdmin).Error; err != nil {
		log.Fatalf("Failed to seeds employee: %v", err)
	}
	if err := db.Create(&empJohn).Error; err != nil {
		log.Fatalf("Failed to seeds employee: %v", err)
	}
	if err := db.Create(&empPeter).Error; err != nil {
		log.Fatalf("Failed to seeds employee: %v", err)
	}

	// Create leave request
	leaveReq := entity.Request{
		EmployeeID:  empPeter.ID,
		StartDate:   ltime.ParseTime(time.DateTime, "2024-01-04 09:00:00", nil),
		EndDate:     ltime.ParseTime(time.DateTime, "2024-01-06 17:00:00", nil),
		Status:      entity.RequestStatusPending,
		Description: "Vacation",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := db.Create(&leaveReq).Error; err != nil {
		log.Fatalf("Failed to seeds leave request: %v", err)
	}

	// Create leave request
	leaveReq2 := entity.Request{
		EmployeeID:  empPeter.ID,
		StartDate:   ltime.ParseTime(time.DateTime, "2024-01-09 12:00:00", nil),
		EndDate:     ltime.ParseTime(time.DateTime, "2024-01-12 14:00:00", nil),
		Status:      entity.RequestStatusPending,
		Description: "Vacation2",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := db.Create(&leaveReq2).Error; err != nil {
		log.Fatalf("Failed to seeds leave request: %v", err)
	}

	log.Println("Database seeding completed successfully.")
}
