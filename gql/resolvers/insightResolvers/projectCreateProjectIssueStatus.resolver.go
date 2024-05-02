package insightResolvers

import (
	"aio-server/enums"
	"aio-server/services/insightServices"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)

func (r *Resolver) ProjectCreateProjectIssueStatus(ctx context.Context, args struct {
	ProjectId graphql.ID
	Id        graphql.ID
}) (*string, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeProjects.String(), enums.PermissionActionTypeWrite.String()); err != nil {
		return nil, err
	}

	service := insightServices.ProjectCreateProjectIssueStatusService{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	} else {
		message := "Successfully Created Project Issue Status"
		return &message, nil
	}
}
