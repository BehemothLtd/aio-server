package insightResolvers

import (
	"aio-server/enums"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/services/insightServices"
	"context"
)

func (r *Resolver) ProjectUpdateIssue(ctx context.Context, args insightInputs.ProjectUpdateIssueInput) (*insightTypes.IssueCreatedType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeProjectIssues.String(), enums.PermissionActionTypeWrite.String()); err != nil {
		return nil, err
	}
	issue := models.Issue{}

	service := insightServices.ProjectUpdateIssueService{
		Ctx:   &ctx,
		Db:    r.Db,
		Args:  args,
		Issue: &issue,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	} else {
		return &insightTypes.IssueCreatedType{
			Issue: &globalTypes.IssueType{
				Issue: &issue,
			},
		}, nil
	}

}
