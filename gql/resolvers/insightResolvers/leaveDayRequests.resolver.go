package insightresolvers

import (
	"aio-server/enums"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"context"
)

func (r *Resolver) LeaveDayRequests(ctx context.Context, args insightInputs.LeaveDayRequestsInput) (*insightTypes.LeaveDayRequestsType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeLeaveDayRequests.String(), enums.PermissionActionTypeRead.String()); err != nil {
		return nil, err
	}

	var leaveDayRequests []*models.LeaveDayRequest
	leaveDayRequestQuery, paginationData := args.ToPaginantionDataAndQuery()

	// repo := repository
}
