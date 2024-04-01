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

func (r *Resolver) ProjectUploadImages(ctx context.Context, args insightInputs.ProjectUploadImagesInput) (*insightTypes.ProjectImagesUploadedType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeProjects.String(), enums.PermissionActionTypeWrite.String()); err != nil {
		return nil, err
	}

	project := models.Project{}
	service := insightServices.ProjectUploadImagesService{
		Ctx:     &ctx,
		Db:      r.Db,
		Args:    args,
		Project: &project,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	}

	return &insightTypes.ProjectImagesUploadedType{
		Project: &globalTypes.ProjectType{
			Project: &project,
		},
	}, nil
}
