package repository

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"
	"fmt"
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

func (r *AttendanceRepository) List(
	attendances *[]*models.Attendance,
	paginateData *models.PaginationData,
	query insightInputs.AttendancesQueryInput,
) error {
	dbTables := r.db.Model(&models.Attendance{})

	return dbTables.Scopes(
		helpers.Paginate(
			dbTables.Scopes(
				r.userIdEq(query.UserIdEq),
				r.checkinAtGteq(query.CheckinAtGteq),
				r.checkinAtLteq(query.CheckinAtLteq),
			), paginateData,
		),
	).Order("id desc").Find(&attendances).Error
}

func (r *AttendanceRepository) AtDateOfUser(
	attendance *models.Attendance,
	user models.User,
	time time.Time,
) error {
	dbTable := r.db.Model(&models.Attendance{})

	return dbTable.Scopes(
		r.AtDate(time),
		r.OfUser(user.Id),
	).First(&attendance).Error
}

func (r *AttendanceRepository) CountAtDateOfUser(
	count *int64,
	userId int32,
	time time.Time,
	id *int32,
) error {
	dbTable := r.db.Model(&models.Attendance{})

	dbTable = dbTable.Scopes(
		r.AtDate(time),
		r.OfUser(userId),
	)

	if id != nil {
		dbTable = dbTable.Scopes(r.FilterOutByID(*id))
	}

	return dbTable.Count(count).Error
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

func (r *AttendanceRepository) Find(attendance *models.Attendance) error {
	dbTables := r.db.Model(&models.Attendance{})

	return dbTables.Where(&attendance).First(&attendance).Error
}

func (r *AttendanceRepository) OfUser(userId int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", userId)
	}
}

func (r *AttendanceRepository) AtDate(time time.Time) func(db *gorm.DB) *gorm.DB {
	dateOfTime := fmt.Sprintf("%d-%d-%d", time.Year(), time.Month(), time.Day())

	return func(db *gorm.DB) *gorm.DB {
		return db.Where("DATE(checkin_at) = ?", dateOfTime)
	}
}

func (r *AttendanceRepository) FilterOutByID(id int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id != ?", id)
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

func (r *AttendanceRepository) userIdEq(userIdEq *int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if userIdEq == nil {
			return db
		} else {
			return db.Where("user_id = ?", userIdEq)
		}
	}
}

func (r *AttendanceRepository) Create(attendance *models.Attendance) error {
	return r.db.Model(&attendance).Create(&attendance).Error
}

func (r *AttendanceRepository) Update(attendance *models.Attendance, updateAttendance models.Attendance) error {
	return r.db.Model(&attendance).Updates(updateAttendance).First(&attendance).Error
}

func (r *AttendanceRepository) Destroy(attendance *models.Attendance) error {
	return r.db.Model(&attendance).Delete(&attendance).Error
}
