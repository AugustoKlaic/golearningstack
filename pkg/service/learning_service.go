package service

import (
	"github.com/AugustoKlaic/golearningstack/pkg/api/response"
	. "github.com/AugustoKlaic/golearningstack/pkg/domain"
)

type LearningService struct {
	repo *LearningRepository
}

func NewLearningService(repo *LearningRepository) *LearningService {
	return &LearningService{repo: repo}
}

func (s *LearningService) GetAllMessages() ([]MessageEntity, error) {
	return s.repo.FindAllMessages()
}

func (s *LearningService) CreateMessage(user *MessageEntity) error {
	return s.repo.CreateMessage(user)
}

func (s *LearningService) GetMessage(id int) (*MessageEntity, error) {
	return s.repo.GetMessage(id)
}

func (s *LearningService) DeleteMessage(id int) error {
	return s.repo.DeleteMessage(id)
}

func (s *LearningService) UpdateMessage(newMessage *response.Message, id int) (*MessageEntity, error) {
	var oldMessage, err = s.GetMessage(id)

	if err == nil {
		oldMessage.Content = newMessage.Content
		oldMessage.DateTime = newMessage.DateTime
		_ = s.repo.UpdateMessage(oldMessage)
		return oldMessage, nil
	} else {
		return nil, err
	}
}
