package resolvers

import (
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/msTypes"
	"aio-server/gql/inputs"
	"aio-server/models"
	"aio-server/repository"
	"context"
)

// MsSnippets resolves the query for retrieving a collection of snippets.
func (r *Resolver) MsSnippets(ctx context.Context, args inputs.MsSnippetsInput) (*msTypes.MsSnippetsType, error) {
	var snippets []*models.Snippet
	snippetsQuery, paginationData := args.ToPaginationDataAndSnippetsQuery()

	repo := repository.NewSnippetRepository(&ctx, r.Db)

	err := repo.List(&snippets, &paginationData, &snippetsQuery)
	if err != nil {
		return nil, err
	}

	return &msTypes.MsSnippetsType{
		Collection: r.SnippetSliceToTypes(snippets),
		Metadata: &globalTypes.MetadataType{
			Metadata: &paginationData.Metadata,
		},
	}, nil
}
