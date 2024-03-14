package repository

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	"gorm.io/gorm"
)

type LeaveDayRequestRepository struct {
	Repository
}

func NewLeaveDayRequestRepository(c *context.Context, db *gorm.DB) *LeaveDayRequestRepository {
	return &LeaveDayRequestRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

func (ldr *LeaveDayRequestRepository) List(
	leaveDayRequests *[]*models.LeaveDayRequest,
	paginationData *models.PaginationData,
	query insightInputs.LeaveDayRequestsQueryInput,
) error {
	dbTables := ldr.db.Model(&models.LeaveDayRequest{})

	return dbTables.Scopes(
		helpers.Paginate(dbTables.Scopes(), paginationData),
	).Order("id desc").Find(&leaveDayRequests).Error
}

func (ldr *LeaveDayRequestRepository) UserIdEq(userId *int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if userId == nil {
			return db
		} else {
			return db.Where(gorm.Expr(`leave_day_request.user_id = ?`, userId))
		}
	}
}
