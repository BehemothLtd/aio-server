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

func (r *Resolver) DeviceTypes(ctx context.Context, args insightInputs.DeviceTypesInput) (*insightTypes.DeviceTypesType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeDevices.String(), enums.PermissionActionTypeRead.String()); err != nil {
		return nil, err
	}

	var deviceTypes []*models.DeviceType
	paginationData := args.ToPaginationData()

	repo := repository.NewDeviceTypeRepository(&ctx, r.Db)

	err := repo.List(&deviceTypes, &paginationData)
	if err != nil {
		return nil, err
	}

	return &insightTypes.DeviceTypesType{
		Collection: r.DeviceTypesSlideToType(deviceTypes),
		Metadata: &globalTypes.MetadataType{
			Metadata: &paginationData.Metadata,
		},
	}, nil
}
