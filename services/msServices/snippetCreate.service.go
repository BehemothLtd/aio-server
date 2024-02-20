package msServices

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/msInputs"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/repository"
	"aio-server/validators"
	"context"

	"gorm.io/gorm"
)

// SnippetCreateService handles the creation of snippets.
type SnippetCreateService struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args msInputs.SnippetCreateInput
}

// Execute creates a new snippet.
func (scs *SnippetCreateService) Execute() (*models.Snippet, error) {
	// Authenticate the user
	user, err := auths.AuthUserFromCtx(*scs.Ctx)
	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	snippet := models.Snippet{
		UserId: user.Id,
	}

	form := validators.NewSnippetFormValidator(
		&scs.Args.Input,
		repository.NewSnippetRepository(scs.Ctx, scs.Db),
		&snippet,
	)

	formErr := form.Save()

	if formErr != nil {
		return nil, formErr
	}

	return &snippet, nil
}
