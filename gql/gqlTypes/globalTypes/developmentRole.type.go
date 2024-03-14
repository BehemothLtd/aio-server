package globalTypes

import (
	"aio-server/pkg/helpers"
	"aio-server/pkg/systems"
	"context"

	"github.com/graph-gophers/graphql-go"
)

type DevelopentRoleType struct {
	DevelopmentRole *systems.DevelopmentRole
}

func (drt *DevelopentRoleType) ID(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(drt.DevelopmentRole.Id)
}

func (drt *DevelopentRoleType) Code(ctx context.Context) *string {
	return &drt.DevelopmentRole.Code
}

func (drt *DevelopentRoleType) Title(ctx context.Context) *string {
	return &drt.DevelopmentRole.Title
}
