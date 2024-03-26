package repository

import (
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
