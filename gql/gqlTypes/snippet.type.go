package gqlTypes

import (
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/pkg/helpers"
	"context"
	"slices"

	graphql "github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type SnippetResolver struct {
	Ctx *context.Context
	Db  *gorm.DB

	Snippet *models.Snippet
}

func (sr *SnippetResolver) ID(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(sr.Snippet.Id)
}

func (sr *SnippetResolver) Title(ctx context.Context) *string {
	return &sr.Snippet.Title
}

func (sr *SnippetResolver) Content(ctx context.Context) *string {
	return &sr.Snippet.Content
}

func (sr *SnippetResolver) UserId(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(sr.Snippet.UserId)
}

func (sr *SnippetResolver) Slug(ctx context.Context) *string {
	return &sr.Snippet.Slug
}

func (sr *SnippetResolver) SnippetType(ctx context.Context) *int32 {
	return helpers.Int32Pointer(int32(sr.Snippet.SnippetType))
}

func (sr *SnippetResolver) FavoritesCount(ctx context.Context) *int32 {
	return helpers.Int32Pointer(int32(sr.Snippet.FavoritesCount))
}

func (sr *SnippetResolver) CreatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(sr.Snippet.CreatedAt)
}

func (sr *SnippetResolver) UpdatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(sr.Snippet.UpdatedAt)
}

func (sr *SnippetResolver) LockVersion(ctx context.Context) int32 {
	return sr.Snippet.LockVersion
}

func (sr *SnippetResolver) Favorited(ctx context.Context) bool {
	user, err := auths.AuthUserFromCtx(ctx)

	if err != nil {
		return false
	}

	favorited := slices.ContainsFunc(sr.Snippet.FavoritedUsers, func(u models.User) bool { return u.Id == user.Id })
	return favorited
}
