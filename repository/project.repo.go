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

// Find finds a project by its attributes.
func (r *ProjectRepository) Find(project *models.Project) error {
	dbTables := r.db.Model(&models.Project{})

	return dbTables.Where(&project).First(&project).Error
}

func (r *ProjectRepository) Create(project *models.Project) error {
	return r.db.Model(&project).Preload("ProjectIssueStatuses.IssueStatus").Preload("ProjectAssignees").Create(&project).First(&project).Error
}

func (r *ProjectRepository) Update(project *models.Project, fields []string) error {
	return r.db.Model(&project).Select(fields).Session(&gorm.Session{FullSaveAssociations: true}).Updates(&project).Error
}
