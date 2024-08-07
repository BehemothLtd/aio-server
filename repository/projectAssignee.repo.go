package repository

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
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
	if err := r.db.Create(&projectAssignee).Error; err != nil {
		return err
	}

	return r.db.Model(&models.ProjectAssignee{}).Where("id = ?", projectAssignee.Id).Preload("User.Avatar.AttachmentBlob").Find(&projectAssignee).Error
}

func (r *ProjectAssigneeRepository) Update(projectAssignee *models.ProjectAssignee, attributes map[string]interface{}) error {
	if err := r.db.Model(&projectAssignee).Select(append(helpers.GetKeys(attributes), "LockVersion")).Updates(attributes).Error; err != nil {
		return err
	}

	return r.db.Model(&models.ProjectAssignee{}).Where("id = ?", projectAssignee.Id).Preload("User.Avatar.AttachmentBlob").Find(&projectAssignee).Error
}

func (r *ProjectAssigneeRepository) Delete(projectAssignee *models.ProjectAssignee) error {
	return r.db.Where("id = ?", &projectAssignee.Id).Delete(&projectAssignee).Error
}
