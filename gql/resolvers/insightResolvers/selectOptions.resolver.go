package insightResolvers

import (
	"aio-server/database"
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/pkg/auths"
	"aio-server/services/insightServices"
	"context"
	"fmt"
)

func (r *Resolver) SelectOptions(ctx context.Context, args insightInputs.SelectOptionsInput) (*insightTypes.SelectOptionsType, error) {
	user, err := auths.AuthUserFromCtx(ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	service := insightServices.FetchSelectOptionsService{
		Db:     database.Db,
		Keys:   args.Input.Keys,
		Params: args.Params,
		User:   &user,
	}

	if err := service.Execute(); err != nil {
		return nil, exceptions.NewBadRequestError(fmt.Sprintf("Error fetching options : %s", err.Error()))
	}

	return &service.Result, nil
}
