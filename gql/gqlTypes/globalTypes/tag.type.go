package globalTypes

import (
	"aio-server/models"
	"context"

	"gorm.io/gorm"
)

type TagType struct {
	Ctx *context.Context
	Db  *gorm.DB

	Tag *models.Tag
}
