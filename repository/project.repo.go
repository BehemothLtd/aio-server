package repository

import (
	"aio-server/models"
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

// FindById finds a project by its attributes.
func (r *ProjectRepository) Find(project *models.Project) error {
	dbTables := r.db.Model(&models.Project{})

	return dbTables.Where(&project).First(&project).Error
}
