package configuration

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

/*
 - Using GORM as connector between application and database postgres
 - Todo hide connection details in a separate file
*/

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=postgres dbname=learningDb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
}