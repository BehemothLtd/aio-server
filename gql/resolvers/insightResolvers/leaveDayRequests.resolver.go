package insightResolvers

import (
	"aio-server/enums"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"context"
)

func (r *Resolver) LeaveDayRequests(ctx context.Context, args insightInputs.LeaveDayRequestsInput) (*insightTypes.LeaveDayRequestsType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeLeaveDayRequests.String(), enums.PermissionActionTypeRead.String()); err != nil {
		return nil, err
	}

	var leaveDayRequests []*models.LeaveDayRequest
	leaveDayRequestQuery, paginationData := args.ToPaginationDataAndQuery()

	repo := repository.NewLeaveDayRequestRepository(&ctx, r.Db)

	err := repo.List(&leaveDayRequests, &paginationData, leaveDayRequestQuery)
	if err != nil {
		return nil, err
	}

	return &insightTypes.LeaveDayRequestsType{
		Collection: r.LeaveDayRequestSliceToTypes(leaveDayRequests),
		Metadata: &globalTypes.MetadataType{
			Metadata: &paginationData.Metadata,
		},
	}, nil
}
