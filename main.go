package main

import (
	"fmt"
	. "github.com/AugustoKlaic/golearningstack/pkg/api/router"
	. "github.com/AugustoKlaic/golearningstack/pkg/configuration"
)

/*
 - For the web api I will be utilizing the GIN web framework
 - I am starting router in default mode
 - I am registering all routes with its respective functions
 - I am running the server at localhost:8080 and checking for error at startup
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
