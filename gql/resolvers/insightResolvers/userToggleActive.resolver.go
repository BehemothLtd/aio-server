package insightResolvers

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"
	"time"

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
	err = repo.FindWithProjectAssignees(&user)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exceptions.NewRecordNotFoundError()
		}
		return nil, err
	}

	currentTime := time.Now().Format(constants.DateTimeZoneFormat)

	if user.State == enums.UserStateActive {
		if !user.Inactiveable() {
			return nil, exceptions.NewBadRequestError("User is inactiveable!")
		}
		user.State = enums.UserStateInactive

		user.Timing.InactiveAt = currentTime
	} else {
		user.State = enums.UserStateActive

		timing := models.UserTiming{
			ActiveAt:   currentTime,
			InactiveAt: "",
		}
		user.Timing = &timing
	}

	if err := repo.Update(&user, []string{"state", "timing"}); err != nil {
		return nil, exceptions.NewUnprocessableContentError("Unable to toggle user state", nil)
	}

	return &globalTypes.UserType{User: &user}, nil
}
