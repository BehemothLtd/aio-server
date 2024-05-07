package repository

import (
	"aio-server/models"
	"context"

	"gorm.io/gorm"
)

type MessageRepository struct {
	Repository
}

func NewMessageRepository(c *context.Context, db *gorm.DB) *MessageRepository {
	return &MessageRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

func (mr *MessageRepository) Find(message *models.RequestMessage) error {
	dbTables := mr.db.Model(&models.RequestMessage{})

	return dbTables.Where(&message).First(&message).Error
}

func (mr *MessageRepository) Create(message *models.RequestMessage) error {
	return mr.db.Table("request_messages").Create(&message).Error
}

func (mr *MessageRepository) Update(message *models.RequestMessage) error {
	originalMessage := models.RequestMessage{Id: message.Id}
	mr.db.Model(&originalMessage).First(&originalMessage)

	return mr.db.Table("request_messges").Save(&message).Error
}
