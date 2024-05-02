package insightResolvers

import (
	"aio-server/enums"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/services/insightServices"
	"context"
)

func (r *Resolver) ProjectDeleteProjectAssignee(ctx context.Context, args insightInputs.ProjectDeleteProjectAssigneeInput) (*string, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeProjectAssignees.String(), enums.PermissionActionTypeWrite.String()); err != nil {
		return nil, err
	}

	service := insightServices.ProjectDeleteProjectAssigneeService{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	} else {
		message := "Successfully deleted Member"
		return &message, nil
	}
}
