package insightResolvers

import (
	"aio-server/enums"
	"aio-server/models"
	"aio-server/services/insightServices"
	"aio-server/tasks"
	"fmt"
	"log"
	"time"

	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"context"

	"github.com/hibiken/asynq"
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
		// Send slack message
		var mentions []*string
		task, err := tasks.NewSlackSendLeaveDayRequestTask(request.Id, &mentions)
		if err != nil {
			log.Fatalf("could not create task: %v", err)
		}

		info, err := tasks.AsynqClient.Enqueue(task, asynq.ProcessIn(5*time.Second))
		if err != nil {
			log.Fatalf("could not enqueue task: %v", err)
		}

		fmt.Print(info)

		return nil, err
	} else {
		return &insightTypes.LeaveDayRequestCreatedType{
			LeaveDayRequest: &globalTypes.LeaveDayRequestType{
				LeaveDayRequest: &request,
			},
		}, nil
	}
}
