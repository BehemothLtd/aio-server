package globalTypes

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type CollectionType struct {
	Ctx *context.Context
	Db  *gorm.DB

	Collection *models.Collection
}

func (ct *CollectionType) ID(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(ct.Collection.Id)
}

func (ct *CollectionType) Title(ctx context.Context) *string {
	return &ct.Collection.Title
}

func (ct *CollectionType) UserId(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(ct.Collection.UserId)
}

func (st *CollectionType) CreatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(st.Collection.CreatedAt)
}

func (st *CollectionType) UpdatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(st.Collection.UpdatedAt)
}

func (st *CollectionType) LockVersion(ctx context.Context) int32 {
	return st.Collection.LockVersion
}
