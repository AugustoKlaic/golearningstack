package configuration

import (
	"github.com/AugustoKlaic/golearningstack/pkg/domain/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

/*
 - Using GORM as connector between application and database postgres
 - AutoMigrate creates the tables automatically
 - Underline is for ignoring the return of function
 - Todo hide connection details in a separate file
 - Todo create a separate func to call automigrate
*/

func ConnectDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=learningDb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	} else {
		log.Println("Database connected!")
	}

	_ = db.AutoMigrate(&entity.MessageEntity{})

	return db
}
