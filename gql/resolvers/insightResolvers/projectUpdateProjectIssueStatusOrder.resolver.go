package insightResolvers

import (
	"aio-server/enums"
	"aio-server/services/insightServices"
	"context"

	"github.com/graph-gophers/graphql-go"
)

func (r *Resolver) ProjectUpdateProjectIssueStatusOrder(ctx context.Context, args struct {
	Id    graphql.ID
	Input []int32
}) (*string, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeProjects.String(), enums.PermissionActionTypeWrite.String()); err != nil {
		return nil, err
	}

	service := insightServices.ProjectUpdateProjectIssueStatusOrderService{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	}

	message := "Success"
	return &message, nil
}
