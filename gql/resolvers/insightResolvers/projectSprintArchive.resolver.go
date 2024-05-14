package insightResolvers

import (
	"aio-server/enums"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/services/insightServices"
	"context"
)

func (r *Resolver) ProjectSprintArchive(ctx context.Context, args insightInputs.ProjectSprintArchiveInput) (*string, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeProjectSprints.String(), enums.PermissionActionTypeWrite.String()); err != nil {
		return nil, err
	}

	service := insightServices.ProjectSprintArchiveService{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	} else {
		message := "Archived Sprint"

		return &message, nil
	}
}
