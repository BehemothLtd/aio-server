package payloads

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)

// SnippetResolver contains the DB and the model for resolving
type SnippetResolver struct {
	Snippet *models.Snippet
}

func (sr *SnippetResolver) ID(context.Context) *graphql.ID {
	return helpers.GqlIDP(sr.Snippet.Id)
}

func (sr *SnippetResolver) Title(context.Context) *string {
	return &sr.Snippet.Title
}

func (sr *SnippetResolver) Content(context.Context) *string {
	return &sr.Snippet.Content
}
