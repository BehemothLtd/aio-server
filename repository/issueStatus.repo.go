package repository

import (
	"aio-server/models"
	"context"

	"gorm.io/gorm"
)

type IssueStatusRepository struct {
	Repository
}

func NewIssueStatusRepository(c *context.Context, db *gorm.DB) *IssueStatusRepository {
	return &IssueStatusRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

// Find finds a issue status by its attributes.
func (r *IssueStatusRepository) Find(issueStatus *models.IssueStatus) error {
	dbTables := r.db.Model(&models.IssueStatus{})

	return dbTables.Where(&issueStatus).First(&issueStatus).Error
}
