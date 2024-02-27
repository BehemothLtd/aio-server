package snippetResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/pkg/auths"

	"aio-server/gql/gqlTypes/snippetTypes"
	"aio-server/gql/inputs/msInputs"
	"aio-server/models"
	"aio-server/repository"
	"context"
)

func (r *Resolver) Collections(ctx context.Context, args msInputs.CollectionsInput) (*snippetTypes.CollectionsType, error) {
	var collections []*models.Collection

	user, err := auths.AuthUserFromCtx(ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	collectionQuery, paginationData := args.ToPaginationDataAndCollectionQuery()

	repo := repository.NewCollectionRepository(&ctx, r.Db)

	fetchErr := repo.List(&collections, &paginationData, &collectionQuery, &user)

	if fetchErr != nil {
		return nil, exceptions.NewBadRequestError("")
	}

	return &snippetTypes.CollectionsType{
		Collection: r.CollectionsSliceToTypes(collections),
		Metadata: &globalTypes.MetadataType{
			Metadata: &paginationData.Metadata,
		},
	}, nil
}
