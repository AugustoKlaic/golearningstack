package controller

import (
	. "github.com/AugustoKlaic/golearningstack/pkg/api/errorvalidation"
	"github.com/AugustoKlaic/golearningstack/pkg/api/request"
	"github.com/AugustoKlaic/golearningstack/pkg/mapper"
	"github.com/AugustoKlaic/golearningstack/pkg/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
)

var controllerLogger = log.New(os.Stdout, "CONTROLLER: ", log.Ldate|log.Ltime|log.Lshortfile)

type LearningController struct {
	service service.LearningServiceInterface
}

func NewLearningController(service service.LearningServiceInterface) *LearningController {
	return &LearningController{service: service}
}

func (ctrl *LearningController) GetAllMessages(c *gin.Context) {
	controllerLogger.Println("Getting all messages...")
	if messages, err := ctrl.service.GetAllMessages(); err != nil {
		HandleError(c, err)
		return
	} else {
		c.IndentedJSON(http.StatusOK, mapper.ToMessageResponses(messages...))
	}
}

func (ctrl *LearningController) GetMessage(c *gin.Context) {
	var id, _ = strconv.Atoi(c.Param("id"))
	controllerLogger.Printf("Getting message with Id: %v", id)

	if message, err := ctrl.service.GetMessage(id); err != nil {
		HandleError(c, err)
		return
	} else {
		c.IndentedJSON(http.StatusOK, mapper.ToMessageResponse(message))
	}
}

func (ctrl *LearningController) CreateMessage(c *gin.Context) {
	controllerLogger.Println("Creating new message ...")
	var message request.MessageRequest

	if err := c.BindJSON(&message); err != nil {
		HandleError(c, err)
		return
	}

	if newMessage, err := ctrl.service.CreateMessage(mapper.ToMessageEntity(message)); err != nil {
		HandleError(c, err)
		return
	} else {
		c.IndentedJSON(http.StatusCreated, mapper.ToMessageResponse(newMessage))
	}
}

func (ctrl *LearningController) DeleteMessage(c *gin.Context) {
	var id, _ = strconv.Atoi(c.Param("id"))
	controllerLogger.Printf("Deleting message with Id: %v", id)

	if err := ctrl.service.DeleteMessage(id); err != nil {
		HandleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (ctrl *LearningController) UpdateMessage(c *gin.Context) {
	var id, _ = strconv.Atoi(c.Param("id"))
	var updateMessage request.MessageRequest
	controllerLogger.Printf("Updating message with Id: %v", id)

	if err := c.BindJSON(&updateMessage); err != nil {
		HandleError(c, err)
		return
	}

	if newMessage, err := ctrl.service.UpdateMessage(mapper.ToMessageEntity(updateMessage), id); err != nil {
		HandleError(c, err)
		return
	} else {
		c.IndentedJSON(http.StatusOK, mapper.ToMessageResponse(newMessage))
	}
}
