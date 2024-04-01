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

func (r *Resolver) IssueStatusCreate(ctx context.Context, args insightInputs.IssueStatusCreateInput) (*insightTypes.IssueStatusCreatedType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeIssueStatuses.String(), enums.PermissionActionTypeWrite.String()); err != nil {
		return nil, err
	}

	IssueStatus := models.IssueStatus{}
	service := insightServices.IssueStatusCreateService{
		Ctx:         &ctx,
		Db:          r.Db,
		Args:        args,
		IssueStatus: &IssueStatus,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	} else {
		return &insightTypes.IssueStatusCreatedType{
			IssueStatus: &globalTypes.IssueStatusType{
				IssueStatus: &IssueStatus,
			},
		}, nil
	}
}
