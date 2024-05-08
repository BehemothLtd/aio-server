package insightResolvers

import (
	"aio-server/enums"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"context"
)

func (r *Resolver) RequestReport(ctx context.Context, args insightInputs.RequestReportInput) (*[]*globalTypes.RequestReportType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeLeaveDayRequests.String(), enums.PermissionActionTypeRead.String()); err != nil {
		return nil, err
	}

	var requestReports []*models.RequestReport
	requestReportQuery := args.ToQuery()

	repo := repository.NewLeaveDayRequestRepository(&ctx, r.Db)

	err := repo.Report(&requestReports, requestReportQuery)
	if err != nil {
		return nil, err
	}

	return r.RequestReportSlideToTypes(requestReports), nil
}
