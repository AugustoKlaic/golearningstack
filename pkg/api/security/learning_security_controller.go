package security

import (
	"fmt"
	"github.com/AugustoKlaic/golearningstack/pkg/api/errorvalidation"
	"github.com/AugustoKlaic/golearningstack/pkg/api/security/request"
	. "github.com/AugustoKlaic/golearningstack/pkg/domain/error"
	"github.com/AugustoKlaic/golearningstack/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

var securityControllerLogger = log.New(os.Stdout, "SECURITY_CONTROLLER: ", log.Ldate|log.Ltime|log.Lshortfile)

type LearningSecurityController struct {
}

func NewLearningSecurityController() *LearningSecurityController {
	return &LearningSecurityController{}
}

//This is just a simulation for login

func (ctrl *LearningSecurityController) Login(c *gin.Context) {
	securityControllerLogger.Printf("loggin in...")
	var userCredentials request.LoginRequest

	if err := c.ShouldBindJSON(&userCredentials); err != nil {
		errorvalidation.HandleError(c, &InvalidCredentialsError{})
		return
	}

	token, err := utils.GenerateToken(userCredentials.UserName)

	if err != nil {
		errorvalidation.HandleError(c, fmt.Errorf("erro ao gerar token"))
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"token": token})
}
