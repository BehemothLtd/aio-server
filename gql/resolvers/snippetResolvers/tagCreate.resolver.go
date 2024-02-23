package snippetResolvers

import (
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/snippetTypes"
	"aio-server/gql/inputs/msInputs"
	"aio-server/services/msServices"
	"context"
)

func (r *Resolver) TagCreate(ctx context.Context, args msInputs.TagCreateInput) (*snippetTypes.TagCreatedType, error) {
	service := msServices.TagCreateService{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	if tag, err := service.Execute(); err != nil {
		return nil, err
	} else {
		return &snippetTypes.TagCreatedType{
			Tag: &globalTypes.TagType{
				Tag: tag,
			},
		}, nil
	}
}
