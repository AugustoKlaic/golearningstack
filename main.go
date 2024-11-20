package main

import (
	"fmt"
	. "github.com/AugustoKlaic/golearningstack/pkg/api"
	"github.com/gin-gonic/gin"
)

/*
 - For the web api I will be utilizing the GIN web framework
 - I am starting router in default mode
 - I am registering all routes with its respective functions
 - I am running the server at localhost:8080 and checking for error at startup
*/

func main() {
	fmt.Println("Iniciando o projeto golearningstack!")

	var router = gin.Default()

	StartLearningApiRouter(router)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
