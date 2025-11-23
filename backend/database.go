package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=1234 dbname=shoppingdb port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	// Auto-migrate models
	if err := database.AutoMigrate(&User{}, &Item{}, &Cart{}, &Order{}); err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}

	DB = database
}
