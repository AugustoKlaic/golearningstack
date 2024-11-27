package domain

import "gorm.io/gorm"

type LearningRepository struct {
	db *gorm.DB
}

func NewLearningRepository(db *gorm.DB) *LearningRepository {
	return &LearningRepository{db: db}
}

func (r *LearningRepository) CreateMessage(message *MessageEntity) error {
	return r.db.Create(message).Error
}

func (r *LearningRepository) FindAllMessages() ([]MessageEntity, error) {
	var messages []MessageEntity
	err := r.db.Find(&messages).Error
	return messages, err
}
