package insightResolvers

import (
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/inputs/insightInputs"
	"context"
)

func (r *Resolver) RequestReport(ctx context.Context, args insightInputs.LeaveDayRequestsInput) (*[]*globalTypes.RequestReportType, error) {
	return nil, nil
}
