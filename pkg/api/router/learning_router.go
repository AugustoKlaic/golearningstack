package router

import (
	"github.com/AugustoKlaic/golearningstack/pkg/api/controller"
	"github.com/AugustoKlaic/golearningstack/pkg/api/security"
	"github.com/gin-gonic/gin"
)

func SetupRouter(messageController *controller.LearningController, middleware *security.MiddlewareTokenValidation) *gin.Engine {
	router := gin.Default()

	messageApi := router.Group("/learning", middleware.JwtAuthMiddleware())
	{
		messageApi.GET("", messageController.GetAllMessages)
		messageApi.GET("/:id", messageController.GetMessage)
		messageApi.POST("", messageController.CreateMessage)
		messageApi.DELETE(":id", messageController.DeleteMessage)
		messageApi.PUT("/:id", messageController.UpdateMessage)
	}

	return router
}
