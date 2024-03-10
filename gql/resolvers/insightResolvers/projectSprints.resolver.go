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

func (r *Resolver) ProjectSprint(ctx context.Context, args insightInputs.ProjectSprintInput) (*globalTypes.ProjectSprintType, error) {

	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeProjects.String(), enums.PermissionActionTypeWrite.String()); err != nil {
		return nil, err
	}

	if args.Id == "" {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}

	projectSprintID, err := helpers.GqlIdToInt32(args.Id)

	if err != nil {
		return nil, err
	}

	projectSprint := models.ProjectSprint{}
	repo := repository.NewProjectSprintRepository(&ctx, r.Db)
	err = repo.FindById(&projectSprint, projectSprintID)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exceptions.NewRecordNotFoundError()
		}
		return nil, err
	}

	return &globalTypes.ProjectSprintType{ProjectSprint: &projectSprint}, nil
}
