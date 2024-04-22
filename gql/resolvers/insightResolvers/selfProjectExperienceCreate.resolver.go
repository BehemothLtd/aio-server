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

func (r *Resolver) SelfProjectExperienceCreate(ctx context.Context, args insightInputs.ProjectExperienceCreateInput) (*insightTypes.ProjectExperienceType, error) {
	user, err := r.Authorize(ctx, enums.PermissionTargetTypeProjectExperiences.String(), enums.PermissionActionTypeWrite.String())
	if err != nil {
		return nil, err
	}
	userId := user.Id

	projectExperience := models.ProjectExperience{UserId: userId}
	service := insightServices.SelfProjectExperienceCreateService{
		Ctx:               &ctx,
		Db:                r.Db,
		Args:              args,
		ProjectExperience: &projectExperience,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	} else {
		return &insightTypes.ProjectExperienceType{
			ProjectExperience: &globalTypes.ProjectExperienceType{
				ProjectExperience: &projectExperience,
			},
		}, nil
	}
}
