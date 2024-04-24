package insightResolvers

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"

	"github.com/graph-gophers/graphql-go"
)

func (r *Resolver) ProjectDestroy(ctx context.Context, args struct{ Id graphql.ID }) (*string, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeProjects.String(), enums.PermissionActionTypeDelete.String()); err != nil {
		return nil, err
	}

	projectId, err := helpers.GqlIdToInt32(args.Id)

	if err != nil {
		return nil, exceptions.NewBadRequestError(err.Error())
	}

	project := models.Project{Id: projectId}
	repo := repository.NewProjectRepository(&ctx, r.Db)

	if err := repo.Delete(&project); err != nil {
		return nil, exceptions.NewBadRequestError(err.Error())
	} else {
		message := "Deleted"
		return &message, nil
	}
}
