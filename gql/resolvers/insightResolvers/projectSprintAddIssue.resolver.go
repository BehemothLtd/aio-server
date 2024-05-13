package insightResolvers

import (
	"aio-server/enums"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/services/insightServices"
	"context"
)

func (r *Resolver) ProjectSprintAddIssue(ctx context.Context, args insightInputs.ProjectSprintIssueModifyInput) (*string, error) {
	if _, err := r.Authorize(ctx, string(enums.PermissionTargetTypeProjectSprints), string(enums.PermissionActionTypeWrite)); err != nil {
		return nil, err
	}
	service := insightServices.ProjectSprintAddIssueService{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	}

	message := "Remove Issue out of sprint successfully"
	return &message, nil
}
