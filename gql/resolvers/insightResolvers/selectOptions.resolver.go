package insightResolvers

import (
	"aio-server/database"
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/pkg/auths"
	"aio-server/services/insightServices"
	"context"
)

func (r *Resolver) SelectOptions(ctx context.Context, args insightInputs.SelectOptionsInput) (*insightTypes.SelectOptionsType, error) {
	_, err := auths.AuthUserFromCtx(ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	service := insightServices.FetchSelectOptionsService{
		Db:   database.Db,
		Args: args,
	}

	if err := service.Execute(); err != nil {
		return nil, exceptions.NewBadRequestError("Error fetching options")
	}

	return nil, nil
}
