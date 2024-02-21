package globalTypes

import (
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/pkg/helpers"
	"context"
	"slices"

	graphql "github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type SnippetType struct {
	Ctx *context.Context
	Db  *gorm.DB

	Snippet *models.Snippet
}

func (sr *SnippetType) ID(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(sr.Snippet.Id)
}

func (sr *SnippetType) Title(ctx context.Context) *string {
	return &sr.Snippet.Title
}

func (sr *SnippetType) Content(ctx context.Context) *string {
	return &sr.Snippet.Content
}

func (sr *SnippetType) UserId(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(sr.Snippet.UserId)
}

func (sr *SnippetType) Slug(ctx context.Context) *string {
	return &sr.Snippet.Slug
}

func (sr *SnippetType) SnippetType(ctx context.Context) *string {
	value := sr.Snippet.SnippetType.String()

	return &value
}

func (sr *SnippetType) FavoritesCount(ctx context.Context) *int32 {
	return helpers.Int32Pointer(int32(sr.Snippet.FavoritesCount))
}

func (sr *SnippetType) CreatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(sr.Snippet.CreatedAt)
}

func (sr *SnippetType) UpdatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(sr.Snippet.UpdatedAt)
}

func (sr *SnippetType) LockVersion(ctx context.Context) int32 {
	return sr.Snippet.LockVersion
}

func (sr *SnippetType) Favorited(ctx context.Context) bool {
	user, err := auths.AuthUserFromCtx(ctx)

	if err != nil {
		return false
	}

	favorited := slices.ContainsFunc(sr.Snippet.FavoritedUsers, func(u models.User) bool { return u.Id == user.Id })
	return favorited
}
