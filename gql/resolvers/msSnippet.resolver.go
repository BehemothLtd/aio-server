package resolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/inputs/msInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"

	"gorm.io/gorm"
)

// MsSnippet resolves the query for retrieving a single snippet.
func (r *Resolver) MsSnippet(ctx context.Context, args msInputs.SnippetInput) (*globalTypes.SnippetType, error) {
	if args.Id == "" {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}

	snippetId, err := helpers.GqlIdToInt32(args.Id)
	if err != nil {
		return nil, err
	}

	snippet := models.Snippet{}
	repo := repository.NewSnippetRepository(&ctx, r.Db)
	err = repo.FindById(&snippet, snippetId)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exceptions.NewRecordNotFoundError()
		}
		return nil, err
	}

	return &globalTypes.SnippetType{Snippet: &snippet}, nil
}
