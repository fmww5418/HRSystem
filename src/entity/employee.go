package entity

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	ID                   uint   `gorm:"primaryKey"`
	Name                 string `gorm:"not null"`
	Position             string `gorm:"not null"`
	ContactInfo          string `gorm:"not null"`
	Salary               float64
	UserID               uint        `gorm:"unique;not null;"`
	User                 User        `gorm:"foreignKey:UserID;references:ID"`
	DepartmentID         *uint       `gorm:""`
	Department           *Department `gorm:"foreignKey:DepartmentID;references:ID"`
	SupervisorEmployeeID *uint       `gorm:""`
	SupervisorEmployee   *Employee   `gorm:"foreignKey:SupervisorEmployeeID;references:ID"`
	RemainedDayOff       int         `gorm:"default:0"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            gorm.DeletedAt `gorm:"index"`
}
