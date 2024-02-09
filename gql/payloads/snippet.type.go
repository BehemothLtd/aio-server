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

func (sr *SnippetResolver) UserId(context.Context) *graphql.ID {
	return helpers.GqlIDP(sr.Snippet.UserId)
}

func (sr *SnippetResolver) Slug(context.Context) *string {
	return &sr.Snippet.Slug
}

func (sr *SnippetResolver) SnippetType(context.Context) *int32 {
	snippetType := int32(sr.Snippet.SnippetType)

	return &snippetType
}

func (sr *SnippetResolver) FavoritesCount(context.Context) *int32 {
	favoritesCount := int32(sr.Snippet.FavoritesCount)

	return &favoritesCount
}

func (sr *SnippetResolver) CreatedAt(context.Context) *graphql.Time {
	createdAt := graphql.Time{Time: sr.Snippet.CreatedAt}

	return &createdAt
}

func (sr *SnippetResolver) UpdatedAt(context.Context) *graphql.Time {
	updatedAt := graphql.Time{Time: sr.Snippet.UpdatedAt}

	return &updatedAt
}
