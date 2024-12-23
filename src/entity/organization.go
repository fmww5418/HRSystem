package entity

import (
	"time"
	
	"gorm.io/gorm"
)

type Organization struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique;not null"`
	AdminID   uint   `gorm:"not null"`
	Admin     User   `gorm:"foreignKey:AdminID;references:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
} 