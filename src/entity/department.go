package entity

import (
	"time"
	
	"gorm.io/gorm"
)

type Department struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"not null"`
	OrganizationID uint   `gorm:"not null"`
	Organization   Organization `gorm:"foreignKey:OrganizationID;references:ID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
} 