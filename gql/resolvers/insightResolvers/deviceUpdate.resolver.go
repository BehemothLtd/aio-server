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

func (r *Resolver) DeviceUpdate(ctx context.Context, args insightInputs.DeviceUpdateInput) (*insightTypes.DeviceModifiedType, error) {
	_, err := r.Authorize(ctx, enums.PermissionTargetTypeDevices.String(), enums.PermissionActionTypeWrite.String())
	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	deviceId, err := helpers.GqlIdToInt32(args.Id)

	if err != nil {
		return nil, err
	}

	device := models.Device{Id: deviceId}
	service := insightServices.DeviceUpdateService{
		Ctx:    &ctx,
		Db:     r.Db,
		Args:   args,
		Device: &device,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	} else {
		return &insightTypes.DeviceModifiedType{
			Device: &globalTypes.DeviceType{
				Device: &device,
			},
		}, nil
	}
}
