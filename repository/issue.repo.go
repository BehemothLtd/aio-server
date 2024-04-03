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

func (r *IssueRepository) FindRecentTasksByUser(issues *[]models.Issue, userId int32) error {
	return r.db.Model(&models.Issue{}).
		Distinct("*").
		Joins("LEFT JOIN issue_assignees on issues.id = issue_assignees.issue_id").
		Where("user_id = ?", userId).
		Preload("IssueAssignees.User.Avatar", "name='avatar'").
		Preload("IssueAssignees.User.Avatar.AttachmentBlob").
		Limit(5).
		Find(&issues).Error
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
	originalIssue := models.Issue{Id: issue.Id}
	r.db.Model(&originalIssue).First(&originalIssue)

	r.db.Transaction(func(tx *gorm.DB) error {
		if err := r.db.Model(&originalIssue).Unscoped().Association("IssueAssignees").Unscoped().Clear(); err != nil {
			return err
		}

		if err := r.db.Model(&originalIssue).Association("IssueAssignees").Append(issue.IssueAssignees); err != nil {
			return err
		}

		return r.db.Model(&originalIssue).
			Preload("Creator").Preload("Project").
			Preload("ProjectSprint").
			Preload("Children").Preload("Parent").
			Preload("IssueAssignees").
			Save(&issue).First(&issue).Error
	})

	return nil
}
