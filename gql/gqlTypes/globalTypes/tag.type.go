package globalTypes

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type TagType struct {
	Ctx *context.Context
	Db  *gorm.DB

	Tag *models.Tag
}

func (tt *TagType) ID(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(tt.Tag.Id)
}

func (tt *TagType) Name(ctx context.Context) *string {
	return &tt.Tag.Name
}
