package errorvalidation

import (
	"errors"
	"fmt"
	. "github.com/AugustoKlaic/golearningstack/pkg/domain/error"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var (
	errMessageNotFound = &MessageNotFoundError{}
	validationErrors   = &validator.ValidationErrors{}
)

func HandleError(c *gin.Context, err error) {

	switch {
	case errors.As(err, &errMessageNotFound):
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Message not found. Error: %v", err.Error())})
	case errors.As(err, validationErrors):
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("invalid request JSON. Error: %v", err.Error())})
	default:
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
}
