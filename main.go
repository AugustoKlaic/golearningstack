package main

import (
	"fmt"
	_ "github.com/AugustoKlaic/golearningstack/docs"
	. "github.com/AugustoKlaic/golearningstack/pkg/api/message/controller"
	. "github.com/AugustoKlaic/golearningstack/pkg/api/router"
	. "github.com/AugustoKlaic/golearningstack/pkg/api/security"
	. "github.com/AugustoKlaic/golearningstack/pkg/api/security/controller"
	. "github.com/AugustoKlaic/golearningstack/pkg/configuration"
	. "github.com/AugustoKlaic/golearningstack/pkg/domain/repository"
	. "github.com/AugustoKlaic/golearningstack/pkg/queue"
	. "github.com/AugustoKlaic/golearningstack/pkg/service"
	"log"
	"os"
)

/*
 - Next-steps:
	- Add custom errors - ok
	- Add global error handler - ok
	- Unit test - ok
	- Add queue (rabbit - ok and kafka - ok)
		- kafka on windows -> https://github.com/confluentinc/confluent-kafka-go/issues/889
	- Add logging - ok
	- Export properties to a separate file with placeHolders - ok
	- Sonar - ok
	- Secure API with jwtToken - ok
	- swagger for golang (http://localhost:8080/swagger/index.html#/)- ok
	- MongoDb to credentials storage - ok
	- Camunda BPMN - ok
*/

var mainLogger = log.New(os.Stdout, "MAIN: ", log.Ldate|log.Ltime|log.Lshortfile)

// @title Message API swagger
// @version 1.0
// @description This is an API that manipulate messages
// @host localhost:8080
func main() {
	mainLogger.Println("Loading env variables")
	LoadConfig("application.yaml")

	ConfigureRabbitMQ()
	defer CloseRabbitMqResources()
	mainLogger.Println("Rabbit configuration and connection done")

	ConfigureKafka()
	defer CloseKafkaResources()
	mainLogger.Println("Kafka configuration and connection done")

	CreateClient()
	mainLogger.Println("Camunda client started")

	messageController, securityController, middleware := initializeDependencies()

	if err := SetupRouter(messageController, securityController, middleware).
		Run(fmt.Sprintf("%s:%s", Props.Gin.Host, Props.Gin.Port)); err != nil {
		mainLogger.Fatalf("Error starting API. Error: %v", err)
	} else {
		mainLogger.Println("Project started on port 8080!")
	}
}

func initializeDependencies() (*LearningController, *LearningSecurityController, *MiddlewareTokenValidation) {
	userCredentialsRepo := NewUserCredentialsRepository(ConnectMongoDatabase())
	userCredentialsService := NewUserCredentialsService(userCredentialsRepo)

	messageRepo := NewLearningRepository(ConnectPostgresDatabase())
	messageService := NewLearningService(messageRepo)

	messageApiConsumer := NewMessageApiConsumer(messageService)
	messageApiConsumer.Consume()
	userApiConsumer := NewUserApiConsumer(userCredentialsService)
	userApiConsumer.Consume()

	middleware := NewMiddlewareTokenValidation()
	messageController := NewLearningController(messageService)
	securityController := NewLearningSecurityController(userCredentialsService)

	return messageController, securityController, middleware
}
