package insightResolvers

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"

	"gorm.io/gorm"
)

func (r *Resolver) UserToggleActive(ctx context.Context, args insightInputs.UserInput) (*globalTypes.UserType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeUsers.String(), enums.PermissionActionTypeWrite.String()); err != nil {
		return nil, err
	}

	userId, err := helpers.GqlIdToInt32(args.Id)

	if err != nil {
		return nil, err
	}

	user := models.User{Id: userId}
	repo := repository.NewUserRepository(&ctx, r.Db)
	err = repo.Find(&user)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exceptions.NewRecordNotFoundError()
		}
		return nil, err
	}

	if user.State == enums.UserStateActive {
		user.State = enums.UserStateInactive
	} else {
		user.State = enums.UserStateActive
	}

	if err := repo.Update(&user, []string{"state"}); err != nil {
		return nil, exceptions.NewUnprocessableContentError("Unable to toggle user state", nil)
	}

	return &globalTypes.UserType{User: &user}, nil
}
