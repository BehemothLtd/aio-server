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

func (r *IssueRepository) Find(issue *models.Issue) error {
	dbTables := r.db.Model(&models.Issue{})

	return dbTables.Where(&issue).First(&issue).Error
}

func (r *IssueRepository) Create(issue *models.Issue) error {
	return r.db.Model(&issue).
		Preload("Creator").Preload("Project").
		Preload("ProjectSprint").
		Preload("Children").Preload("Parent").
		Preload("IssueAssignees").
		Create(&issue).First(&issue).Error
}

func (r *IssueRepository) Update(issue *models.Issue) error {
	// db.Transaction(func(tx *gorm.DB) error {
	// 	if err := db.Model(&issue).Unscoped().Association("IssueAssignees").Unscoped().Clear(); err != nil {
	// 		return err
	// 	}

	// 	if err := db.Model(&issue).Association("IssueAssignees").Append(newIssueAssignees); err != nil {
	// 		return err
	// 	}

	// 	return nil
	// })
	return nil
}
