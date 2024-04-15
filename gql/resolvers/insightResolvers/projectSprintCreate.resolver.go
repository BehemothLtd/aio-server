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

func (r *Resolver) ProjectSprintCreate(ctx context.Context, args insightInputs.ProjectSprintCreateInput) (*insightTypes.ProjectSprintType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeProjects.String(), enums.PermissionActionTypeWrite.String()); err != nil {
		return nil, err
	}
	projectSprint := models.ProjectSprint{}
	services := insightServices.ProjectSprintCreateService{
		Ctx:           &ctx,
		Db:            r.Db,
		Args:          args,
		ProjectSprint: &projectSprint,
	}

	if err := services.Execute(); err != nil {
		return nil, err
	} else {
		return &insightTypes.ProjectSprintType{
			ProjectSprint: &globalTypes.ProjectSprintType{
				ProjectSprint: &projectSprint,
			},
		}, nil
	}
}
