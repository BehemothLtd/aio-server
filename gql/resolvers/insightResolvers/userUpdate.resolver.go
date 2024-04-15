package insightResolvers

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/services/insightServices"
	"context"
)

func (r *Resolver) UserUpdate(ctx context.Context, args insightInputs.UserUpdateInput) (*globalTypes.UserUpdatedType, error) {
	_, err := r.Authorize(ctx, string(enums.PermissionTargetTypeUsers), string(enums.PermissionActionTypeWrite))
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	userId := args.Id
	updatedUser := models.User{Id: userId}

	service := insightServices.UserUpdateService{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
		User: &updatedUser,
	}
	if err := service.Execute(); err != nil {
		return nil, err
	}

	return &globalTypes.UserUpdatedType{
		User: &globalTypes.UserType{User: &updatedUser},
	}, nil
}
