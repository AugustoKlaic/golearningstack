package controller

import (
	"github.com/AugustoKlaic/golearningstack/pkg/api/response"
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
	if err != nil {
		c.IndentedJSON(http.StatusOK, mapper.ToMessageResponse(message))
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Message not found"})
	}
}

func (ctrl *LearningController) CreateMessage(c *gin.Context) {
	var newMessage response.Message

	if err := c.BindJSON(&newMessage); err != nil {
		return
	}

	_ = ctrl.service.CreateMessage(mapper.ToMessageEntity(newMessage))

	c.IndentedJSON(http.StatusCreated, newMessage)
}

func (ctrl *LearningController) DeleteMessage(c *gin.Context) {
	//var id, _ = strconv.Atoi(c.Param("id"))
	//
	//for i, a := range Messages {
	//	if a.Id == id {
	//		Messages = append(Messages[:i], Messages[i+1:]...)
	//		c.Status(http.StatusNoContent)
	//		return
	//	}
	//}
	//
	//c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Message not found"})
}

func (ctrl *LearningController) UpdateMessage(c *gin.Context) {
	//var id, _ = strconv.Atoi(c.Param("id"))
	//
	//for i, a := range Messages {
	//	if a.Id == id {
	//		var updateMessage Message
	//
	//		if err := c.BindJSON(&updateMessage); err != nil {
	//			return
	//		}
	//
	//		Messages[i].Content = updateMessage.Content
	//		Messages[i].DateTime = updateMessage.DateTime
	//
	//		c.IndentedJSON(http.StatusOK, Messages[i])
	//		return
	//	}
	//}
	//c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Message not found"})
}
