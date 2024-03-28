package insightResolvers

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"
	"fmt"
)

func (r *Resolver) LeaveDayRequestDelete(ctx context.Context, args insightInputs.LeaveDayRequestInput) (*string, error) {
	user, err := r.Authorize(ctx, string(enums.PermissionTargetTypeLeaveDayRequests), string(enums.PermissionActionTypeWrite))

	if err != nil {
		return nil, err
	}
	if args.Id == "" {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}

	requestId, err := helpers.GqlIdToInt32(args.Id)
	if err != nil {
		return nil, err
	}

	leaveDayRequest := models.LeaveDayRequest{
		Id:     requestId,
		UserId: user.Id,
	}
	repo := repository.NewLeaveDayRequestRepository(&ctx, r.Db)

	if err := repo.Find(&leaveDayRequest); err != nil {
		return nil, exceptions.NewRecordNotFoundError()
	}
	if leaveDayRequest.RequestState != enums.RequestStateTypePending {
		return nil, exceptions.NewBadRequestError(fmt.Sprintf("Request was %s", leaveDayRequest.RequestState))
	}

	if err = repo.Destroy(&leaveDayRequest); err != nil {
		return nil, exceptions.NewBadRequestError(fmt.Sprintf("Can not delete this request %s", err.Error()))
	} else {
		message := "Deleted"

		return &message, nil
	}
}
