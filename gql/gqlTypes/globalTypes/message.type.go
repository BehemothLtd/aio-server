package globalTypes

import (
	"aio-server/models"
	"context"

	"gorm.io/gorm"
)

type MessageType struct {
	Ctx *context.Context
	Db  *gorm.DB

	Message *models.Message
}
