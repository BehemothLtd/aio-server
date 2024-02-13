package payloads

import (
	"aio-server/exceptions"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"
	"slices"

	graphql "github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type MsSnippetResolver struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args struct{ Id graphql.ID }

	Snippet *models.Snippet
}

func (msr *MsSnippetResolver) Resolve() error {
	if msr.Args.Id == "" {
		return exceptions.NewBadRequestError("Invalid Id")
	}

	snippetId, err := helpers.GqlIdToInt32(msr.Args.Id)
	if err != nil {
		return err
	}

	snippet := models.Snippet{}

	repo := repository.NewSnippetRepository(msr.Ctx, msr.Db)
	snippetFindErr := repo.FindSnippetById(&snippet, snippetId)

	if snippetFindErr != nil {
		return exceptions.NewRecordNotFoundError()
	}

	msr.Snippet = &snippet

	return nil
}

func (msr *MsSnippetResolver) ID(context.Context) *graphql.ID {
	return helpers.GqlIDP(msr.Snippet.Id)
}

func (msr *MsSnippetResolver) Title(context.Context) *string {
	return &msr.Snippet.Title
}

func (msr *MsSnippetResolver) Content(context.Context) *string {
	return &msr.Snippet.Content
}

func (msr *MsSnippetResolver) UserId(context.Context) *graphql.ID {
	return helpers.GqlIDP(msr.Snippet.UserId)
}

func (msr *MsSnippetResolver) Slug(context.Context) *string {
	return &msr.Snippet.Slug
}

func (msr *MsSnippetResolver) SnippetType(context.Context) *int32 {
	snippetType := int32(msr.Snippet.SnippetType)

	return &snippetType
}

func (msr *MsSnippetResolver) FavoritesCount(context.Context) *int32 {
	favoritesCount := int32(msr.Snippet.FavoritesCount)

	return &favoritesCount
}

func (msr *MsSnippetResolver) CreatedAt(context.Context) *graphql.Time {
	createdAt := graphql.Time{Time: msr.Snippet.CreatedAt}

	return &createdAt
}

func (msr *MsSnippetResolver) UpdatedAt(context.Context) *graphql.Time {
	updatedAt := graphql.Time{Time: msr.Snippet.UpdatedAt}

	return &updatedAt
}

func (msr *MsSnippetResolver) Favorited(ctx context.Context) bool {
	user, err := auths.AuthUserFromCtx(ctx)

	if err != nil {
		return false
	}

	favorited := slices.ContainsFunc(msr.Snippet.FavoritedUsers, func(u models.User) bool { return u.Id == user.Id })

	return favorited
}
