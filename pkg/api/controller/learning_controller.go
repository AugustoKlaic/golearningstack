package controller

import (
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

	var message, err = ctrl.service.GetMessage(id)
	if err == nil {
		c.IndentedJSON(http.StatusOK, mapper.ToMessageResponse(message))
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Message not found"})
	}
}

func (ctrl *LearningController) CreateMessage(c *gin.Context) {
	var newMessage request.MessageRequest

	if err := c.BindJSON(&newMessage); err != nil {
		return
	}

	_ = ctrl.service.CreateMessage(mapper.ToMessageEntity(newMessage))

	c.IndentedJSON(http.StatusCreated, newMessage)
}

func (ctrl *LearningController) DeleteMessage(c *gin.Context) {
	var id, _ = strconv.Atoi(c.Param("id"))

	err := ctrl.service.DeleteMessage(id)
	if err == nil {
		c.Status(http.StatusNoContent)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Message not found"})
	}
}

func (ctrl *LearningController) UpdateMessage(c *gin.Context) {
	var id, _ = strconv.Atoi(c.Param("id"))
	var updateMessage request.MessageRequest

	if err := c.BindJSON(&updateMessage); err != nil {
		return
	}

	newMessage, err := ctrl.service.UpdateMessage(mapper.ToMessageEntity(updateMessage), id)

	if err == nil {
		c.IndentedJSON(http.StatusOK, mapper.ToMessageResponse(newMessage))
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Message not found"})
	}
}
