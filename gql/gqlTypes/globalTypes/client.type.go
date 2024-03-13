package globalTypes

import (
	"context"
	"aio-server/pkg/helpers"
	"aio-server/models"
	"gorm.io/gorm"

	graphql "github.com/graph-gophers/graphql-go"
)

type ClientType struct {
	Ctx    *context.Context
	Db     *gorm.DB

	Client *models.Client
}


func (ct *ClientType) ID(ctx context.Context) *graphql.ID {
    return helpers.GqlIDP(ct.Client.Id)
}

func (ct *ClientType) Name(ctx context.Context) *string {
    return &ct.Client.Name
}

func (ct *ClientType) ShowOnHomePage(ctx context.Context) *bool {
    return &ct.Client.ShowOnHomePage
}

func (ct *ClientType) LockVersion(ctx context.Context) *int32 {
    return &ct.Client.LockVersion
}

func (ct *ClientType) CreatedAt(ctx context.Context) *graphql.Time {
    return helpers.GqlTimePointer(ct.Client.CreatedAt)
}

func (ct *ClientType) UpdatedAt(ctx context.Context) *graphql.Time {
    return helpers.GqlTimePointer(ct.Client.UpdatedAt)
}
