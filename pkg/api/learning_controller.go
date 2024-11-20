package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
 - Here will be created all routes for the specific domain "learning"
 - The only function exported here will be the one with the router
*/

func getMessages(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Messages)
}

func StartLearningApiRouter(router *gin.Engine) {
	router.GET("/learning", getMessages)
}
