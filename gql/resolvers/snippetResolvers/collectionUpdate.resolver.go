package snippetResolvers

import (
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/snippetTypes"
	"aio-server/gql/inputs/msInputs"
	"aio-server/services/msServices"
	"context"
)

func (r *Resolver) CollectionUpdate(ctx context.Context, args msInputs.CollectionUpdateInput) (*snippetTypes.CollectionUpdatedType, error) {
	service := msServices.CollectionUpdateService{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	if collection, err := service.Execute(); err != nil {
		return nil, err
	} else {
		return &snippetTypes.CollectionUpdatedType{
			Collection: &globalTypes.CollectionType{
				Collection: collection,
			},
		}, nil
	}
}
