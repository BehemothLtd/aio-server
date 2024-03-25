package insightResolvers

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/services/insightServices"
	"context"
)

func (r *Resolver) IssueStatusUpdate(ctx context.Context, args insightInputs.IssueStatusInput) (*insightTypes.IssueStatusUpdatedType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeIssueStatuses.String(), enums.PermissionActionTypeWrite.String()); err != nil {
		return nil, err
	}

	if args.Id == "" {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}

	issueStatusId, err := helpers.GqlIdToInt32(args.Id)

	if err != nil {
		return nil, err
	}

	issueStatus := models.IssueStatus{Id: issueStatusId}
	service := insightServices.IssueStatusUpdateService{
		Ctx:         &ctx,
		Db:          r.Db,
		Args:        args,
		IssueStatus: &issueStatus,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	} else {
		return &insightTypes.IssueStatusUpdatedType{
			IssueStatus: &globalTypes.IssueStatusType{
				IssueStatus: &issueStatus,
			},
		}, nil
	}
}
