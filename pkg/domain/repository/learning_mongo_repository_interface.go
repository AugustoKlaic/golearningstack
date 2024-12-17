package repository

import (
	"github.com/AugustoKlaic/golearningstack/pkg/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserCredentialsRepositoryInterface interface {
	Create(entity entity.UserCredentials) (*mongo.InsertOneResult, error)
	FindByID(id string) (*entity.UserCredentials, error)
	Update(id string, entity entity.UserCredentials) (*mongo.UpdateResult, error)
	Delete(id string) (*mongo.DeleteResult, error)
}

// this directive is for creating a mock of this interface
//go:generate mockgen -source=learning_mongo_repository_interface.go -destination=../../../test/mock/repository/learning_mongo_repository_mock.go
