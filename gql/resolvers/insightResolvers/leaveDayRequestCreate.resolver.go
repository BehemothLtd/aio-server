package insightResolvers

import (
	"aio-server/enums"
	"aio-server/models"
	"aio-server/services/insightServices"

	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"context"
)

func (r *Resolver) LeaveDayRequestCreate(ctx context.Context, args insightInputs.LeaveDayRequestCreateInput) (*insightTypes.LeaveDayRequestCreatedType, error) {
	user, err := r.Authorize(ctx, string(enums.PermissionTargetTypeLeaveDayRequests), string(enums.PermissionActionTypeWrite))
	if err != nil {
		return nil, err
	}

	request := models.LeaveDayRequest{
		UserId:       user.Id,
		RequestState: enums.RequestStateTypePending,
	}
	service := insightServices.LeaveDayRequestService{
		Ctx:     &ctx,
		Db:      r.Db,
		Args:    args,
		Request: &request,
	}

	if err := service.Excecute(); err != nil {
		return nil, err
	} else {
		// Redis Job to send message to slack channel

		return &insightTypes.LeaveDayRequestCreatedType{
			LeaveDayRequest: &globalTypes.LeaveDayRequestType{
				LeaveDayRequest: &request,
			},
		}, nil
	}
}
