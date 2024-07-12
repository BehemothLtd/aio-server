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
	devicesUsingHistory := models.DevicesUsingHistory{}

	service := insightServices.DeviceUpdateService{
		Ctx:    &ctx,
		Db:     r.Db,
		Args:   args,
		Device: &device,
	}

	devicesUsingHistoryInput := insightInputs.DevicesUsingHistoryCreateInput{
		Input: insightInputs.DevicesUsingHistoryCreateFormInput{
			UserId:   args.Input.UserId,
			DeviceId: &device.Id,
			State:    args.Input.State,
		},
	}

	devicesUsingHistoryCreateService := insightServices.DevicesUsingHistoryCreateService{
		Ctx:                 &ctx,
		Db:                  r.Db,
		Args:                devicesUsingHistoryInput,
		DevicesUsingHistory: &devicesUsingHistory,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	} else {

		if err := devicesUsingHistoryCreateService.Execute(); err != nil {
			return nil, err
		}
		return &insightTypes.DeviceModifiedType{
			Device: &globalTypes.DeviceType{
				Device: &device,
			},
		}, nil
	}
}
