package services

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"aio-server/validators"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type SnippetUpdateService struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args struct {
		Id    graphql.ID
		Input inputs.MsSnippetInput
	}
}

func (sus *SnippetUpdateService) Execute() (*models.Snippet, error) {
	user, err := auths.AuthUserFromCtx(*sus.Ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	snippetRepo := repository.NewSnippetRepository(sus.Ctx, sus.Db)

	snippetId, idTransErr := helpers.GqlIdToInt32(sus.Args.Id)

	if idTransErr != nil {
		return nil, exceptions.NewBadRequestError("Invalid ID")
	}

	snippet := models.Snippet{}
	snippetErr := snippetRepo.FindSnippetById(&snippet, snippetId)

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

	form.Validate()

	if form.Valid {
		updateErr := form.Update()

		if updateErr != nil {
			return nil, updateErr
		}

		return &snippet, nil
	} else {
		return nil, exceptions.NewUnprocessableContentError(nil, &form.Errors)
	}
}
