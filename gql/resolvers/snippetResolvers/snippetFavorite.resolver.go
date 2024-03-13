package snippetResolvers

import (
	snippetTypes "aio-server/gql/gqlTypes/snippetTypes"
	"aio-server/gql/inputs/msInputs"
	"aio-server/pkg/helpers"
	"aio-server/services/msServices"
	"context"
)

// SnippetFavorite resolves the mutation for favoriting a snippet.
func (r *Resolver) SnippetFavorite(ctx context.Context, args msInputs.SnippetFavoriteInput) (*snippetTypes.SnippetFavoriteType, error) {
	service := msServices.SnippetFavoriteService{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	err := service.Execute()
	if err != nil {
		return nil, err
	}

	return &snippetTypes.SnippetFavoriteType{
		Id:             *&service.Result.Id,
		Favorited:      *&service.Result.Favorited,
		FavoritesCount: *helpers.Int32Pointer(int32(service.Result.FavoritesCount)),
	}, nil
}
