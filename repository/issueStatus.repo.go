package repository

import (
	"aio-server/enums"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/pkg/helpers"
	"context"
	"strings"

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

func (r *IssueStatusRepository) Create(issueStatus *models.IssueStatus) error {
	return r.db.Model(&issueStatus).Create(&issueStatus).First(&issueStatus).Error
}

func (r *IssueStatusRepository) Update(issueStatus *models.IssueStatus) error {
	originalIssueStatus := models.IssueStatus{Id: issueStatus.Id}
	r.db.Model(&originalIssueStatus).First(&originalIssueStatus)

	return r.db.Model(&originalIssueStatus).Save(&issueStatus).First(&issueStatus).Error
}

func (r *IssueStatusRepository) All(issueStatuses *[]*models.IssueStatus) error {
	return r.db.Table("issue_statuses").Order("id DESC").Find(&issueStatuses).Error
}

func (r *Repository) titleLike(titleLike *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if titleLike == nil {
			return db
		} else {
			return db.Where(gorm.Expr(`lower(issue_statuses.title) LIKE ?`, strings.ToLower("%"+*titleLike+"%")))
		}
	}
}

func (r *Repository) statusTypeEq(statusTypeEq *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if statusTypeEq == nil {
			return db
		} else {
			statusTypeInInt, _ := enums.ParseIssueStatusStatusType(*statusTypeEq)

			return db.Where(gorm.Expr(`issue_statuses.status_type = ?`, statusTypeInInt))
		}
	}
}

// List retrieves a list of issue statuses based on provided pagination data and query.
func (r *IssueStatusRepository) List(
	issueStatuses *[]*models.IssueStatus,
	paginateData *models.PaginationData,
	query insightInputs.IssueStatusesQueryInput,
) error {
	dbTables := r.db.Model(&models.IssueStatus{})

	return dbTables.Scopes(
		helpers.Paginate(dbTables.Scopes(
			r.titleLike(query.TitleCont),
			r.statusTypeEq(query.StatusTypeEq),
		), paginateData),
	).Order("id desc").Find(&issueStatuses).Error
}

func (r *IssueStatusRepository) DefaultScrum(issueStatus *[]models.IssueStatus) error {
	return r.db.Model(models.IssueStatus{}).
		Where("title IN ?", constants.ScrumDefaultIssueStatus()).
		Find(&issueStatus).
		Error
}

func (r *IssueStatusRepository) FetchIdsOnProject(projectId int32, ids *[]int32) error {
	return r.db.Model(&models.IssueStatus{}).
		Select("DISTINCT issue_statuses.id").
		Joins("LEFT JOIN project_issue_statuses ON project_issue_statuses.issue_status_id = issue_statuses.id").
		Joins("LEFT JOIN projects ON projects.id = project_issue_statuses.project_id").
		Where("projects.id = ?", projectId).Scan(&ids).Error
}
