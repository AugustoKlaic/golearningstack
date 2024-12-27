package controller

import (
	_ "github.com/AugustoKlaic/golearningstack/docs" // importa os arquivos gerados
	"github.com/AugustoKlaic/golearningstack/pkg/api/errorvalidation"
	"github.com/AugustoKlaic/golearningstack/pkg/api/security/request"
	"github.com/AugustoKlaic/golearningstack/pkg/mapper"
	. "github.com/AugustoKlaic/golearningstack/pkg/service"
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
// @Tags Security
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
		errorvalidation.HandleError(c, err)
		return
	}

	if token, err := ctrl.service.GenerateUserToken(mapper.ToUserCredentialsEntity(&userCredentials)); err != nil {
		errorvalidation.HandleError(c, err)
		return
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"token": token})
	}
}

// @Summary Create a new user
// @Description Create a new user for the system
// @Tags Security
// @Accept json
// @Produce json
// @Param userCredentials	body	request.LoginRequest	true	"Create a new user"
// @Success 201 {interface} mongo.InsertOneResult
// @Failure 404 {string} message
// @Failure 400 {string} message
// @Failure 500 {string} message
// @Router /security/add-user [post]
func (ctrl *LearningSecurityController) CreateUser(c *gin.Context) {
	securityControllerLogger.Printf("Creating new user...")
	var userCredentials request.LoginRequest

	if err := c.ShouldBindJSON(&userCredentials); err != nil {
		errorvalidation.HandleError(c, err)
		return
	}

	if newUserId, err := ctrl.service.CreateUser(mapper.ToUserCredentialsEntity(&userCredentials)); err != nil {
		errorvalidation.HandleError(c, err)
		return
	} else {
		c.IndentedJSON(http.StatusCreated, gin.H{"userId": newUserId})
	}
}
