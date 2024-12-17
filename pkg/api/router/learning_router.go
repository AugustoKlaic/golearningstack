package router

import (
	"github.com/AugustoKlaic/golearningstack/pkg/api/message/controller"
	"github.com/AugustoKlaic/golearningstack/pkg/api/security"
	. "github.com/AugustoKlaic/golearningstack/pkg/api/security/controller"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(messageController *controller.LearningController,
	securityController *LearningSecurityController,
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
