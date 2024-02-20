package snippetResolvers

import (
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/snippetTypes"
	"aio-server/gql/inputs/msInputs"
	"aio-server/models"
	"aio-server/repository"
	"context"
)

// Snippets resolves the query for retrieving a collection of snippets.
func (r *Resolver) Snippets(ctx context.Context, args msInputs.SnippetsInput) (*snippetTypes.SnippetsType, error) {
	var snippets []*models.Snippet
	snippetsQuery, paginationData := args.ToPaginationDataAndSnippetsQuery()

	repo := repository.NewSnippetRepository(&ctx, r.Db)

	err := repo.List(&snippets, &paginationData, &snippetsQuery)
	if err != nil {
		return nil, err
	}

	return &snippetTypes.SnippetsType{
		Collection: r.SnippetSliceToTypes(snippets),
		Metadata: &globalTypes.MetadataType{
			Metadata: &paginationData.Metadata,
		},
	}, nil
}
