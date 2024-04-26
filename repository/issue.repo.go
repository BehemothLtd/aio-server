package repository

import (
	"aio-server/enums"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"
	"strings"
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
func (r *IssueRepository) titleLike(titleLike *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if titleLike == nil {
			return db
		} else {
			return db.Where(gorm.Expr(`lower(issues.title) LIKE ?`, strings.ToLower("%"+*titleLike+"%")))
		}
	}
}

func (r *IssueRepository) codeLike(codeLike *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if codeLike == nil {
			return db
		} else {
			return db.Where(gorm.Expr(`lower(issues.code) LIKE ?`, strings.ToLower("%"+*codeLike+"%")))
		}
	}
}

func (r *Repository) issueTypeEq(IssueTypeEq *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if IssueTypeEq == nil {
			return db
		} else {
			IssueTypeInInt, _ := enums.ParseIssueType(*IssueTypeEq)

			return db.Where(gorm.Expr(`issues.Issue_type = ?`, IssueTypeInInt))
		}
	}
}
func (r *IssueRepository) List(
	issues *[]*models.Issue,
	paginateData *models.PaginationData,
	query insightInputs.IssuesQueryInput,
) error {
	dbTables := r.db.Model(&models.Issue{})

	return dbTables.Scopes(
		helpers.Paginate(dbTables.Scopes(r.titleLike(query.TitleCont), r.codeLike(query.CodeCont), r.issueTypeEq(query.IssueTypeEq)), paginateData),
	).Order("id desc").Find(&issues).Error
}
