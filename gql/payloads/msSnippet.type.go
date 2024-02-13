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

func (sr *MsSnippetResolver) Resolve() error {
	if sr.Args.Id == "" {
		return exceptions.NewBadRequestError("Invalid Id")
	}

	snippetId, err := helpers.GqlIdToInt32(sr.Args.Id)
	if err != nil {
		return err
	}

	snippet := models.Snippet{}

	repo := repository.NewSnippetRepository(sr.Ctx, sr.Db)
	snippetFindErr := repo.FindSnippetById(&snippet, snippetId)

	if snippetFindErr != nil {
		return exceptions.NewRecordNotFoundError()
	}

	sr.Snippet = &snippet

	return nil
}

func (sr *MsSnippetResolver) ID(context.Context) *graphql.ID {
	return helpers.GqlIDP(sr.Snippet.Id)
}

func (sr *MsSnippetResolver) Title(context.Context) *string {
	return &sr.Snippet.Title
}

func (sr *MsSnippetResolver) Content(context.Context) *string {
	return &sr.Snippet.Content
}

func (sr *MsSnippetResolver) UserId(context.Context) *graphql.ID {
	return helpers.GqlIDP(sr.Snippet.UserId)
}

func (sr *MsSnippetResolver) Slug(context.Context) *string {
	return &sr.Snippet.Slug
}

func (sr *MsSnippetResolver) SnippetType(context.Context) *int32 {
	snippetType := int32(sr.Snippet.SnippetType)

	return &snippetType
}

func (sr *MsSnippetResolver) FavoritesCount(context.Context) *int32 {
	favoritesCount := int32(sr.Snippet.FavoritesCount)

	return &favoritesCount
}

func (sr *MsSnippetResolver) CreatedAt(context.Context) *graphql.Time {
	createdAt := graphql.Time{Time: sr.Snippet.CreatedAt}

	return &createdAt
}

func (sr *MsSnippetResolver) UpdatedAt(context.Context) *graphql.Time {
	updatedAt := graphql.Time{Time: sr.Snippet.UpdatedAt}

	return &updatedAt
}

func (sr *MsSnippetResolver) Favorited(ctx context.Context) bool {
	user, err := auths.AuthUserFromCtx(ctx)

	if err != nil {
		return false
	}

	favorited := slices.ContainsFunc(sr.Snippet.FavoritedUsers, func(u models.User) bool { return u.Id == user.Id })

	return favorited
}
