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

func (r *Resolver) Devices(ctx context.Context, args insightInputs.DevicesInput) (*insightTypes.DevicesType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeDevices.String(), enums.PermissionActionTypeRead.String()); err != nil {
		return nil, err
	}

	var devices []*models.Device
	deviceQuery, paginationData := args.ToPaginationDataAndQuery()

	repo := repository.NewDeviceRepository(&ctx, r.Db)

	err := repo.List(&devices, &paginationData, deviceQuery)
	if err != nil {
		return nil, err
	}

	return &insightTypes.DevicesType{
		Collection: r.DeviceSlideToTypes(devices),
		Metadata: &globalTypes.MetadataType{
			Metadata: &paginationData.Metadata,
		},
	}, nil
}
