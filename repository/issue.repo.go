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
			r.archivedEq(query.ArchivedEq),
			r.priorityEq(query.PriorityEq),
			r.issueStatusIdEq(query.IssueStatusIdEq),
			r.issueTypeEq(query.IssueTypeEq),
			r.projectSprintIdEq(query.ProjectSprintIdEq),
			r.userIdIn(query.UserIdIn),
			r.deadLineAtGteq(query.DeadLineAtGteq),
			r.deadLineAtLteq(query.DeadLineAtLteq),
		), paginateData),
	).Order("id desc").Find(&issues).Error
}

func (r *IssueRepository) AllByProjectId(issues *[]*models.Issue, projectId int32) error {
	return r.db.Model(&models.Issue{}).Scopes(r.projectIdEq(&projectId)).Order("id DESC").Find(&issues).Error
}

func (r *IssueRepository) ListByUser(
	issues *[]*models.Issue,
	query insightInputs.IssuesQueryInput,
	paginateData *models.PaginationData,
	user models.User,
) error {
	dbTables := r.db.Model(&models.Issue{})

	return dbTables.Scopes(
		helpers.Paginate(
			dbTables.Scopes(
				r.titleLike(query.TitleCont),
				r.issueTypeEq(query.IssueTypeEq),
				r.codeLike(query.CodeCont),
				r.OfUser(user.Id),
				r.projectIdEq(query.ProjectIdEq),
				r.deadLineAtGteq(query.DeadLineAtGteq),
				r.deadLineAtLteq(query.DeadLineAtLteq),
			), paginateData,
		),
	).Order("id desc").Find(&issues).Error
}

func (r *IssueRepository) ListByUserAndProject(
	issues *[]*models.Issue,
	userId int32,
	projectId int32,
) error {
	dbTables := r.db.Model(&models.Issue{})

	return dbTables.Scopes(
		r.OfUser(userId),
		r.projectIdEq(&projectId),
	).Order("id desc").Find(&issues).Error
}

func (r *IssueRepository) OfUser(userId int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Joins("LEFT JOIN issue_assignees ON issues.id = issue_assignees.issue_id").
			Where("issue_assignees.user_id = ?", userId).
			Group("issues.id")
	}
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

func (r *IssueRepository) issueStatusIdEq(issueStatusId *int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if issueStatusId == nil {
			return db
		} else {
			return db.Where("issues.issue_status_id = ?", issueStatusId)
		}
	}
}

func (r *IssueRepository) archivedEq(archived *bool) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if archived == nil {
			return db
		} else {
			return db.Where("issues.archived = ?", archived)
		}
	}
}

func (r *IssueRepository) priorityEq(priorityEq *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if priorityEq == nil {
			return db
		} else {
			priorityEqInInt, _ := enums.ParseIssuePriority(*priorityEq)

			return db.Where("issues.priority = ?", priorityEqInInt)
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

func (r *IssueRepository) userIdIn(userIdIn *[]int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if userIdIn == nil {
			return db
		} else {
			var userIds []int32
			includeUnassigned := false

			for _, id := range *userIdIn {
				if id == 0 {
					includeUnassigned = true
				} else {
					userIds = append(userIds, id)
				}
			}

			if includeUnassigned {
				if len(userIds) == 0 {
					return db.Where("NOT EXISTS (SELECT 1 FROM issue_assignees WHERE issue_assignees.issue_id = issues.id)")

				} else {
					return db.
						Joins("LEFT JOIN issue_assignees on issues.id = issue_assignees.issue_id").
						Where("issue_assignees.user_id IN (?) OR NOT EXISTS (SELECT 1 FROM issue_assignees WHERE issue_assignees.issue_id = issues.id)", userIds).
						Group("issues.id")
				}
			} else {
				return db.
					Joins("LEFT JOIN issue_assignees ON issues.id = issue_assignees.issue_id").
					Where("issue_assignees.user_id IN (?)", userIds).
					Group("issues.id")
			}
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

func (r *IssueRepository) Update(issue *models.Issue, updates map[string]interface{}) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Issue{Id: issue.Id}).Unscoped().Association("IssueAssignees").Unscoped().Clear(); err != nil {
			return err
		}

		if err := tx.Model(&models.Issue{Id: issue.Id}).Association("IssueAssignees").Append(issue.IssueAssignees); err != nil {
			return err
		}

		if err := tx.Model(&issue).Select(append(helpers.GetKeys(updates), "LockVersion")).Updates(updates).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return r.db.Model(&models.Issue{}).
		Where("id = ?", issue.Id).
		Preload("Creator").Preload("Project").
		Preload("ProjectSprint").
		Preload("Children").Preload("Parent").
		Preload("IssueAssignees").
		First(&issue).Error
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

func (r *IssueRepository) UpdateSprint(
	issue *models.Issue,
	sprintId int32,
) error {
	return r.db.Model(issue).Update("project_sprint_id", sprintId).Error
}

func (r IssueRepository) FetchProjectBoardIssues(project models.Project, issues *[]*models.Issue) error {
	dbTables := r.db.Model(&models.Issue{})

	// TODO: add page query
	scopes := dbTables.Scopes(
		r.projectIdEq(&project.Id),
	)

	if project.ProjectType == enums.ProjectTypeScrum {
		scopes = scopes.Where("project_sprint_id = ?", project.CurrentSprintId)
	}

	return scopes.Order("position asc").Find(&issues).Error
}

func (r *IssueRepository) FindMinAndMaxPositionWithCount(projectIssueStatus *models.ProjectIssueStatus) (*int, *int, *int, error) {
	var result struct {
		MinPosition int
		MaxPosition int
		Count       int
	}

	if err := r.db.Model(&models.Issue{}).
		Select("MIN(position) AS min_position, MAX(position) AS max_position, COUNT(*) AS count").
		Where("issue_status_id = ?", projectIssueStatus.IssueStatusId).
		Where("project_id = ?", projectIssueStatus.ProjectId).
		Scan(&result).
		Error; err != nil {
		return nil, nil, nil, err
	}

	return &result.MinPosition, &result.MaxPosition, &result.Count, nil
}

func (r *IssueRepository) FindIssueOfProjectByStatusAndPosition(projectIssueStatus *models.ProjectIssueStatus, position int) ([]models.Issue, error) {
	var issues []models.Issue

	if err := r.db.Model(&models.Issue{}).
		Where("issue_status_id = ?", projectIssueStatus.IssueStatusId).
		Where("project_id = ?", projectIssueStatus.ProjectId).
		Where("position >= ?", position).
		Find(&issues).
		Error; err != nil {
		return nil, err
	}

	return issues, nil
}

func (r *IssueRepository) UpdatePosition(
	issue *models.Issue,
	position int32,
	issueStatusId int32,
) error {
	return r.db.Model(issue).Updates(map[string]interface{}{"position": position, "issue_status_id": issueStatusId}).Error
}
