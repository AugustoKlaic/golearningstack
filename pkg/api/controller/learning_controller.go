package controller

import (
	"fmt"
	"github.com/AugustoKlaic/golearningstack/pkg/api/request"
	"github.com/AugustoKlaic/golearningstack/pkg/mapper"
	"github.com/AugustoKlaic/golearningstack/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/*
 - Here will be created all routes for the specific domain "learning"
 - The only function exported here will be the one with the router
 - The for loop iterates over copies of the original value <- pay attention to that
*/

type LearningController struct {
	service *service.LearningService
}

func NewLearningController(service *service.LearningService) *LearningController {
	return &LearningController{service: service}
}

func (ctrl *LearningController) GetAllMessages(c *gin.Context) {
	var messages, _ = ctrl.service.GetAllMessages()
	c.IndentedJSON(http.StatusOK, mapper.ToMessageResponses(messages...))
}

func (ctrl *LearningController) GetMessage(c *gin.Context) {
	var id, _ = strconv.Atoi(c.Param("id"))

	if message, err := ctrl.service.GetMessage(id); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Message not found. Error: %v", err.Error())})
	} else {
		c.IndentedJSON(http.StatusOK, mapper.ToMessageResponse(message))
	}
}

func (ctrl *LearningController) CreateMessage(c *gin.Context) {
	var message request.MessageRequest

	if err := c.BindJSON(&message); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("could not bind request body. Error: %v", err.Error())})
	}

	if newMessage, err := ctrl.service.CreateMessage(mapper.ToMessageEntity(message)); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error occured: %v", err.Error())})
	} else {
		c.IndentedJSON(http.StatusCreated, newMessage)
	}
}

func (ctrl *LearningController) DeleteMessage(c *gin.Context) {
	var id, _ = strconv.Atoi(c.Param("id"))

	if err := ctrl.service.DeleteMessage(id); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Message not found. Error: %v", err.Error())})
	}

	c.Status(http.StatusNoContent)
}

func (ctrl *LearningController) UpdateMessage(c *gin.Context) {
	var id, _ = strconv.Atoi(c.Param("id"))
	var updateMessage request.MessageRequest

	if err := c.BindJSON(&updateMessage); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("could not bind request body. Error: %v", err.Error())})
	}

	if newMessage, err := ctrl.service.UpdateMessage(mapper.ToMessageEntity(updateMessage), id); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Message not found. Error: %v", err.Error())})
	} else {
		c.IndentedJSON(http.StatusOK, mapper.ToMessageResponse(newMessage))
	}
}
