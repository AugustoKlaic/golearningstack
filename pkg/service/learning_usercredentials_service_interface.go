package service

import "github.com/AugustoKlaic/golearningstack/pkg/domain/entity"

type UserCredentialsServiceInterface interface {
	CreateUser(newUser *entity.UserCredentials) (interface{}, error)
}
