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

func (r *ProjectRepository) Update(project *models.Project, updateProject models.Project) error {
	if err := r.db.Model(&project).Session(&gorm.Session{FullSaveAssociations: true}).Updates(&updateProject).Error; err != nil {
		return err
	}

	return r.db.Model(&project).Preload("ProjectIssueStatuses.IssueStatus").Preload("ProjectAssignees").Where("id = ?", project.Id).First(&project).Error
}

func (r *Repository) UpdateFiles(project *models.Project) error {
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		if project.Logo != nil {
			if err := r.db.Model(&models.Project{}).Unscoped().Where("name = 'logo'").Association("Logo").Unscoped().Clear(); err != nil {
				return err
			}
		}

		if len(project.Files) > 0 {
			if err := r.db.Model(&models.Project{}).Unscoped().Where("name = 'files'").Association("Files").Unscoped().Clear(); err != nil {
				return err
			}
		}

		return r.db.Model(&project).Updates(&project).Error
	}); err != nil {
		return err
	}

	return r.db.Model(&project).Where("id = ?", project.Id).
		Preload("Logo", "name = 'logo'").Preload("Logo.AttachmentBlob").
		Preload("Files", "name = 'files'").Preload("Files.AttachmentBlob").
		First(&project).Error
}
