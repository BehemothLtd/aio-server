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

// SnippetPinService handles pin and unpin snippets
type SnippetPinService struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args msInputs.SnippetPinInput

	user    models.User
	snippet models.Snippet
}

// Execute performs the pin/unpin action
func (sps *SnippetPinService) Execute() (*bool, error) {
	if err := sps.validate(); err != nil {
		return nil, err
	}

	sps.snippet = models.Snippet{
		Id: sps.snippet.Id,
	}

	// Retrive the snippet
	snippetRepo := repository.NewSnippetRepository(sps.Ctx, sps.Db)
	if err := snippetRepo.FindByAttr(&sps.snippet); err != nil {
		return nil, exceptions.NewRecordNotFoundError()
	}

	// Toggle pin
	pinRepo := repository.NewPinRepository(sps.Ctx, sps.Db)
	result, err := pinRepo.Toggle(sps.snippet, sps.user)
	if err != nil {
		return &result, exceptions.NewUnprocessableContentError("Unable to perform this action", nil)
	}

	return &result, nil
}

func (sps *SnippetPinService) validate() error {
	if sps.Args.Id == "" {
		return exceptions.NewBadRequestError("Invalid Id")
	}

	snippetId, err := helpers.GqlIdToInt32(sps.Args.Id)
	if err != nil {
		return exceptions.NewBadRequestError("Invalid Id")
	}

	// Authenticate user
	user, err := auths.AuthUserFromCtx(*sps.Ctx)
	if err != nil {
		return exceptions.NewUnauthorizedError("")
	}

	sps.user = user
	sps.snippet = models.Snippet{
		Id: snippetId,
	}
	return nil
}
