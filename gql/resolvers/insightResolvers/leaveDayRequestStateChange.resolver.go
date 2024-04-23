package insightResolvers

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"
	"fmt"
)

func (r *Resolver) LeaveDayRequestStateChange(ctx context.Context, args insightInputs.LeaveDayRequestStateChangeInput) (*insightTypes.LeaveDayRequestUpdatedType, error) {
	user, err := r.Authorize(ctx, string(enums.PermissionTargetTypeLeaveDayRequests), string(enums.PermissionActionTypeChangeState))

	if err != nil {
		return nil, err
	}

	if args.Id == "" {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}

	requetId, err := helpers.GqlIdToInt32(args.Id)
	if err != nil {
		return nil, err
	}

	if requestStateEnum, err := enums.ParseRequestStateType(args.RequestState); err != nil {
		return nil, exceptions.NewBadRequestError("Invalid request state")
	} else {
		request := models.LeaveDayRequest{
			Id: requetId,
		}

		repo := repository.NewLeaveDayRequestRepository(&ctx, r.Db)

		if err := repo.Find(&request); err != nil {
			return nil, exceptions.NewRecordNotFoundError()
		}

		request.ApproverId = &user.Id
		request.RequestState = requestStateEnum

		if err = repo.Update(&request); err != nil {
			return nil, exceptions.NewBadRequestError(fmt.Sprintf("Can not change this request's state %s", err.Error()))
		} else {

			response := &insightTypes.LeaveDayRequestUpdatedType{
				LeaveDayRequest: &globalTypes.LeaveDayRequestType{
					LeaveDayRequest: &request,
				}}

			return response, nil
		}
	}
}
