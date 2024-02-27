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

func (st *SnippetType) ID(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(st.Snippet.Id)
}

func (st *SnippetType) Title(ctx context.Context) *string {
	return &st.Snippet.Title
}

func (st *SnippetType) Content(ctx context.Context) *string {
	return &st.Snippet.Content
}

func (st *SnippetType) UserId(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(st.Snippet.UserId)
}

func (st *SnippetType) Slug(ctx context.Context) *string {
	return &st.Snippet.Slug
}

func (st *SnippetType) SnippetType(ctx context.Context) *string {
	value := st.Snippet.SnippetType.String()

	return &value
}

func (st *SnippetType) FavoritesCount(ctx context.Context) *int32 {
	return helpers.Int32Pointer(int32(st.Snippet.FavoritesCount))
}

func (st *SnippetType) CreatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(st.Snippet.CreatedAt)
}

func (st *SnippetType) UpdatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(st.Snippet.UpdatedAt)
}

func (st *SnippetType) LockVersion(ctx context.Context) int32 {
	return st.Snippet.LockVersion
}

func (st *SnippetType) Favorited(ctx context.Context) bool {
	user, err := auths.AuthUserFromCtx(ctx)

	if err != nil {
		return false
	}

	favorited := slices.ContainsFunc(st.Snippet.FavoritedUsers, func(u models.User) bool { return u.Id == user.Id })
	return favorited
}

func (st *SnippetType) Pinned(ctx context.Context) bool {
	user, err := auths.AuthUserFromCtx(ctx)

	if err != nil {
		return false
	}

	pinned := slices.ContainsFunc(st.Snippet.Pins, func(p models.Pin) bool { return p.UserId == user.Id })
	return pinned
}
