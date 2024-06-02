package globalTypes

import (
	"aio-server/models"
	"context"

	"gorm.io/gorm"
)

type DeviceUsingHistory struct {
	Ctx *context.Context
	Db  *gorm.DB

	DeviceUsingHistory *models.DevicesUsingHistory
}
