package entity

import "time"

type RequestStatus uint

const (
	RequestStatusPending RequestStatus = iota
	RequestStatusAccepted
	RequestStatusRejected
	RequestStatusMax
)

type Request struct {
	ID          uint          `gorm:"primaryKey"`
	EmployeeID  uint          `gorm:"not null"`
	Employee    Employee      `gorm:"foreignKey:EmployeeID;references:ID"`
	StartDate   time.Time     `gorm:"not null;index"`
	EndDate     time.Time     `gorm:"not null"`
	Status      RequestStatus `gorm:"default:0;index"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
