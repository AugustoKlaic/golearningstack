package service

import (
	"fmt"
	"github.com/AugustoKlaic/golearningstack/camundabpmn"
	. "github.com/AugustoKlaic/golearningstack/pkg/configuration"
	. "github.com/AugustoKlaic/golearningstack/pkg/domain/entity"
	. "github.com/AugustoKlaic/golearningstack/pkg/domain/error"
	. "github.com/AugustoKlaic/golearningstack/pkg/domain/repository"
	"github.com/AugustoKlaic/golearningstack/pkg/queue/rabbitmq"
	"github.com/AugustoKlaic/golearningstack/pkg/utils"
	"log"
	"os"
)

var serviceLogger = log.New(os.Stdout, "SERVICE: ", log.Ldate|log.Ltime|log.Lshortfile)

type LearningService struct {
	repo       LearningRepositoryInterface
	camundaAdm *camundabpmn.CamundaAdmin
}

func NewLearningService(repo LearningRepositoryInterface) *LearningService {
	return &LearningService{
		repo:       repo,
		camundaAdm: camundabpmn.NewCamundaAdmin(),
	}
}

func (s *LearningService) GetAllMessages() ([]MessageEntity, error) {
	serviceLogger.Println("Getting all messages...")
	if messages, err := s.repo.FindAllMessages(); err != nil {
		return nil, fmt.Errorf("problem retrieving messages: %v", err)
	} else {
		return messages, nil
	}
}

func (s *LearningService) CreateMessage(message *MessageEntity) (*MessageEntity, error) {
	serviceLogger.Println("Creating message...")
	if newMessage, err := s.repo.CreateMessage(message); err != nil {
		return nil, fmt.Errorf("problem creating message: %v", err)
	} else {
		return newMessage, nil
	}
}

func (s *LearningService) GetMessage(id int) (*MessageEntity, error) {
	serviceLogger.Printf("Getting message with Id: %d", id)
	if message, err := s.repo.GetMessage(id); err != nil {
		return nil, &MessageNotFoundError{Id: id}
	} else {
		publishToRabbit(message)
		s.camundaAdm.ExecuteProcess(message)
		return message, nil
	}
}

func (s *LearningService) DeleteMessage(id int) error {
	serviceLogger.Printf("Deleting message with Id: %d", id)
	if err := s.repo.DeleteMessage(id); err != nil {
		return fmt.Errorf("problem deleting message with id: %d. Error: %v", id, err)
	}
	return nil
}

func (s *LearningService) UpdateMessage(newMessage *MessageEntity, id int) (*MessageEntity, error) {
	serviceLogger.Printf("Updating message with Id: %d", id)
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

func publishToRabbit(message *MessageEntity) {
	encodedJson := utils.JsonEncoder(message)
	rabbitmq.PublishMessage(ExchangeName, RoutingKey, encodedJson, GetConnection(GetRabbitMQURL()))
}
