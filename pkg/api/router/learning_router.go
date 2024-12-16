package router

import (
	"github.com/AugustoKlaic/golearningstack/pkg/api/controller"
	"github.com/AugustoKlaic/golearningstack/pkg/api/security"
	"github.com/gin-gonic/gin"
)

func SetupRouter(messageController *controller.LearningController,
	securityController *security.LearningSecurityController,
	middleware *security.MiddlewareTokenValidation) *gin.Engine {

	router := gin.Default()

	messageApi := router.Group("/learning", middleware.JwtAuthMiddleware())
	{
		messageApi.GET("", messageController.GetAllMessages)
		messageApi.GET("/:id", messageController.GetMessage)
		messageApi.POST("", messageController.CreateMessage)
		messageApi.DELETE(":id", messageController.DeleteMessage)
		messageApi.PUT("/:id", messageController.UpdateMessage)
	}

	securityApi := router.Group("/security")
	{
		securityApi.POST("/login", securityController.Login)
	}

	return router
}
