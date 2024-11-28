package domain

import "gorm.io/gorm"

/*
 - Operator Unscoped means that it will execute the designated operation on soft deleted rows
 -
*/

type LearningRepository struct {
	db *gorm.DB
}

func NewLearningRepository(db *gorm.DB) *LearningRepository {
	return &LearningRepository{db: db}
}

func (r *LearningRepository) CreateMessage(message *MessageEntity) error {
	return r.db.Create(&message).Error
}

func (r *LearningRepository) UpdateMessage(message *MessageEntity) error {
	return r.db.Save(&message).Error
}

func (r *LearningRepository) FindAllMessages() ([]MessageEntity, error) {
	var messages []MessageEntity
	err := r.db.Find(&messages).Error
	return messages, err
}

func (r *LearningRepository) GetMessage(id int) (*MessageEntity, error) {
	var message MessageEntity

	err := r.db.First(&message, id).Error

	return &message, err
}

func (r *LearningRepository) DeleteMessage(id int) error {
	return r.db.Unscoped().Delete(&MessageEntity{}, id).Error
}
