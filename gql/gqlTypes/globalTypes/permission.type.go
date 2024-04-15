package globalTypes

import (
	"aio-server/pkg/helpers"
	"aio-server/pkg/systems"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type PermissionType struct {
	Ctx *context.Context
	Db  *gorm.DB

	Permission systems.Permission
}

func (ct *PermissionType) ID(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(ct.Permission.Id)
}

func (ct *PermissionType) Target(ctx context.Context) *string {
	target := ct.Permission.Target.String()
	return &target
}

func (ct *PermissionType) Action(ctx context.Context) *string {
	action := ct.Permission.Action.String()
	return &action
}
