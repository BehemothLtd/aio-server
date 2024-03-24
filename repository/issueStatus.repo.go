package repository

import (
	"aio-server/enums"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
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
