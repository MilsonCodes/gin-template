package models

import (
	"github.com/jinzhu/gorm"
)

// User - model using gorm.Model definition
type User struct {
	Name  string
	Email *string
	gorm.Model
}
