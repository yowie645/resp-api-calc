package db

import (
	"log"
	calcservice "resp-api-calc/internal/calcService"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=yowie645 dbname=postgres port=5342 sslmode=disable"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	if err := db.AutoMigrate(&calcservice.Calculation{}); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	return db, nil
}
