package repository

import (
	"aio-server/models"
	"context"

	"gorm.io/gorm"
)

type ProjectAssigneeRepository struct {
	Repository
}

func NewProjectAssigneeRepository(c *context.Context, db *gorm.DB) *ProjectAssigneeRepository {
	return &ProjectAssigneeRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

func (r *ProjectAssigneeRepository) Create(projectAssignee *models.ProjectAssignee) error {
	return r.db.Create(&projectAssignee).Error
}
