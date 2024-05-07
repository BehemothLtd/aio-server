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

// Querying Functions
func (r *LeaveDayRequestRepository) Find(request *models.LeaveDayRequest) error {
	dbTables := r.db.Model(&models.LeaveDayRequest{})

	return dbTables.Where(&request).First(&request).Error
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

func (ldr *LeaveDayRequestRepository) Report(
	requestReport *[]*models.RequestReport,
	query insightInputs.RequestReportInput,
) error {
	// listRequestReports := []models.RequestReport{}
	// ldr.db.Table(models.LeaveDayRequest{}).Select("SUM( CASE When)")

	// return dbTables.Scopes(
	// 	ldr.userIdEq(query.UserIdEq),
	// 	ldr.requestTypeIn(query.RequestTypeIn),
	// 	ldr.createdAtBetween(query.CreatedAtBetween),
	// ).Order("user_id asc").Find()

	return nil
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

func (ldr *LeaveDayRequestRepository) requestTypeIn(requestTypeIn *[]*string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if requestTypeIn == nil || len(*requestTypeIn) == 0 {
			return db
		} else {
			var requestTypes []enums.RequestType

			for _, requestType := range *requestTypeIn {
				requestTypeInInt, err := enums.ParseRequestType(*requestType)

				if err != nil {
					continue
				}

				requestTypes = append(requestTypes, requestTypeInInt)
			}
			return db.Where(gorm.Expr(`leave_day_requests.request_type IN (?)`, requestTypes))
		}
	}
}

func (ldr *LeaveDayRequestRepository) createdAtBetween(createdAtBetween *[]*string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if createdAtBetween == nil || len(*createdAtBetween) == 0 {
			return db
		} else {
			if len(*createdAtBetween) == 2 {
				dateRange := *createdAtBetween
				startDateStr := dateRange[0]
				endDateStr := dateRange[1]

				query := db
				if startDateStr != nil || *startDateStr != "" {
					query = query.Where(gorm.Expr(`leave_day_requests.created_at >= ?`), startDateStr)
				}
				if endDateStr != nil || *endDateStr != "" {
					query = query.Where(gorm.Expr(`leave_day_requests.created_at <= ?`), endDateStr)
				}
				return query
			}

			return db
		}
	}
}

// Mutating Functions
func (ldr *LeaveDayRequestRepository) Create(request *models.LeaveDayRequest) error {
	return ldr.db.Table("leave_day_requests").Create(&request).Error
}

func (ldr *LeaveDayRequestRepository) Update(request *models.LeaveDayRequest) error {
	originalRequest := models.LeaveDayRequest{Id: request.Id}
	ldr.db.Model(&originalRequest).First(&originalRequest)

	return ldr.db.Model(&originalRequest).Save(&request).Error
}

func (ldr *LeaveDayRequestRepository) Destroy(request *models.LeaveDayRequest) error {
	return ldr.db.Table("leave_day_requests").Delete(&request).Error
}
