package insightResolvers

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"

	"github.com/graph-gophers/graphql-go"
)

func (r *Resolver) ProjectSprints(ctx context.Context, args struct{ Id graphql.ID }) (*[]*globalTypes.ProjectSprintType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeProjects.String(), enums.PermissionActionTypeRead.String()); err != nil {
		return nil, err
	}

	if args.Id == "" {
		return nil, exceptions.NewBadRequestError("Invalid Project Id")
	}

	projectId, err := helpers.GqlIdToInt32(args.Id)

	if err != nil || projectId == 0 {
		return nil, exceptions.NewBadRequestError("Invalid Project")
	}

	project := models.Project{Id: projectId}
	projectRepo := repository.NewProjectRepository(&ctx, r.Db)

	if err := projectRepo.Find(&project); err != nil {
		return nil, exceptions.NewRecordNotFoundError()
	}

	if project.ProjectType == enums.ProjectTypeKanban {
		return nil, exceptions.NewBadRequestError("Only Scrum Projects have sprints")
	}

	var sprints []models.ProjectSprint
	repo := repository.NewProjectSprintRepository(&ctx, r.Db.Preload("Project"))

	if err := repo.FindAllByProject(project.Id, &sprints); err != nil {
		return nil, exceptions.NewUnprocessableContentError("Error happend", nil)
	}

	result := make([]*globalTypes.ProjectSprintType, len(sprints))

	for i := range sprints {
		result[i] = &globalTypes.ProjectSprintType{ProjectSprint: &sprints[i]}
	}

	return &result, nil
}
