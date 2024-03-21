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

func (r *ProjectAssigneeRepository) Find(projectAssignee *models.ProjectAssignee) error {
	return r.db.Table("project_assignees").Where(&projectAssignee).First(&projectAssignee).Error
}

func (r *ProjectAssigneeRepository) Create(projectAssignee *models.ProjectAssignee) error {
	return r.db.Create(&projectAssignee).Error
}

func (r *ProjectAssigneeRepository) Update(projectAssignee *models.ProjectAssignee) error {
	projectAssignee.LockVersion += 1

	return r.db.Updates(&projectAssignee).Error
}
