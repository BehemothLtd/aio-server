package repository

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"
	"strings"
	"time"

	"gorm.io/gorm"
)

type WorkingTimelogRepository struct {
	Repository
}

func NewWorkingTimelogRepository(c *context.Context, db *gorm.DB) *WorkingTimelogRepository {
	return &WorkingTimelogRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

func (wtr *WorkingTimelogRepository) FindById(workingTimeLog *models.WorkingTimelog, id int32) error {
	dbTables := wtr.db.Model(&models.WorkingTimelog{})

	return dbTables.First(&workingTimeLog, id).Error
}

func (wtr *WorkingTimelogRepository) List(workingTimeLogs *[]*models.WorkingTimelog, paginateData *models.PaginationData, query insightInputs.WorkingTimelogsQueryInput) error {
	dbTables := wtr.db.Model(&models.WorkingTimelog{})

	return dbTables.Scopes(
		helpers.Paginate(dbTables.Scopes(
			wtr.DescriptionLike(query.DescriptionCont),
			wtr.IssueCodeEq(query.IssueCodeEq),
			wtr.IssueTitleLike(query.IssueTitleCont),
		), paginateData),
	).Order("id desc").Find(&workingTimeLogs).Error
}

func (wtr *WorkingTimelogRepository) Create(workingTimelog *models.WorkingTimelog) error {
	return wtr.db.Model(&workingTimelog).Create(&workingTimelog).First(&workingTimelog).Error
}

func (wtr *WorkingTimelogRepository) GetWorkingTimelogsByLoggedAt(workingTimeLogs *[]*models.WorkingTimelog, loggedAt time.Time) error {
	dbTables := wtr.db.Model(&models.WorkingTimelog{})

	dateOfLogging := loggedAt.Format("2006-01-02")

	return dbTables.Where("logged_at = ?", dateOfLogging).Find(&workingTimeLogs).Error
}

func (wtr *WorkingTimelogRepository) Update(workingTimelog *models.WorkingTimelog) error {
	return wtr.db.Model(&workingTimelog).Save(&workingTimelog).First(&workingTimelog).Error
}

// RANSACK
func (r *Repository) DescriptionLike(descriptionLike *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if descriptionLike == nil {
			return db
		} else {
			return db.Where(gorm.Expr(`lower(working_timelogs.description) LIKE ?`, strings.ToLower("%"+*descriptionLike+"%")))
		}
	}
}

func (r *Repository) IssueTitleLike(issueTitleLike *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if issueTitleLike == nil {
			return db
		} else {
			// return db.Where(gorm.Expr(`lower(issues.title) LIKE ?`, strings.ToLower("%"+*issueTitleLike+"%")))
			return db.Joins("LEFT JOIN issues on lower(issues.title) LIKE = ?", strings.ToLower("%"+*issueTitleLike+"%"))
		}
	}
}

func (r *Repository) IssueCodeEq(IssueCodeEq *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if IssueCodeEq == nil {
			return db
		} else {
			return db.Joins("LEFT JOIN issues on issues.id = ?", IssueCodeEq)
		}
	}
}
