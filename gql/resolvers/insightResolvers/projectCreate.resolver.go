package insightResolvers

import (
	"aio-server/enums"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/services/insightServices"
	"context"
)

func (r *Resolver) ProjectCreate(ctx context.Context, args insightInputs.ProjectCreateInput) (*insightTypes.ProjectCreatedType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeProjects.String(), enums.PermissionActionTypeWrite.String()); err != nil {
		return nil, err
	}

	project := models.Project{}

	service := insightServices.ProjectCreateService{
		Ctx:     &ctx,
		Db:      r.Db,
		Args:    args,
		Project: &project,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	} else {
		return &insightTypes.ProjectCreatedType{
			Project: &globalTypes.ProjectType{
				Project: &project,
			},
		}, nil
	}
}
