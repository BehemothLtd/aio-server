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

func (ct *CollectionType) Snippets(ctx context.Context) *[]*SnippetType {
	results := make([]*SnippetType, len(ct.Collection.Snippets))

	for i, s := range ct.Collection.Snippets {
		results[i] = &SnippetType{Snippet: s}
	}
	return &results
}

func (ct *CollectionType) CreatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(ct.Collection.CreatedAt)
}

func (ct *CollectionType) UpdatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(ct.Collection.UpdatedAt)
}
