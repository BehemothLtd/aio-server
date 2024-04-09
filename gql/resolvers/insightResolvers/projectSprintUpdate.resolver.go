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

func (r *Resolver) ProjectSprintUpdate(ctx context.Context, args insightInputs.ProjectSprintUpdateInput) (*insightTypes.ProjectSprintType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeProjects.String(), enums.PermissionActionTypeWrite.String()); err != nil {
		return nil, err
	}

	if args.Id == "" {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}

	projectSprintId, err := helpers.GqlIdToInt32(args.Id)

	if err != nil {
		return nil, err
	}

	projectSprint := models.ProjectSprint{Id: projectSprintId}

	service := insightServices.ProjectSprintUpdateService{
		Ctx:           &ctx,
		Db:            r.Db,
		Args:          args,
		ProjectSprint: &projectSprint,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	} else {
		return &insightTypes.ProjectSprintType{
			ProjectSprint: &globalTypes.ProjectSprintType{
				ProjectSprint: &projectSprint,
			},
		}, nil
	}
}
