package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // Import for sqlite3
)

// DB - the global variable for the gorm database
var DB *gorm.DB

// ConnectDataBase - function to connect to simple sqlite database
func ConnectDataBase() {
	database, err := gorm.Open("sqlite3", "database/test.db")
	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&User{})

	DB = database
}
