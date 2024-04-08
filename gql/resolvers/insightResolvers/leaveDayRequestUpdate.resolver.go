package insightResolvers

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/services/insightServices"
	"aio-server/tasks"
	"context"
	"fmt"
)

func (r *Resolver) LeaveDayRequestUpdate(ctx context.Context, args insightInputs.LeaveDayRequestUpdateInput) (*insightTypes.LeaveDayRequestUpdatedType, error) {
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

	request := models.LeaveDayRequest{Id: requestId, UserId: user.Id}
	service := insightServices.LeaveDayRequestUpdateService{
		Ctx:     &ctx,
		Db:      r.Db,
		Args:    args,
		Request: &request,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	} else {
		// Update slack request message task
		task, err := tasks.NewSlackUpdateLeaveDayRequestMessageTask(request)
		if err != nil {
			return nil, err
		}

		info, err := tasks.AsynqClient.Enqueue(task)
		if err != nil {
			return nil, err
		}
		fmt.Printf("Task ID: %+v => completed at %v\n", info.ID, info.CompletedAt)

		return &insightTypes.LeaveDayRequestUpdatedType{
			LeaveDayRequest: &globalTypes.LeaveDayRequestType{
				LeaveDayRequest: &request,
			},
		}, nil
	}
}
