package utils

import (
	"github.com/AugustoKlaic/golearningstack/pkg/configuration"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateToken(userName string) (string, error) {
	claims := jwt.MapClaims{
		"username": userName,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(configuration.Props.Jwt.Secret)
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, verifyTokenSignature)
}

func verifyTokenSignature(token *jwt.Token) (interface{}, error) {
	if _, isHmacSigned := token.Method.(*jwt.SigningMethodHMAC); !isHmacSigned {
		return nil, jwt.ErrSignatureInvalid
	}
	return configuration.Props.Jwt.Secret, nil
}

func GetClaims(token *jwt.Token) (interface{}, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims["username"], nil
	} else {
		return nil, jwt.ErrTokenInvalidClaims
	}
}
