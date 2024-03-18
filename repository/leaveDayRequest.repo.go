package repository

import (
	"aio-server/enums"
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

func (r *LeaveDayRequestRepository) FindById(request *models.LeaveDayRequest, id int32) error {
	dbTables := r.db.Model(&models.LeaveDayRequest{})

	return dbTables.First(&request, id).Error
}

func (ldr *LeaveDayRequestRepository) List(
	leaveDayRequests *[]*models.LeaveDayRequest,
	paginationData *models.PaginationData,
	query insightInputs.LeaveDayRequestsQueryInput,
) error {
	dbTables := ldr.db.Model(&models.LeaveDayRequest{}).Preload("User").Preload("Approver")

	return dbTables.Scopes(
		helpers.Paginate(dbTables.Scopes(
			ldr.requestTypeEq(query.RequestTypeEq),
			ldr.requestStateEq(query.RequestStateEq),
			ldr.userIdEq(query.UserIdEq),
		), paginationData),
	).Order("id desc").Find(&leaveDayRequests).Error
}

func (ldr *LeaveDayRequestRepository) requestTypeEq(requestTypeEq *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if requestTypeEq == nil {
			return db
		} else {
			requestTypeEqInInt, _ := enums.ParseRequestType(*requestTypeEq)

			return db.Where(gorm.Expr(`leave_day_requests.request_type = ?`, requestTypeEqInInt))
		}
	}
}

func (ldr *LeaveDayRequestRepository) requestStateEq(requestStateEq *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if requestStateEq == nil {
			return db
		} else {
			requestStateEqInInt, _ := enums.ParseRequestStateType(*requestStateEq)

			return db.Where(gorm.Expr(`leave_day_requests.request_state = ?`, requestStateEqInInt))
		}
	}
}

func (ldr *LeaveDayRequestRepository) userIdEq(userIdEq *int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if userIdEq == nil {
			return db
		} else {
			return db.Where(gorm.Expr(`leave_day_requests.user_id = ?`, userIdEq))
		}
	}
}
