package insightResolvers

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)

func (r *Resolver) ProjectSprintActive(ctx context.Context, args struct {
	ProjectId graphql.ID
	Id        graphql.ID
}) (*string, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeProjectSprints.String(), enums.PermissionActionTypeWrite.String()); err != nil {
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

	repo := repository.NewProjectSprintRepository(&ctx, r.Db)
	if err = repo.Find(&projectSprint); err != nil {
		return nil, exceptions.NewRecordNotFoundError()
	}

	if err := projectRepo.ActiveSprint(&project, projectSprint); err != nil {
		return nil, exceptions.NewUnprocessableContentError(err.Error(), nil)
	}

	message := "Active Sprint Successfully"

	return &message, nil
}
