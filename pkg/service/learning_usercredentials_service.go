package service

import "github.com/AugustoKlaic/golearningstack/pkg/domain/repository"

type UserCredentialsService struct {
	repository repository.UserCredentialsRepositoryInterface
}

func NewUserCredentialsService(repo repository.UserCredentialsRepositoryInterface) *UserCredentialsService {
	return &UserCredentialsService{
		repository: repo,
	}
}
