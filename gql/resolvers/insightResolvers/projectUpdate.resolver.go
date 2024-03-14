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
	"context"
)

func (r *Resolver) ProjectUpdate(ctx context.Context, args insightInputs.ProjectUpdateInput) (*insightTypes.ProjectUpdatedType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeProjects.String(), enums.PermissionActionTypeWrite.String()); err != nil {
		return nil, err
	}

	if args.Id == "" {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}

	projectId, err := helpers.GqlIdToInt32(args.Id)

	if err != nil {
		return nil, err
	}

	project := models.Project{Id: projectId}
	service := insightServices.ProjectUpdateService{
		Ctx:     &ctx,
		Db:      r.Db,
		Args:    args,
		Project: &project,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	} else {
		return &insightTypes.ProjectUpdatedType{
			Project: &globalTypes.ProjectType{
				Project: &project,
			},
		}, nil
	}
}
