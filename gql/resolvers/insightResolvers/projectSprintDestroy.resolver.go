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

func (r *Resolver) ProjectSprintDestroy(ctx context.Context, args insightInputs.ProjectSprintInput) (*string, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeProjectSprints.String(), enums.PermissionActionTypeDelete.String()); err != nil {
		return nil, err
	}

	projectSprintId, err := helpers.GqlIdToInt32(args.Id)

	if err != nil {
		return nil, exceptions.NewBadRequestError("Invalid ID")
	}

	projectSprint := models.ProjectSprint{
		Id: projectSprintId,
	}

	repo := repository.NewProjectSprintRepository(&ctx, r.Db.Preload("Project").Preload("Issues"))
	if err = repo.Find(&projectSprint); err != nil {
		return nil, exceptions.NewRecordNotFoundError()
	}

	if err := repo.Destroy(&projectSprint); err != nil {
		return nil, exceptions.NewBadRequestError(fmt.Sprintf("Can not delete this project sprint %s", err.Error()))
	} else {
		message := "Deleted"

		return &message, nil
	}
}
