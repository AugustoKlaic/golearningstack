package main

import (
	"fmt"
	"github.com/AugustoKlaic/golearningstack/pkg/api/controller"
	. "github.com/AugustoKlaic/golearningstack/pkg/api/router"
	. "github.com/AugustoKlaic/golearningstack/pkg/configuration"
	"github.com/AugustoKlaic/golearningstack/pkg/domain/repository"
	"github.com/AugustoKlaic/golearningstack/pkg/service"
)

/*
 - Next-steps:
	- Add custom errors - ok
	- Add global error handler - ok
	- Unit test - ok
	- Add queue (rabbit and kafka)
	- Add logging
	- Do a more complex entity to test GORM framework
	- Export properties (connections, passwords...) to a separate file with placeHolders
	- Secure API with jwtToken
	- MongoDb?
	- Sonar? It exists for golang?
*/

func main() {
	fmt.Println("Iniciando o projeto golearningstack!")

	messageRepo := repository.NewLearningRepository(ConnectDatabase())
	messageService := service.NewLearningService(messageRepo)
	messageController := controller.NewLearningController(messageService)

	if err := SetupRouter(messageController).Run("localhost:8080"); err != nil {
		return
	}
}
