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

func (r *IssueRepository) List(
	issues *[]*models.Issue,
	query insightInputs.ProjectIssuesQueryInput,
	paginateData *models.PaginationData,
) error {
	dbTables := r.db.Model(&models.Issue{})

	return dbTables.Scopes(
		helpers.Paginate(dbTables.Scopes(
			r.projectIdEq(query.ProjectIdEq),
			r.titleLike(query.TitleCont),
			r.codeLike(query.CodeCont),
			r.issueTypeEq(query.IssueTypeEq),
			r.projectSprintIdEq(query.ProjectSprintIdEq),
			r.userIdIn(query.UserIdIn),
			r.deadLineAtGteq(query.DeadLineAtGteq),
			r.deadLineAtLteq(query.DeadLineAtLteq),
		), paginateData),
	).Order("id desc").Find(&issues).Error
}

func (r *IssueRepository) projectIdEq(projectId *int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if projectId == nil {
			return db
		} else {
			return db.Where("project_id = ?", projectId)
		}
	}
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

func (r *IssueRepository) issueTypeEq(issueTypeEq *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if issueTypeEq == nil {
			return db
		} else {
			issueTypeEqInInt, _ := enums.ParseIssueType(*issueTypeEq)

			return db.Where(gorm.Expr(`issues.issue_type = ?`, issueTypeEqInInt))
		}
	}
}

func (r *IssueRepository) projectSprintIdEq(projectSprintIdEq *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if projectSprintIdEq == nil {
			return db.Where("issues.project_sprint_id IS NULL")
		} else if *projectSprintIdEq == "" {
			return db
		} else {
			return db.Where(gorm.Expr(`issues.project_sprint_id = ?`, projectSprintIdEq))
		}
	}
}

func (r *IssueRepository) userIdIn(userIdIn *[]*int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if userIdIn == nil {
			return db.Where("NOT EXISTS (SELECT id FROM issue_assignees WHERE issue_assignees.issue_id = issues.id)")
		} else if len(*userIdIn) > 0 {
			var userIds []int32
			for _, id := range *userIdIn {
				userIds = append(userIds, *id)
			}
			return db.
				Joins("LEFT JOIN issue_assignees on issues.id = issue_assignees.issue_id").
				Where("issue_assignees.user_id IN (?)", userIds).
				Group("issues.id")
		} else {
			return db
		}
	}
}

func (r *IssueRepository) deadLineAtGteq(DeadLineAtGteq *time.Time) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if DeadLineAtGteq == nil {
			return db
		} else {
			return db.Where("deadline >= ?", DeadLineAtGteq)
		}
	}
}

func (r *IssueRepository) deadLineAtLteq(DeadLineAtLteq *time.Time) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if DeadLineAtLteq == nil {
			return db
		} else {
			return db.Where("deadline <= ?", DeadLineAtLteq)
		}
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

func (r *IssueRepository) CountByProjectAndIssueStatus(projectId int32, issueStatusId int32, count *int32) error {
	return r.db.
		Model(&models.Issue{}).
		Select("count(*)").
		Where("project_id = ? AND issue_status_id = ?", projectId, issueStatusId).
		Scan(&count).
		Error
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

func (r *IssueRepository) UpdateRemoveSprint(
	issue *models.Issue,
) error {
	return r.db.Model(issue).Update("project_sprint_id", nil).Error
}
