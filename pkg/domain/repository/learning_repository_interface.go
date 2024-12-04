package repository

import "github.com/AugustoKlaic/golearningstack/pkg/domain/entity"

type LearningRepositoryInterface interface {
	CreateMessage(message *entity.MessageEntity) (*entity.MessageEntity, error)
	UpdateMessage(message *entity.MessageEntity) (*entity.MessageEntity, error)
	FindAllMessages() ([]entity.MessageEntity, error)
	GetMessage(id int) (*entity.MessageEntity, error)
	DeleteMessage(id int) error
}

// this directive is for creating a mock of this interface
//go:generate mockgen -source=learning_repository_interface.go -destination=../../../test/mock/repository/learning_repository_mock.go
