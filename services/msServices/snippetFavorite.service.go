package msServices

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/msInputs"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"

	"gorm.io/gorm"
)

// SnippetFavoriteService handles favoriting and unfavoriting snippets.
type SnippetFavoriteService struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args msInputs.SnippetFavoriteInput

	user    models.User
	snippet models.Snippet
	Result  *bool
}

// Execute performs the favoriting or unfavoriting action.
func (sfs *SnippetFavoriteService) Execute() error {
	if err := sfs.validate(); err != nil {
		return err
	}
	snippet := models.Snippet{Id: sfs.snippet.Id}
	// Retrieve the snippet
	snippetRepo := repository.NewSnippetRepository(sfs.Ctx, sfs.Db)
	if err := snippetRepo.FindByAttr(&snippet); err != nil {
		return exceptions.NewRecordNotFoundError()
	}

	// Toggle favorite status
	favoriteSnippetRepo := repository.NewSnippetFavoriteRepository(sfs.Ctx, sfs.Db)
	favorited, err := favoriteSnippetRepo.Toggle(&sfs.snippet, &sfs.user)
	if err != nil {
		return exceptions.NewUnprocessableContentError("Unable to perform this action", nil)
	}

	sfs.Result = &favorited
	return nil
}

// validate validates the input and retrieves the user information.
func (sfs *SnippetFavoriteService) validate() error {
	if sfs.Args.Id == "" {
		return exceptions.NewBadRequestError("Invalid Id")
	}

	snippetId, err := helpers.GqlIdToInt32(sfs.Args.Id)
	if err != nil {
		return exceptions.NewBadRequestError("Invalid Id")
	}

	// Authenticate user
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
