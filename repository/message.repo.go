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

func (mr *MessageRepository) Find(message *models.Message) error {
	dbTables := mr.db.Table("messages")

	return dbTables.Where(&message).First(&message).Error
}

func (mr *MessageRepository) Create(message *models.Message) error {
	return mr.db.Table("messages").Create(&message).Error
}

func (mr *MessageRepository) Update(message *models.Message) error {
	originalMessage := models.Message{Id: message.Id}
	mr.db.Table("messages").First(&originalMessage)

	return mr.db.Table("messages").Save(&message).Error
}
