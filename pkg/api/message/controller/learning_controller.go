package controller

import (
	. "github.com/AugustoKlaic/golearningstack/pkg/api/errorvalidation"
	"github.com/AugustoKlaic/golearningstack/pkg/api/message/request"
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

// @Summary Get all messages
// @Description Fetch all messages saved
// @Tags Message
// @Produce json
// @Param	Authorization	header	string	true "Bearer token"
// @Success 200 {array} response.Message
// @Failure 403 {string} message
// @Failure 500 {string} message
// @Router /learning [get]
func (ctrl *LearningController) GetAllMessages(c *gin.Context) {
	controllerLogger.Println("Getting all messages...")
	if messages, err := ctrl.service.GetAllMessages(); err != nil {
		HandleError(c, err)
		return
	} else {
		c.IndentedJSON(http.StatusOK, mapper.ToMessageResponses(messages...))
	}
}

// @Summary Get message by Id
// @Description Fetch message by Id
// @Tags Message
// @Produce json
// @Param	Authorization	header	string	true "Bearer token"
// @Param			id		path		int 				true	"Message ID"
// @Success 200 {object} response.Message
// @Failure 403 {string} message
// @Failure 404 {string} message
// @Failure 500 {string} message
// @Router /learning/{id} [get]
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

// @Summary Create a new message
// @Description Create a new message
// @Tags Message
// @Accept json
// @Produce json
// @Param	Authorization	header	string	true "Bearer token"
// @Param			message	body		request.MessageRequest	true	"Create message"
// @Success 201 {object} response.Message
// @Failure 403 {string} message
// @Failure 400 {string} message
// @Failure 500 {string} message
// @Router /learning [post]
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

// @Summary Delete message by Id
// @Description Delete message by Id
// @Tags Message
// @Param	Authorization	header	string	true "Bearer token"
// @Param			id		path		int 				true	"Message ID"
// @Success 204
// @Failure 403 {string} message
// @Failure 500 {string} message
// @Router /learning/{id} [delete]
func (ctrl *LearningController) DeleteMessage(c *gin.Context) {
	var id, _ = strconv.Atoi(c.Param("id"))
	controllerLogger.Printf("Deleting message with Id: %v", id)

	if err := ctrl.service.DeleteMessage(id); err != nil {
		HandleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary Update a message
// @Description Update a message
// @Tags Message
// @Accept json
// @Produce json
// @Param	Authorization	header	string	true "Bearer token"
// @Param			message	body		request.MessageRequest	true	"Update message"
// @Param			id		path		int 				true	"Message ID"
// @Success 200 {object} response.Message
// @Failure 404 {string} message
// @Failure 403 {string} message
// @Failure 400 {string} message
// @Failure 500 {string} message
// @Router /learning/{id} [put]
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
