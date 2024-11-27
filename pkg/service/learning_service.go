package service

import . "github.com/AugustoKlaic/golearningstack/pkg/domain"

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
