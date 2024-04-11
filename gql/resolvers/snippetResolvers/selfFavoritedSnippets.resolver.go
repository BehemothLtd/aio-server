package snippetResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/snippetTypes"
	"aio-server/gql/inputs/msInputs"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/repository"
	"context"
)

func (r *Resolver) SelfFavoritedSnippets(ctx context.Context, args msInputs.SnippetsInput) (*snippetTypes.SnippetsType, error) {
	var snippets []*models.Snippet

	user, err := auths.AuthUserFromCtx(ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	snippetsQuery, paginationData := args.ToPaginationDataAndQuery()

	repo := repository.NewSnippetRepository(&ctx, r.Db)
	fetchErr := repo.ListByUserFavorited(&snippets, &paginationData, snippetsQuery, &user)

	if fetchErr != nil {
		return nil, exceptions.NewBadRequestError("")
	}

	return &snippetTypes.SnippetsType{
		Collection: r.SnippetSliceToTypes(snippets),
		Metadata: &globalTypes.MetadataType{
			Metadata: &paginationData.Metadata,
		},
	}, nil
}
