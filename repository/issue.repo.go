package repository

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"
	"time"

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

type IssueCountingOnProjectAndState struct {
	Count         int
	ProjectId     int32
	IssueStatusId int32
}

func (r *IssueRepository) IssueCountingOnProjectAndState(
	issueCountingOnProjectAndState *[]IssueCountingOnProjectAndState,
	issueStatusIds []interface{},
	projectIds []interface{},
) error {
	return r.db.
		Model(models.Issue{}).
		Select("Count(id) as count, project_id, issue_status_id").
		Where("issue_status_id IN ? and project_id IN ?", issueStatusIds, projectIds).
		Group("project_id, issue_status_id").
		Order("project_id desc, issue_status_id asc").
		Scan(&issueCountingOnProjectAndState).
		Error
}

func (r *IssueRepository) UserAllWeekIssuesState(
	user models.User,
	issueDateBaseState *[]models.IssuesDeadlineBaseState,
) error {
	startTime, endTime := helpers.StartAndEndOfWeek(time.Now())

	return r.db.
		Model(models.Issue{}).
		Select("deadline as date, SUM(if(issue_status_id = 7, 1, 0)) as done, SUM(if(issue_status_id != 7, 1, 0)) as not_done").
		Preload("IssueStatus").
		Joins("LEFT JOIN issue_assignees ON issue_assignees.issue_id = issues.id").
		Where("deadline BETWEEN ? AND ?", startTime, endTime).
		Where("issue_assignees.user_id = ?", user.Id).
		Group("date").
		Order("date asc").
		Scan(&issueDateBaseState).
		Error
}
