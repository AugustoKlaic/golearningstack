package service

import (
	"github.com/AugustoKlaic/golearningstack/pkg/domain/entity"
	. "github.com/AugustoKlaic/golearningstack/pkg/domain/error"
	"github.com/AugustoKlaic/golearningstack/pkg/domain/repository"
	"github.com/AugustoKlaic/golearningstack/pkg/utils"
)

type UserCredentialsService struct {
	repository repository.UserCredentialsRepositoryInterface
}

func NewUserCredentialsService(repo repository.UserCredentialsRepositoryInterface) *UserCredentialsService {
	return &UserCredentialsService{
		repository: repo,
	}
}

func (r *UserCredentialsService) CreateUser(newUser *entity.UserCredentials) (interface{}, error) {
	var hashPassword, err = utils.HashPassword(newUser.Password)

	if err != nil {
		return nil, &UnhashablePasswordError{}
	} else {
		newUser.Password = hashPassword
	}

	if createdUser, err := r.repository.Create(newUser); err != nil {
		return nil, err
	} else {
		return createdUser.InsertedID, nil
	}
}

func (r *UserCredentialsService) GenerateUserToken(userCredentials *entity.UserCredentials) (string, error) {

	if foundUser, err := r.repository.FindByUserName(userCredentials.Username); err != nil {
		return "", &UserNotFoundError{}
	} else {
		if utils.CheckPassword(foundUser.Password, userCredentials.Password) {
			if token, err := utils.GenerateToken(userCredentials.Username); err != nil {
				return token, err
			} else {
				return token, nil
			}
		} else {
			return "", &InvalidCredentialsError{}
		}
	}
}
