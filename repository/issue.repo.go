package repository

import (
	"aio-server/models"
	"context"

	"gorm.io/gorm"
)

type IssueRepository struct {
	Repository
}

func NewIssueRepository(c *context.Context, db *gorm.DB) *IssueRepository {
	return &IssueRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

func (r *IssueRepository) Create(issue *models.Issue) error {
	return r.db.Model(&issue).
		Preload("Creator").Preload("Project").
		Preload("ProjectSprint").
		Preload("Children").Preload("Parent").
		Preload("IssueAssignees").
		Create(&issue).First(&issue).Error
}
