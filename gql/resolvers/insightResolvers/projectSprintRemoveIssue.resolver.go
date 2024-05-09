package insightResolvers

import (
	"aio-server/enums"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/services/insightServices"
	"context"
)

func (r *Resolver) ProjectSprintRemoveIssue(ctx context.Context, args insightInputs.ProjectSprintIssueModifyInput) (*string, error) {
	if _, err := r.Authorize(ctx, string(enums.PermissionTargetTypeProjectSprints), string(enums.PermissionActionTypeWrite)); err != nil {
		return nil, err
	}
	service := insightServices.ProjectSprintRemoveIssueService{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	}

	return nil, nil
}
