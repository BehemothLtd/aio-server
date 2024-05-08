package insightResolvers

import (
	"aio-server/enums"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"context"
)

func (r *Resolver) RequestReport(ctx context.Context, args insightInputs.RequestReportInput) (*insightTypes.RequestReportType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeLeaveDayRequests.String(), enums.PermissionActionTypeRead.String()); err != nil {
		return nil, err
	}

	var requestReport []*models.RequestReport
	requestReportQuery := args.ToQuery()

	repo := repository.NewLeaveDayRequestRepository(&ctx, r.Db)

	err := repo.Report(&requestReport, requestReportQuery)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
