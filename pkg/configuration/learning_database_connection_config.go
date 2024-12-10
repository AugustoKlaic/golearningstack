package configuration

import (
	"fmt"
	"github.com/AugustoKlaic/golearningstack/pkg/domain/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

/*
 - Using GORM as connector between application and database postgres
 - AutoMigrate creates the tables automatically
 - Underline is for ignoring the return of function
 - Todo create a separate func to call automigrate
*/

var databaseConfigLogger = log.New(os.Stdout, "CONFIGURATION: ", log.Ldate|log.Ltime|log.Lshortfile)

func ConnectDatabase() *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		Props.DB.Host, Props.DB.User, Props.DB.Password, Props.DB.Dbname, Props.DB.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		databaseConfigLogger.Fatalf("Failed to connect to database: %v", err)
	} else {
		databaseConfigLogger.Println("Database connected!")
	}

	_ = db.AutoMigrate(&entity.MessageEntity{})

	return db
}
