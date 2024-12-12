package security

import (
	"github.com/AugustoKlaic/golearningstack/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type MiddlewareTokenValidation struct {
}

func NewMiddlewareTokenValidation() *MiddlewareTokenValidation {
	return &MiddlewareTokenValidation{}
}

func (middleware *MiddlewareTokenValidation) JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "token not provided or invalid"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := utils.ValidateToken(tokenString)

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}

		if claims, err := utils.GetClaims(token); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		} else {
			c.Set("username", claims)
		}
		c.Next()
	}
}
