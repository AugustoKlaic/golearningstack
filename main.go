package main

import (
	"fmt"
	. "github.com/AugustoKlaic/golearningstack/pkg/api/router"
	. "github.com/AugustoKlaic/golearningstack/pkg/configuration"
)

/*
 - Next-steps:
	- Add custom errors
	- Add global error handler
	- Add queue (rabbit or kafka)
	- Add logging
	- Do a more complex entity to test GORM framework
	- Export properties (connections, passwords...) to a separate file with placeHolders
	- Secure API with jwtToken
	- MongoDb?
*/

func main() {
	fmt.Println("Iniciando o projeto golearningstack!")

	ConnectDatabase()
	var router = SetupRouter()

	err := router.Run("localhost:8080")

	if err != nil {
		return
	}
}
