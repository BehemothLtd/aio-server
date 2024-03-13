package repository

import (
	"aio-server/models"
	"context"

	"gorm.io/gorm"
)

type ProjectSprintRepository struct {
	Repository
}

func NewProjectSprintRepository(c *context.Context, db *gorm.DB) *ProjectSprintRepository {
	return &ProjectSprintRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

func (r *ProjectSprintRepository) Find(projectSprint *models.ProjectSprint) error {
	dbTables := r.db.Model(&models.ProjectSprint{})

	return dbTables.Where(&projectSprint).First(&projectSprint).Error
}
