package models

import (
	"github.com/jinzhu/gorm"
)

// User model using gorm.Model definition
type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     *string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
