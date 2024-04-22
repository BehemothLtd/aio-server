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

func (r *Resolver) SelfProjectExperienceUpdate(ctx context.Context, args insightInputs.ProjectExperienceUpdateInput) (*insightTypes.ProjectExperienceType, error) {
	user, err := r.Authorize(ctx, enums.PermissionTargetTypeProjectExperiences.String(), enums.PermissionActionTypeWrite.String())
	if err != nil {
		return nil, err
	}
	userId := user.Id

	if args.Id == "" {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}

	id, err := helpers.GqlIdToInt32(args.Id)
	if err != nil {
		return nil, err
	}

	projectExperience := models.ProjectExperience{Id: id, UserId: userId}

	service := insightServices.SelfProjectExperienceUpdateService{
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
