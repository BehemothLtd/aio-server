package payloads

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type MmUserGroupResolver struct {
	Ctx       *context.Context
	Db        *gorm.DB
	Args      struct{ Id graphql.ID }
	UserGroup *models.UserGroup
}

func (mugr *MmUserGroupResolver) ID(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(mugr.UserGroup.Id)
}

func (mugr *MmUserGroupResolver) Title(ctx context.Context) *string {
	return &mugr.UserGroup.Title
}

func (mugr *MmUserGroupResolver) CreatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(mugr.UserGroup.CreatedAt)
}

func (mugr *MmUserGroupResolver) UpdatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(mugr.UserGroup.UpdatedAt)
}
