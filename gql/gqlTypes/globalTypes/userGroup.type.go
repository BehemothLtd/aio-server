package globalTypes

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)

type UserGroupType struct {
	UserGroup *models.UserGroup
}

func (ugt *UserGroupType) ID(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(ugt.UserGroup.Id)
}

func (ugt *UserGroupType) Title(ctx context.Context) *string {
	return &ugt.UserGroup.Title
}

func (ugt *UserGroupType) CreatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(ugt.UserGroup.CreatedAt)
}

func (ugt *UserGroupType) UpdatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(ugt.UserGroup.UpdatedAt)
}
