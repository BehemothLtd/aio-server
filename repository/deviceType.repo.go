package repository

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	"gorm.io/gorm"
)

type DeviceTypeRepository struct {
	Repository
}

func NewDeviceTypeRepository(c *context.Context, db *gorm.DB) *DeviceTypeRepository {
	return &DeviceTypeRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

func (r *DeviceTypeRepository) List(
	deviceTypes *[]*models.DeviceType,
	paginateData *models.PaginationData,
) error {
	dbTable := r.db.Model(&models.DeviceType{})

	return dbTable.Scopes(
		helpers.Paginate(dbTable.Scopes(), paginateData),
	).Order("id desc").Find(&deviceTypes).Error
}

func (r *DeviceTypeRepository) FindById(deviceType *models.DeviceType, id int32) error {
	dbTables := r.db.Model(&models.DeviceType{})

	return dbTables.First(&deviceType, id).Error
}

func (r *DeviceTypeRepository) Create(deviceType *models.DeviceType) error {
	return r.db.Model(&deviceType).Create(&deviceType).Error
}
