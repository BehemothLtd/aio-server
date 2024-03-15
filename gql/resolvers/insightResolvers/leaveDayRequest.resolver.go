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

func (r *Resolver) LeaveDayRequest(ctx context.Context, args insightInputs.LeaveDayRequestInput) (*globalTypes.LeaveDayRequestType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeLeaveDayRequests.String(), enums.PermissionActionTypeRead.String()); err != nil {
		return nil, err
	}

	if args.Id == "" {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}

	leaveDayRequestId, err := helpers.GqlIdToInt32(args.Id)

	if err != nil {
		return nil, err
	}

	leaveDayRequest := models.LeaveDayRequest{}
	repo := repository.NewLeaveDayRequestRepository(&ctx, r.Db)
	err = repo.FindById(&leaveDayRequest, leaveDayRequestId)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exceptions.NewRecordNotFoundError()
		}

		return nil, err
	}

	return &globalTypes.LeaveDayRequestType{LeaveDayRequest: &leaveDayRequest}, nil
}
