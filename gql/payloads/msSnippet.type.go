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
	Ctx     *context.Context
	Db      *gorm.DB
	Args    struct{ Id graphql.ID }
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
	err = repo.FindById(&snippet, snippetId)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return exceptions.NewRecordNotFoundError()
		}
		return err
	}

	msr.Snippet = &snippet

	return nil
}

func (msr *MsSnippetResolver) ID(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(msr.Snippet.Id)
}

func (msr *MsSnippetResolver) Title(ctx context.Context) *string {
	return &msr.Snippet.Title
}

func (msr *MsSnippetResolver) Content(ctx context.Context) *string {
	return &msr.Snippet.Content
}

func (msr *MsSnippetResolver) UserId(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(msr.Snippet.UserId)
}

func (msr *MsSnippetResolver) Slug(ctx context.Context) *string {
	return &msr.Snippet.Slug
}

func (msr *MsSnippetResolver) SnippetType(ctx context.Context) *int32 {
	return helpers.Int32Pointer(int32(msr.Snippet.SnippetType))
}

func (msr *MsSnippetResolver) FavoritesCount(ctx context.Context) *int32 {
	return helpers.Int32Pointer(int32(msr.Snippet.FavoritesCount))
}

func (msr *MsSnippetResolver) CreatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(msr.Snippet.CreatedAt)
}

func (msr *MsSnippetResolver) UpdatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(msr.Snippet.UpdatedAt)
}

func (msr *MsSnippetResolver) Favorited(ctx context.Context) bool {
	user, err := auths.AuthUserFromCtx(ctx)

	if err != nil {
		return false
	}

	favorited := slices.ContainsFunc(msr.Snippet.FavoritedUsers, func(u models.User) bool { return u.Id == user.Id })
	return favorited
}
