package insightResolvers

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"

	"gorm.io/gorm"
)

func (r *Resolver) ProjectExperience(ctx context.Context, args insightInputs.ProjectExperienceInput) (*globalTypes.ProjectExperienceType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeProjectExperiences.String(), enums.PermissionActionTypeRead.String()); err != nil {
		return nil, err
	}

	if args.Id == "" {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}
	Id, err := helpers.GqlIdToInt32(args.Id)
	if err != nil {
		return nil, err
	}

	projectExperience := models.ProjectExperience{Id: Id}
	repo := repository.NewProjectExperienceRepository(&ctx, r.Db)
	err = repo.Find(&projectExperience)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exceptions.NewRecordNotFoundError()
		}
		return nil, err
	}
	return &globalTypes.ProjectExperienceType{ProjectExperience: &projectExperience}, nil
}
