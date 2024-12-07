package service

import (
	"encoding/json"
	"fmt"
	. "github.com/AugustoKlaic/golearningstack/pkg/domain/entity"
	. "github.com/AugustoKlaic/golearningstack/pkg/domain/error"
	. "github.com/AugustoKlaic/golearningstack/pkg/domain/repository"
	"github.com/AugustoKlaic/golearningstack/pkg/queue/rabbitmq"
	"github.com/rabbitmq/amqp091-go"
	"log"
)

type LearningService struct {
	repo       LearningRepositoryInterface
	rabbitConn *amqp091.Connection
}

func NewLearningService(repo LearningRepositoryInterface, rabbitConn *amqp091.Connection) *LearningService {
	return &LearningService{
		repo:       repo,
		rabbitConn: rabbitConn,
	}
}

func (s *LearningService) GetAllMessages() ([]MessageEntity, error) {
	if messages, err := s.repo.FindAllMessages(); err != nil {
		return nil, fmt.Errorf("problem retrieving messages: %v", err)
	} else {
		return messages, nil
	}
}

func (s *LearningService) CreateMessage(message *MessageEntity) (*MessageEntity, error) {
	if newMessage, err := s.repo.CreateMessage(message); err != nil {
		return nil, fmt.Errorf("problem creating message: %v", err)
	} else {
		return newMessage, nil
	}
}

func (s *LearningService) GetMessage(id int) (*MessageEntity, error) {
	if message, err := s.repo.GetMessage(id); err != nil {
		return nil, &MessageNotFoundError{Id: id}
	} else {
		publishToRabbit(message, s.rabbitConn)
		return message, nil
	}
}

func (s *LearningService) DeleteMessage(id int) error {
	if err := s.repo.DeleteMessage(id); err != nil {
		return fmt.Errorf("problem deleting message with id: %d. Error: %v", id, err)
	}
	return nil
}

func (s *LearningService) UpdateMessage(newMessage *MessageEntity, id int) (*MessageEntity, error) {
	if oldMessage, err := s.GetMessage(id); err != nil {
		return nil, err
	} else {
		oldMessage.Content = newMessage.Content
		oldMessage.DateTime = newMessage.DateTime
		if updatedMessage, err := s.repo.UpdateMessage(oldMessage); err != nil {
			return nil, fmt.Errorf("problem updating message with id: %d. Error: %v", id, err)
		} else {
			return updatedMessage, nil
		}
	}
}

func publishToRabbit(message *MessageEntity, rabbitConn *amqp091.Connection) {
	if encodedJson, err := json.Marshal(message); err != nil {
		log.Printf("Erro ao converter mensagem: %v", err)
	} else {
		_ = rabbitmq.PublishMessage("", "", encodedJson, rabbitConn)
	}
}
