package insightResolvers

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"
	"fmt"

	"github.com/graph-gophers/graphql-go"
)

func (r *Resolver) ProjectSprintDestroy(ctx context.Context, args struct {
	ProjectId graphql.ID
	Id        graphql.ID
}) (*string, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeProjectSprints.String(), enums.PermissionActionTypeDelete.String()); err != nil {
		return nil, err
	}

	projectSprintId, err := helpers.GqlIdToInt32(args.Id)

	if err != nil {
		return nil, exceptions.NewBadRequestError("Invalid ID")
	}

	projectId, err := helpers.GqlIdToInt32(args.ProjectId)
	if err != nil || projectId == 0 {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}

	project := models.Project{Id: projectId}
	projectRepo := repository.NewProjectRepository(&ctx, r.Db)
	if err := projectRepo.Find(&project); err != nil {
		return nil, exceptions.NewBadRequestError("Invalid Project")
	}

	projectSprint := models.ProjectSprint{
		Id:        projectSprintId,
		ProjectId: projectId,
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
