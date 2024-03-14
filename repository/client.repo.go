package repository

import (
	"aio-server/models"
	"context"

	"gorm.io/gorm"
)

type ClientRepository struct {
	Repository
}

func NewClientRepository(c *context.Context, db *gorm.DB) *ClientRepository {
	return &ClientRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

func (r *ClientRepository) Find(client *models.Client) error {
	return r.db.Where(&client).First(&client).Error
}
