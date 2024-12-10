package main

import (
	"github.com/AugustoKlaic/golearningstack/pkg/api/controller"
	. "github.com/AugustoKlaic/golearningstack/pkg/api/router"
	. "github.com/AugustoKlaic/golearningstack/pkg/configuration"
	"github.com/AugustoKlaic/golearningstack/pkg/domain/repository"
	"github.com/AugustoKlaic/golearningstack/pkg/queue"
	"github.com/AugustoKlaic/golearningstack/pkg/service"
	"log"
	"os"
)

/*
 - Next-steps:
	- Add custom errors - ok
	- Add global error handler - ok
	- Unit test - ok
	- Add queue (rabbit - ok and kafka - in progress)
	- Add logging
	- Do a more complex entity to test GORM framework
	- Export properties (connections, passwords...) to a separate file with placeHolders
	- Secure API with jwtToken
	- MongoDb?
	- Sonar? It exists for golang?
*/

var mainLogger = log.New(os.Stdout, "MAIN: ", log.Ldate|log.Ltime|log.Lshortfile)

func main() {
	mainLogger.Println("Starting golearningstack project!")

	rabbitConn := GetConnection(GetRabbitMQURL())
	ConfigureRabbitMQ(rabbitConn)
	defer CloseConnection()
	mainLogger.Println("Rabbit configuration and connection done")

	messageRepo := repository.NewLearningRepository(ConnectDatabase())
	messageService := service.NewLearningService(messageRepo, rabbitConn)
	messageController := controller.NewLearningController(messageService)
	messageApiConsumer := queue.NewMessageApiConsumer(messageService)
	messageApiConsumer.Consume()

	if err := SetupRouter(messageController).Run("localhost:8080"); err != nil {
		mainLogger.Fatalf("Error starting API. Error: %v", err)
	} else {
		mainLogger.Println("Project started on port 8080!")
	}
}
