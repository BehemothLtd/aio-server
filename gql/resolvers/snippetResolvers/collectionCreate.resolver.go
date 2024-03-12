package snippetResolvers

import (
	"aio-server/gql/gqlTypes/globalTypes"
	snippetTypes "aio-server/gql/gqlTypes/snippetTypes"
	"aio-server/gql/inputs/msInputs"
	"aio-server/services/msServices"
	"context"
)

func (r *Resolver) CollectionCreate(ctx context.Context, args msInputs.CollectionCreateInput) (*snippetTypes.CollectionCreatedType, error) {
	service := msServices.CollectionCreateService{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	if collection, err := service.Excecute(); err != nil {
		return nil, err
	} else {
		return &snippetTypes.CollectionCreatedType{
			Collection: &globalTypes.CollectionType{
				Collection: collection,
			},
		}, nil
	}
}
