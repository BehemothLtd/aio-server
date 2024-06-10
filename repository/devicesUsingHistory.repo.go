package repository

import (
	"aio-server/models"
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
