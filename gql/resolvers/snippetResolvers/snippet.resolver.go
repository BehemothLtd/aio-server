package snippetResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/inputs/msInputs"
	"aio-server/models"
	"aio-server/repository"
	"context"

	"gorm.io/gorm"
)

// Snippet resolves the query for retrieving a single snippet.
func (r *Resolver) Snippet(ctx context.Context, args msInputs.SnippetInput) (*globalTypes.SnippetType, error) {
	if args.Slug == "" {
		return nil, exceptions.NewBadRequestError("Invalid Slug")
	}

	snippet := models.Snippet{
		Slug: args.Slug,
	}

	repo := repository.NewSnippetRepository(&ctx, r.Db.Preload("FavoritedUsers").Preload("Pins"))
	err := repo.FindSnippetByAttr(&snippet)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exceptions.NewRecordNotFoundError()
		}
		return nil, err
	}

	return &globalTypes.SnippetType{Snippet: &snippet}, nil
}
