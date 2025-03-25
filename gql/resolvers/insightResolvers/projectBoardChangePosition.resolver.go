package insightResolvers

import (
	"aio-server/enums"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/services/insightServices"
	"context"
)

func (r *Resolver) ProjectBoardChangePosition(ctx context.Context, args insightInputs.ProjectBoardChangePositionInput) (*string, error) {
	_, err := r.Authorize(ctx, enums.PermissionTargetTypeProjectIssues.String(), enums.PermissionActionTypeWrite.String())

	if err != nil {
		return nil, err
	}

	service := insightServices.ChangePositionBoardService{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	} else {
		message := "Successfully Change Position"
		return &message, nil
	}

}
