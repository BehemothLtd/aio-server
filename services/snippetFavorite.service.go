package services

import (
	"aio-server/exceptions"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type SnippetFavoriteService struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args struct{ Id graphql.ID }

	user    models.User
	snippet models.Snippet

	Result *bool
}

func (sfs *SnippetFavoriteService) Execute() error {
	validationErr := sfs.validate()

	if validationErr != nil {
		return validationErr
	}

	snippetRepo := repository.NewSnippetRepository(sfs.Ctx, sfs.Db)
	snippetErr := snippetRepo.FindSnippetById(&sfs.snippet, sfs.snippet.Id)

	if snippetErr != nil {
		return exceptions.NewRecordNotFoundError()
	}

	favoriteSnippetRepo := repository.NewFavoriteSnippetRepository(sfs.Ctx, sfs.Db)
	favorited, toggleFavoriteErr := favoriteSnippetRepo.ToggleFavoriteSnippet(&sfs.snippet, &sfs.user)

	if toggleFavoriteErr != nil {
		return exceptions.NewUnprocessableContentError("Unable to perform this action", nil)
	}

	sfs.Result = &favorited

	return nil
}

func (sfs *SnippetFavoriteService) validate() error {
	if sfs.Args.Id == "" {
		return exceptions.NewBadRequestError("Invalid Id")
	}

	snippetId, err := helpers.GqlIdToInt32(sfs.Args.Id)
	if err != nil {
		return exceptions.NewBadRequestError("Invalid Id")
	}

	user, err := auths.AuthUserFromCtx(*sfs.Ctx)

	if err != nil {
		return exceptions.NewUnauthorizedError("")
	}

	sfs.user = user
	sfs.snippet = models.Snippet{
		Id: snippetId,
	}

	return nil
}
