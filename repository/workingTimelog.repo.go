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

	if query.IssueCodeEq != nil || query.IssueTitleCont != nil {
		dbTables = dbTables.Joins("LEFT OUTER JOIN issues on working_timelogs.issue_id = issues.id")
	}
	return dbTables.Scopes(
		helpers.Paginate(dbTables.Scopes(
			wtr.DescriptionLike(query.DescriptionCont),
			wtr.IssueCodeEq(query.IssueCodeEq),
			wtr.IssueTitleLike(query.IssueTitleCont),
		), paginateData),
	).Order("logged_at desc").Order("user_id desc").Find(&workingTimeLogs).Error
}

func (wtr *WorkingTimelogRepository) FindByAttr(workingTimeLog *models.WorkingTimelog) error {
	findRecord := models.WorkingTimelog{
		ProjectId: workingTimeLog.ProjectId,
		IssueId:   workingTimeLog.IssueId,
		UserId:    workingTimeLog.UserId,
	}
	dbTables := wtr.db.Model(&findRecord)

	if !workingTimeLog.LoggedAt.IsZero() {
		dateOfLogging := workingTimeLog.LoggedAt.Format(constants.YYYYMMDD_DateFormat)
		dbTables = dbTables.Where("logged_at = ?", dateOfLogging)

	}

	return dbTables.Where(&findRecord).First(&workingTimeLog).Error
}

func (wtr *WorkingTimelogRepository) Update(workingTimelog *models.WorkingTimelog, updateRecord models.WorkingTimelog) error {
	return wtr.db.Model(&workingTimelog).Preload("User").Preload("Issue").Preload("Project").Updates(&updateRecord).First(&workingTimelog).Error
}

func (wtr *WorkingTimelogRepository) Create(workingTimelog *models.WorkingTimelog) error {
	return wtr.db.Model(&workingTimelog).Preload("User").Preload("Issue").Preload("Project").Create(&workingTimelog).First(&workingTimelog).Error
}

func (wtr *WorkingTimelogRepository) GetWorkingTimelogsByLoggedAt(workingTimeLogs *[]*models.WorkingTimelog, loggedAt time.Time, id int32) error {
	dbTables := wtr.db.Model(&models.WorkingTimelog{})

	dateOfLogging := loggedAt.Format(constants.YYYYMMDD_DateFormat)

	return dbTables.Where("logged_at = ? AND id != ?", dateOfLogging, id).Find(&workingTimeLogs).Error
}

func (wtr *WorkingTimelogRepository) SelfWorkingTimelogHistory(
	workingTimelogHistory *[]*models.WorkingTimelogHistory,
	userId int32,
	query insightInputs.SelfWorkingTimelogQueryInput) error {
	return wtr.db.Model(&models.WorkingTimelog{}).Select(
		`working_timelogs.id as id,
		projects.id as project_id,
		issues.title as issue_name,
		issues.description as issue_description,
		working_timelogs.logged_at,
		working_timelogs.issue_id,
		working_timelogs.minutes`).
		Joins("INNER JOIN issues on working_timelogs.issue_id = issues.id").
		Joins("INNER JOIN projects on projects.id = working_timelogs.project_id").
		Scopes(
			wtr.UserIdEq(&userId),
			wtr.LoggedAtBetween(query.LoggedAtBetween),
		).
		Scan(&workingTimelogHistory).Error
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
			return db.Where(gorm.Expr(`lower(issues.title) LIKE ?`, strings.ToLower("%"+*issueTitleLike+"%")))
		}
	}
}

func (r *Repository) IssueCodeEq(IssueCodeEq *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if IssueCodeEq == nil {
			return db
		} else {
			return db.Where(gorm.Expr("issues.code = ?", IssueCodeEq))
		}
	}
}

func (r *Repository) UserIdEq(userIdEq *int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if userIdEq == nil {
			return db
		} else {
			return db.Where(gorm.Expr(`working_timelogs.user_id = ?`, userIdEq))
		}
	}
}

func (r *Repository) LoggedAtBetween(loggedAtBetween *[]*string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if loggedAtBetween == nil || len(*loggedAtBetween) == 0 {
			return db
		} else {
			if len(*loggedAtBetween) == 2 {
				dateRange := *loggedAtBetween
				startDateStr := dateRange[0]
				endDateStr := dateRange[1]
				query := db

				if startDateStr != nil && *startDateStr != "" {
					startDateTime, err := time.ParseInLocation(constants.YYYYMMDD_DateFormat, *startDateStr, time.Local)
					if err != nil {
						return db
					}

					query = query.Where(gorm.Expr(`working_timelogs.logged_at >= ?`, startDateTime))
				}
				if endDateStr != nil && *endDateStr != "" {
					endDateTime, err := time.ParseInLocation(constants.YYYYMMDD_DateFormat, *endDateStr, time.Local)
					if err != nil {
						return db
					}

					query = query.Where(gorm.Expr(`working_timelogs.logged_at <= ?`, endDateTime))
				}
				return query
			}

			return db
		}
	}
}
