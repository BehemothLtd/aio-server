package services

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/repository"
	"aio-server/validators"
	"context"

	"gorm.io/gorm"
)

type SnippetCreateService struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args struct{ Input inputs.MsSnippetInput }
}

func (scs *SnippetCreateService) Execute() (*models.Snippet, error) {
	user, err := auths.AuthUserFromCtx(*scs.Ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	snippet := models.Snippet{
		UserId: user.Id,
	} // New

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
