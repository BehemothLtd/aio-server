package snippetResolvers

import (
	"aio-server/gql/gqlTypes/globalTypes"

	"aio-server/gql/gqlTypes/snippetTypes"
	"aio-server/gql/inputs/msInputs"
	"aio-server/models"
	"aio-server/repository"
	"context"
)

func (r *Resolver) Collections(ctx context.Context, args msInputs.CollectionsInput) (*snippetTypes.CollectionsType, error) {
	var collections []*models.Collection
	collectionQuery, paginationData := args.ToPaginationDataAndCollectionQuery()

	repo := repository.NewCollectionRepository(&ctx, r.Db)

	err := repo.List(&collections, &paginationData, &collectionQuery)
	if err != nil {
		return nil, err
	}

	return &snippetTypes.CollectionsType{
		Collection: r.CollectionsSliceToTypes(collections),
		Metadata: &globalTypes.MetadataType{
			Metadata: &paginationData.Metadata,
		},
	}, nil
}
