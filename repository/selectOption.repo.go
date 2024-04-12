package repository

import (
	"aio-server/models"
	"context"

	"gorm.io/gorm"
)

type SelectOptionRepository struct {
	Repository
}

func NewSelectOptionRepository(c *context.Context, db *gorm.DB) *SelectOptionRepository {
	return &SelectOptionRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

func (r *SelectOptionRepository) FetchUsers() ([]*models.User, error) {
	var users []*models.User
	err := r.db.Find(&users).Error
	return users, err
}
