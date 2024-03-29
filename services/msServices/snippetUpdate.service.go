package msServices

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/msInputs"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"aio-server/validators"
	"context"

	"gorm.io/gorm"
)

// SnippetUpdateService handles the updation of snippets.
type SnippetUpdateService struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args msInputs.SnippetUpdateInput
}

// Execute updates an existed snippet.
func (sus *SnippetUpdateService) Execute() (*models.Snippet, error) {
	// Authenticate the user
	user, err := auths.AuthUserFromCtx(*sus.Ctx)
	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	snippetRepo := repository.NewSnippetRepository(sus.Ctx, sus.Db)

	snippetId, idTransErr := helpers.GqlIdToInt32(sus.Args.Id)

	if idTransErr != nil {
		return nil, exceptions.NewBadRequestError("Invalid ID")
	}

	snippet := models.Snippet{
		Id: snippetId,
	}
	snippetErr := snippetRepo.FindSnippetByAttr(&snippet)

	if snippetErr != nil {
		return nil, exceptions.NewRecordNotFoundError()
	}

	if snippet.UserId != user.Id {
		return nil, exceptions.NewRecordNotFoundError()
	}

	form := validators.NewSnippetFormValidator(
		&sus.Args.Input,
		snippetRepo,
		&snippet,
	)

	formErr := form.Save()

	if formErr != nil {
		return nil, formErr
	}

	return &snippet, nil
}
