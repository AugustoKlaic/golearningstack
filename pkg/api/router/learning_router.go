package router

import (
	"github.com/AugustoKlaic/golearningstack/pkg/api/controller"
	"github.com/AugustoKlaic/golearningstack/pkg/configuration"
	"github.com/AugustoKlaic/golearningstack/pkg/domain"
	"github.com/AugustoKlaic/golearningstack/pkg/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	messageRepo := domain.NewLearningRepository(configuration.DB)
	messageService := service.NewLearningService(messageRepo)
	messageController := controller.NewLearningController(messageService)

	messageApi := router.Group("/learning")
	{
		messageApi.GET("", messageController.GetAllMessages)
		messageApi.GET("/:id", messageController.GetMessage)
		messageApi.POST("", messageController.CreateMessage)
		messageApi.DELETE(":id", messageController.DeleteMessage)
		messageApi.PATCH("/:id", messageController.UpdateMessage)
	}

	return router
}
