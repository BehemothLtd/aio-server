package payloads

import (
	// "aio-server/exceptions"
	"aio-server/models"
	// "aio-server/pkg/auths"
	"aio-server/pkg/helpers"
	// "aio-server/repository"
	"context"
	// "slices"

	graphql "github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type MmUserGroupResolver struct {
	Ctx       *context.Context
	Db        *gorm.DB
	Args      struct{ Id graphql.ID }
	UserGroup *models.UserGroup
}

func (mug *MmUserGroupResolver) ID(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(mug.UserGroup.Id)
}

func (msr *MmUserGroupResolver) Title(ctx context.Context) *string {
	return &msr.UserGroup.Title
}

func (msr *MmUserGroupResolver) CreatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(msr.UserGroup.CreatedAt)
}

func (msr *MmUserGroupResolver) UpdatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(msr.UserGroup.UpdatedAt)
}
