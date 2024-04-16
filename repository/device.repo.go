package repository

import (
	"aio-server/enums"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"
	"strings"

	"gorm.io/gorm"
)

type DeviceRepository struct {
	Repository
}

func NewDeviceRepository(c *context.Context, db *gorm.DB) *DeviceRepository {
	return &DeviceRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

func (dr *DeviceRepository) deviceTypeIdIn(deviceTypeIdIn *[]int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if deviceTypeIdIn == nil {
			return db
		} else {
			return db.Where(gorm.Expr(`device_type_id IN (?)`, *deviceTypeIdIn))
		}
	}
}

func (dr *DeviceRepository) nameCont(nameCont *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if nameCont == nil {
			return db
		} else {
			return db.Where(gorm.Expr(`lower(devices.name) LIKE ?`, strings.ToLower("%"+*nameCont+"%")))
		}
	}
}

func (dr *DeviceRepository) stateIn(stateIn *[]string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if stateIn == nil {
			return db
		} else {
			stateInInInt := []enums.DeviceStateType{}
			for _, state := range *stateIn {
				stateInInt, err := enums.ParseDeviceStateType(state)
				if err == nil {
					stateInInInt = append(stateInInInt, stateInInt)
				}
			}

			return db.Where(gorm.Expr(`state IN (?)`, stateInInInt))
		}
	}
}

func (dr *DeviceRepository) userIdIn(userIdIn *[]int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if userIdIn == nil {
			return db
		} else {
			return db.Where(gorm.Expr(`devices.user_id IN (?)`, *userIdIn))
		}
	}
}

func (dr *DeviceRepository) List(
	devices *[]*models.Device,
	paginationData *models.PaginationData,
	query insightInputs.DevicesQueryInput,
) error {
	dbTable := dr.db.Model(&models.Device{})

	return dbTable.Scopes(
		helpers.Paginate(dbTable.Scopes(
			dr.deviceTypeIdIn(query.DeviceTypeIdIn),
			dr.nameCont(query.NameCont),
			dr.stateIn(query.StateIn),
			dr.userIdIn(query.UserIdIn),
		), paginationData),
	).Order("id desc").Find(&devices).Error
}

func (r *DeviceRepository) Find(device *models.Device) error {
	dbTables := r.db.Model(&models.Device{})

	return dbTables.Where(&device).First(&device).Error
}
