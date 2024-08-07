package globalTypes

import (
	"aio-server/models"
	"aio-server/pkg/auths"
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

func (tt *TagType) Self(ctx context.Context) bool {
	user, err := auths.AuthUserFromCtx(ctx)

	if err != nil {
		return false
	}

	selfTag := tt.Tag.UserId == user.Id
	return selfTag
}

func (tt *TagType) LockVersion(ctx context.Context) int32 {
	return tt.Tag.LockVersion
}

func (tt *TagType) NumberOfUsage(ctx context.Context) int32 {
	count := len(tt.Tag.Snippets)

	return int32(count)
}
