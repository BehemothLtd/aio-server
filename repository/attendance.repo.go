package repository

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"
	"time"

	"gorm.io/gorm"
)

type AttendanceRepository struct {
	Repository
}

func NewAttendanceRepository(c *context.Context, db *gorm.DB) *AttendanceRepository {
	return &AttendanceRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

func (r *AttendanceRepository) ListByUser(
	attendances *[]*models.Attendance,
	paginateData *models.PaginationData,
	query insightInputs.SelfAttendancesQueryInput,
	user models.User,
) error {
	dbTables := r.db.Model(&models.Attendance{})

	return dbTables.Scopes(
		helpers.Paginate(
			dbTables.Scopes(
				r.OfUser(user.Id),
				r.checkinAtGteq(query.CheckinAtGteq),
				r.checkinAtLteq(query.CheckinAtLteq),
			), paginateData,
		),
	).Order("id desc").Find(&attendances).Error
}

func (r *AttendanceRepository) OfUser(userId int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", userId)
	}
}

func (r *AttendanceRepository) checkinAtGteq(checkinAtGteq *time.Time) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if checkinAtGteq == nil {
			return db
		} else {
			return db.Where("checkin_at >= ?", checkinAtGteq)
		}
	}
}

func (r *AttendanceRepository) checkinAtLteq(checkinAtLteq *time.Time) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if checkinAtLteq == nil {
			return db
		} else {
			return db.Where("checkin_at <= ?", checkinAtLteq)
		}
	}
}
