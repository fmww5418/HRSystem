package entity

import (
	"time"
	
	"gorm.io/gorm"
)

type Role uint8

const (
	RoleAdmin Role = iota
	RoleEmployee
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Role      Role   `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
