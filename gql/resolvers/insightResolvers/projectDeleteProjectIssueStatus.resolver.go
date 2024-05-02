package insightResolvers

import (
	"aio-server/enums"
	"aio-server/services/insightServices"
	"context"

	"github.com/graph-gophers/graphql-go"
)

func (r *Resolver) ProjectDeleteProjectIssueStatus(ctx context.Context, args struct {
	ProjectId graphql.ID
	Id        graphql.ID
}) (*string, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeProjectIssueStatuses.String(), enums.PermissionActionTypeDelete.String()); err != nil {
		return nil, err
	}

	service := insightServices.ProjectDeleteProjectIssueStatusService{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	} else {
		message := "Successfully deleted Project Issue Status"
		return &message, nil
	}
}
