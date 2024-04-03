package repository

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/constants"
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

func (wtr *WorkingTimelogRepository) FindByAttr(workingTimeLog *models.WorkingTimelog) error {
	findRecord := models.WorkingTimelog{
		ProjectId: workingTimeLog.ProjectId,
		IssueId:   workingTimeLog.IssueId,
		UserId:    workingTimeLog.UserId,
	}
	dbTables := wtr.db.Model(&findRecord)

	if !workingTimeLog.LoggedAt.IsZero() {
		dateOfLogging := workingTimeLog.LoggedAt.Format(constants.YYMMDD_DateFormat)
		dbTables = dbTables.Where("logged_at = ?", dateOfLogging)

	}

	return dbTables.Where(&findRecord).First(&workingTimeLog).Error
}

func (wtr *WorkingTimelogRepository) Update(workingTimelog *models.WorkingTimelog, updateRecord models.WorkingTimelog) error {
	return wtr.db.Model(&workingTimelog).Preload("User").Preload("Issue").Preload("Project").Updates(&updateRecord).First(&workingTimelog).Error
}

func (wtr *WorkingTimelogRepository) Create(workingTimelog *models.WorkingTimelog) error {
	return wtr.db.Model(workingTimelog).Preload("User").Preload("Issue").Preload("Project").Create(&workingTimelog).Error
}

func (wtr *WorkingTimelogRepository) GetWorkingTimelogsByLoggedAt(workingTimeLogs *[]*models.WorkingTimelog, loggedAt time.Time, id int32) error {
	dbTables := wtr.db.Model(&models.WorkingTimelog{})

	dateOfLogging := loggedAt.Format(constants.YYMMDD_DateFormat)

	return dbTables.Where("logged_at = ? AND id != ?", dateOfLogging, id).Find(&workingTimeLogs).Error
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
