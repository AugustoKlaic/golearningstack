package controller

import (
	"fmt"
	_ "github.com/AugustoKlaic/golearningstack/docs" // importa os arquivos gerados
	"github.com/AugustoKlaic/golearningstack/pkg/api/errorvalidation"
	"github.com/AugustoKlaic/golearningstack/pkg/api/security/request"
	. "github.com/AugustoKlaic/golearningstack/pkg/domain/error"
	"github.com/AugustoKlaic/golearningstack/pkg/mapper"
	. "github.com/AugustoKlaic/golearningstack/pkg/service"
	"github.com/AugustoKlaic/golearningstack/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

var securityControllerLogger = log.New(os.Stdout, "SECURITY_CONTROLLER: ", log.Ldate|log.Ltime|log.Lshortfile)

type LearningSecurityController struct {
	service UserCredentialsServiceInterface
}

func NewLearningSecurityController(service UserCredentialsServiceInterface) *LearningSecurityController {
	return &LearningSecurityController{
		service: service,
	}
}

// @Summary Login for token generation
// @Description Generates JWT token based on user credentials
// @Tags Login
// @Accept json
// @Produce json
// @Param userCredentials	body	request.LoginRequest	true	"Create jwt Token for user"
// @Success 200 {string} token
// @Failure 403 {string} message
// @Failure 500 {string} message
// @Router /security/login [post]
func (ctrl *LearningSecurityController) Login(c *gin.Context) {
	//This is just a simulation for login for now
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

func (ctrl *LearningSecurityController) CreateUser(c *gin.Context) {
	securityControllerLogger.Printf("Creating new user...")
	var userCredentials request.LoginRequest

	if err := c.ShouldBindJSON(&userCredentials); err != nil {
		errorvalidation.HandleError(c, &InvalidCredentialsError{})
		return
	}

	if newUserId, err := ctrl.service.CreateUser(mapper.ToUserCredentialsEntity(&userCredentials)); err != nil {
		errorvalidation.HandleError(c, err)
	} else {
		c.IndentedJSON(http.StatusCreated, gin.H{"userId": newUserId})
	}
}
