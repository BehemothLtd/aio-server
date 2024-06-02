package repository

import (
	"aio-server/models"
	"context"

	"gorm.io/gorm"
)

type DeviceUsingHistoryRepository struct {
	Repository
}

func NewDeviceUsingHistoryRepository(c *context.Context, db *gorm.DB) *DeviceUsingHistoryRepository {
	return &DeviceUsingHistoryRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

func (r *DeviceUsingHistoryRepository) Create(deviceUsingHistory *models.DevicesUsingHistory) error {
	return r.db.Model(&deviceUsingHistory).Create(&deviceUsingHistory).Error
}
