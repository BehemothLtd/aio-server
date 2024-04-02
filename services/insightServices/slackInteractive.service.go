package insightServices

import (
	"aio-server/models"
	"context"

	"gorm.io/gorm"
)

type SlackInteractiveService struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args models.SlackInteractivePayload
}
