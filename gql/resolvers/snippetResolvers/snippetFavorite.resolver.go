package snippetResolvers

import (
	snippetTypes "aio-server/gql/gqlTypes/snippetTypes"
	"aio-server/gql/inputs/msInputs"
	"aio-server/services"
	"context"
)

// MsSnippetFavorite resolves the mutation for favoriting a snippet.
func (r *Resolver) SnippetFavorite(ctx context.Context, args msInputs.SnippetFavoriteInput) (*snippetTypes.SnippetFavoriteType, error) {
	service := services.SnippetFavoriteService{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	err := service.Execute()
	if err != nil {
		return nil, err
	}

	return &snippetTypes.SnippetFavoriteType{
		Favorited: *service.Result,
	}, nil
}
