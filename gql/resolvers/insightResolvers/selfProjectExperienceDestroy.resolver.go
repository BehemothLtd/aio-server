package insightResolvers

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"
	"fmt"
)

func (r *Resolver) SelfProjectExperienceDestroy(ctx context.Context, args insightInputs.ProjectExperienceInput) (*string, error) {
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
	repo := repository.NewProjectExperienceRepository(&ctx, r.Db)

	if err := repo.Find(&projectExperience); err != nil {
		return nil, exceptions.NewRecordNotFoundError()
	}

	if err = repo.Destroy(&projectExperience); err != nil {
		return nil, exceptions.NewBadRequestError(fmt.Sprintf("Can not delete this project experience %s", err.Error()))
	} else {
		message := "Deleted"
		return &message, nil
	}
}
