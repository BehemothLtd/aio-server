package repository

import (
	"context"

	"gorm.io/gorm"
)

type ProjectRepository struct {
	Repository
}

func NewProjectRepository(c *context.Context, db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}
