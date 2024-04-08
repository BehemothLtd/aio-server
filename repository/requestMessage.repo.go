package repository

import (
	"aio-server/models"
	"context"

	"gorm.io/gorm"
)

type RequestMessageRepository struct {
	Repository
}

func NewRequestMessageRepository(c *context.Context, db *gorm.DB) *RequestMessageRepository {
	return &RequestMessageRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

func (mr *RequestMessageRepository) Find(message *models.RequestMessage) error {
	dbTables := mr.db.Model(&models.RequestMessage{})

	return dbTables.Where(&message).First(&message).Error
}

func (mr *RequestMessageRepository) Create(message *models.RequestMessage) error {
	return mr.db.Table("request_messages").Create(&message).Error
}

func (mr *RequestMessageRepository) Update(message *models.RequestMessage) error {
	return mr.db.Table("request_messges").Update(&message).Error
}
