package configuration

import (
	"context"
	"fmt"
	"github.com/AugustoKlaic/golearningstack/pkg/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

/*
 - Using GORM as connector between application and database postgres
 - AutoMigrate creates the tables automatically
 - Underline is for ignoring the return of function
 - [ctx, cancel := ...] line, set a timeout for when a problem occurs with the mongo connection and invalidate it
 - in golang to create "constraints" it has to be inside the connector, there is no annotations like java
 - EnsureUniqueIndex is defined in each entity file to be organized
 - Todo create a separate func to call automigrate
*/

var databaseConfigLogger = log.New(os.Stdout, "CONFIGURATION: ", log.Ldate|log.Ltime|log.Lshortfile)

func ConnectPostgresDatabase() *gorm.DB {

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

func ConnectMongoDatabase() *mongo.Database {
	uri := fmt.Sprintf("mongodb://%s:%s/%s", Props.Mongo.Host, Props.Mongo.Port, Props.Mongo.Dbname)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		databaseConfigLogger.Fatalf("Failed to connect to MongoDB: %v", err)
	} else {
		databaseConfigLogger.Println("MongoDB connected!")
	}

	var database = client.Database(Props.Mongo.Dbname)

	_ = entity.EnsureUniqueIndex(database.Collection(entity.CollectionName))

	return database
}
