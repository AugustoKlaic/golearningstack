package router

import (
	"github.com/AugustoKlaic/golearningstack/pkg/api/message/controller"
	"github.com/AugustoKlaic/golearningstack/pkg/api/security"
	controller2 "github.com/AugustoKlaic/golearningstack/pkg/api/security/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter(messageController *controller.LearningController,
	securityController *controller2.LearningSecurityController,
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
