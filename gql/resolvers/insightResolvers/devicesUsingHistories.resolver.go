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

func (r *Resolver) DevicesUsingHistories(ctx context.Context, args insightInputs.DevicesUsingHistoriesInput) (*insightTypes.DevicesUsingHistoriesType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeDevices.String(), enums.PermissionActionTypeRead.String()); err != nil {
		return nil, err
	}

	var devicesUsingHistories []*models.DevicesUsingHistory
	devicesUsingHistoryQuery, paginationData := args.ToPaginationDataAndQuery()

	repo := repository.NewDevicesUsingHistoryRepository(&ctx, r.Db.Preload("User"))

	err := repo.List(&devicesUsingHistories, &paginationData, devicesUsingHistoryQuery)
	if err != nil {
		return nil, err
	}

	return &insightTypes.DevicesUsingHistoriesType{
		Collection: r.DevicesUsingHistorySlideToTypes(devicesUsingHistories),
		Metadata: &globalTypes.MetadataType{
			Metadata: &paginationData.Metadata,
		},
	}, nil
}
