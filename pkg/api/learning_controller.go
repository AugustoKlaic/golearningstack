package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/*
 - Here will be created all routes for the specific domain "learning"
 - The only function exported here will be the one with the router
*/

func getAllMessages(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Messages)
}

func getMessage(c *gin.Context) {
	var id, _ = strconv.Atoi(c.Param("id"))

	for _, a := range Messages {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Message not found"})
}

func StartLearningApiRouter(router *gin.Engine) {
	router.GET("/learning", getAllMessages)
	router.GET("/learning/:id", getMessage)
}
