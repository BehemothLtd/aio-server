package ms

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

// SnippetResolver contains the DB and the model for resolving
type SnippetResolver struct {
	Db  *gorm.DB
	Ctx *context.Context
	M   *models.Snippet
}

func (sr *SnippetResolver) ID(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(sr.M.Id)
}

func (sr *SnippetResolver) Title(ctx context.Context) *string {
	return &sr.M.Title
}

func (sr *SnippetResolver) Content(ctx context.Context) *string {
	return &sr.M.Content
}
