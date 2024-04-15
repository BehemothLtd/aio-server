package insightResolvers

import (
	"aio-server/enums"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/services/insightServices"
	"context"
)

func (r *Resolver) UserCreate(ctx context.Context, args insightInputs.UserCreateInput) (*globalTypes.UserType, error) {
	_, err := r.Authorize(ctx, string(enums.PermissionTargetTypeUsers), string(enums.PermissionActionTypeWrite))
	if err != nil {
		return nil, err
	}

	user := models.User{}
	service := insightServices.UserCreateService{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args.Input,
		User: &user,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	} else {
		return &globalTypes.UserType{
			User: &user,
		}, nil
	}
}
