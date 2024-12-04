package service

import "github.com/AugustoKlaic/golearningstack/pkg/domain/entity"

type LearningServiceInterface interface {
	GetAllMessages() ([]entity.MessageEntity, error)
	CreateMessage(message *entity.MessageEntity) (*entity.MessageEntity, error)
	GetMessage(id int) (*entity.MessageEntity, error)
	DeleteMessage(id int) error
	UpdateMessage(newMessage *entity.MessageEntity, id int) (*entity.MessageEntity, error)
}

// this directive is for creating a mock of this interface
//go:generate mockgen -source=learning_service_interface.go -destination=../../test/mock/service/learning_service_mock.go
