package insightResolvers

import (
	"aio-server/enums"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/services/insightServices"
	"context"
)

func (r *Resolver) ProjectUpdateProjectAssignee(ctx context.Context, args insightInputs.ProjectUpdateProjectAssigneeInput) (*insightTypes.ProjectAssigneeModificationType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeProjectAssignees.String(), enums.PermissionActionTypeWrite.String()); err != nil {
		return nil, err
	}

	service := insightServices.ProjectUpdateProjectAssigneeService{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	} else {
		return &insightTypes.ProjectAssigneeModificationType{
			ProjectAssignee: &globalTypes.ProjectAssigneeType{
				ProjectAssignee: service.ProjectAssignee,
			},
		}, nil
	}
}
