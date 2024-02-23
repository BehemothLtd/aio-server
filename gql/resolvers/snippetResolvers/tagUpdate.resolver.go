package snippetResolvers

import (
	"aio-server/gql/gqlTypes/globalTypes"
	snippetTypes "aio-server/gql/gqlTypes/snippetTypes"
	"aio-server/gql/inputs/msInputs"
	"aio-server/services/msServices"
	"context"
)

func (r *Resolver) TagUpdate(ctx context.Context, args msInputs.TagUpdateInput) (*snippetTypes.TagUpdatedType, error) {
	service := msServices.TagUpdateService{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	if tag, err := service.Execute(); err != nil {
		return nil, err
	} else {
		return &snippetTypes.TagUpdatedType{
			Tag: &globalTypes.TagType{
				Tag: tag,
			},
		}, nil
	}
}
