package main

import (
	"fmt"
	"github.com/AugustoKlaic/golearningstack/pkg/api/controller"
	. "github.com/AugustoKlaic/golearningstack/pkg/api/router"
	"github.com/AugustoKlaic/golearningstack/pkg/api/security"
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
	- Add logging - ok
	- Export properties to a separate file with placeHolders - ok
	- Sonar - ok
	- Secure API with jwtToken - ok
	- swagger for golang
	- Adjust code to be more Object-Oriented
 	- Jenkins
	- MongoDb?
	- Do a more complex entity to test GORM framework
*/

var mainLogger = log.New(os.Stdout, "MAIN: ", log.Ldate|log.Ltime|log.Lshortfile)

func main() {
	mainLogger.Println("Starting golearningstack project!")
	LoadConfig("application.yaml")

	rabbitConn := GetConnection(GetRabbitMQURL())
	ConfigureRabbitMQ(rabbitConn)
	defer CloseConnection()
	mainLogger.Println("Rabbit configuration and connection done")

	messageRepo := repository.NewLearningRepository(ConnectDatabase())
	messageService := service.NewLearningService(messageRepo, rabbitConn)
	messageController := controller.NewLearningController(messageService)
	messageApiConsumer := queue.NewMessageApiConsumer(messageService)
	messageApiConsumer.Consume()
	middleware := security.NewMiddlewareTokenValidation()
	securityController := security.NewLearningSecurityController()

	if err := SetupRouter(messageController, securityController, middleware).
		Run(fmt.Sprintf("%s:%s", Props.Gin.Host, Props.Gin.Port)); err != nil {
		mainLogger.Fatalf("Error starting API. Error: %v", err)
	} else {
		mainLogger.Println("Project started on port 8080!")
	}
}
