package insightResolvers

import (
	"aio-server/enums"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/services/insightServices"
	"context"
)

func (r *Resolver) UserCreate(ctx context.Context, args insightInputs.UserCreateInput) (*string, error) {
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

	if userInfo, err := service.Execute(); err != nil {
		return nil, err
	} else {
		return &userInfo, nil
	}
}
