package snippetResolvers

import (
	"aio-server/gql/gqlTypes/msTypes"
	"aio-server/gql/inputs/msInputs"
	"aio-server/services"
	"context"
)

// MsSnippetFavorite resolves the mutation for favoriting a snippet.
func (r *Resolver) SnippetFavorite(ctx context.Context, args msInputs.SnippetFavoriteInput) (*msTypes.MsSnippetFavoriteType, error) {
	service := services.SnippetFavoriteService{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	err := service.Execute()
	if err != nil {
		return nil, err
	}

	return &msTypes.MsSnippetFavoriteType{
		Favorited: *service.Result,
	}, nil
}
