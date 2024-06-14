package repository

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	"gorm.io/gorm"
)

type DevicesUsingHistoryRepository struct {
	Repository
}

func NewDevicesUsingHistoryRepository(c *context.Context, db *gorm.DB) *DevicesUsingHistoryRepository {
	return &DevicesUsingHistoryRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

func (r *DevicesUsingHistoryRepository) Create(devicesUsingHistory *models.DevicesUsingHistory) error {
	return r.db.Model(&devicesUsingHistory).Create(&devicesUsingHistory).Error
}

func (r *DevicesUsingHistoryRepository) deviceIdIn(deviceIdIn *[]int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if deviceIdIn == nil {
			return db
		} else {
			return db.Where(gorm.Expr(`device_id IN (?)`, *deviceIdIn))
		}
	}
}

func (r *DevicesUsingHistoryRepository) List(
	devicesUsingHistories *[]*models.DevicesUsingHistory,
	paginationData *models.PaginationData,
	query insightInputs.DevicesUsingHistoriesQueryInput,
) error {
	dbTable := r.db.Model(&models.DevicesUsingHistory{})

	return dbTable.Scopes(
		helpers.Paginate(dbTable.Scopes(
			r.deviceIdIn(query.DeviceIdIn),
		), paginationData),
	).Order("id desc").Find(&devicesUsingHistories).Error
}
