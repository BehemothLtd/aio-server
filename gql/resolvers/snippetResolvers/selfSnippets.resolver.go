package snippetResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	snippetTypes "aio-server/gql/gqlTypes/snippetTypes"
	"aio-server/gql/inputs/msInputs"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/repository"
	"context"
)

// SelfSnippets resolves the query for retrieving self-owned snippets.
func (r *Resolver) SelfSnippets(ctx context.Context, args msInputs.SnippetsInput) (*snippetTypes.SnippetsType, error) {
	var snippets []*models.Snippet

	user, err := auths.AuthUserFromCtx(ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	snippetsQuery, paginationData := args.ToPaginationDataAndSnippetsQuery()

	repo := repository.NewSnippetRepository(&ctx, r.Db.Model(&models.Snippet{}).Preload("FavoritedUsers").Preload("Pins"))

	fetchErr := repo.ListByUser(&snippets, &paginationData, &snippetsQuery, &user)

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
