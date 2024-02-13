package payloads

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

type MsSnippetCreateResolver struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args struct{ Input inputs.MsSnippetInput }

	model *models.Snippet
}

func (msc *MsSnippetCreateResolver) Resolve() error {
	user, err := auths.AuthUserFromCtx(*msc.Ctx)

	if err != nil {
		return exceptions.NewUnauthorizedError("")
	}

	msc.model = &models.Snippet{
		UserId: user.Id,
	}

	snippet := models.Snippet{
		UserId: user.Id,
	} // New

	form := validators.NewSnippetFormValidator(
		&msc.Args.Input,
		repository.NewSnippetRepository(msc.Ctx, msc.Db),
		&snippet,
	)

	form.Validate()

	if form.Valid {
		createErr := form.Create()

		if createErr != nil {
			return createErr
		}
	} else {
		return exceptions.NewUnprocessableContentError(nil, &form.Errors)
	}

	return nil
}

func (msc *MsSnippetCreateResolver) Snippet() *MsSnippetResolver {
	return &MsSnippetResolver{
		Snippet: msc.model,
	}
}
