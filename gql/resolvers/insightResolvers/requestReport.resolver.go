package insightResolvers

import (
	"aio-server/enums"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/inputs/insightInputs"
	"context"
)

func (r *Resolver) RequestReport(ctx context.Context, args insightInputs.RequestReportInput) (*[]*globalTypes.RequestReportType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeLeaveDayRequests.String(), enums.PermissionActionTypeRead.String()); err != nil {
		return nil, err
	}

	// var requestReport []*models.RequestReport
	// requestReportQuery := args.ToQuery()

	// repo := repository.NewLeaveDayRequestRepository(&ctx, r.Db)



	return nil, nil
}
