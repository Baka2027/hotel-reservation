package config

import (
	"hotel/pkg/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB			//pointer to db, so that all functions across 
						//application modify the same db WITHOUT making a copy of it

func ConnectDB(){
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&models.User{}, &models.Listing{}, &models.Booking{})
	DB = database
}
