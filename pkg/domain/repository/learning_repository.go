package repository

import (
	"github.com/AugustoKlaic/golearningstack/pkg/domain/entity"
	"gorm.io/gorm"
)

/*
 - Operator Unscoped means that it will execute the designated operation on soft deleted rows
*/

type LearningRepository struct {
	db *gorm.DB
}

func NewLearningRepository(db *gorm.DB) *LearningRepository {
	return &LearningRepository{db: db}
}

func (r *LearningRepository) CreateMessage(message *entity.MessageEntity) (*entity.MessageEntity, error) {
	if err := r.db.Create(message).Error; err != nil {
		return nil, err
	}
	return message, nil
}

func (r *LearningRepository) UpdateMessage(message *entity.MessageEntity) (*entity.MessageEntity, error) {
	if err := r.db.Save(&message).Error; err != nil {
		return nil, err
	}
	return message, nil
}

func (r *LearningRepository) FindAllMessages() ([]entity.MessageEntity, error) {
	var messages []entity.MessageEntity
	err := r.db.Find(&messages).Error
	return messages, err
}

func (r *LearningRepository) GetMessage(id int) (*entity.MessageEntity, error) {
	var message entity.MessageEntity
	err := r.db.First(&message, id).Error
	return &message, err
}

func (r *LearningRepository) DeleteMessage(id int) error {
	return r.db.Unscoped().Delete(&entity.MessageEntity{}, id).Error
}
