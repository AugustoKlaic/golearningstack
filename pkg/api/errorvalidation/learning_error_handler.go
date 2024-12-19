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
	errMessageNotFound    = &MessageNotFoundError{}
	errInvalidCredentials = &InvalidCredentialsError{}
	validationErrors      = &validator.ValidationErrors{}
	errUserNotFound       = &UserNotFoundError{}
)

func HandleError(c *gin.Context, err error) {

	switch {
	case errors.As(err, &errMessageNotFound) || errors.As(err, &errUserNotFound):
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Not found. Error: %v", err.Error())})
	case errors.As(err, &errInvalidCredentials):
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
	case errors.As(err, validationErrors):
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("invalid request JSON. Error: %v", err.Error())})
	default:
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
}
