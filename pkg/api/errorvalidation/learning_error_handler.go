package errorvalidation

import (
	"errors"
	"fmt"
	. "github.com/AugustoKlaic/golearningstack/pkg/domain/error"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	ErrMessageNotFound = &MessageNotFoundError{}
)

func HandleError(c *gin.Context, err error) {

	switch {
	case errors.As(err, &ErrMessageNotFound):
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Message not found. Error: %v", err.Error())})
	default:
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
}
