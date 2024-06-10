package insightResolvers

import (
	"aio-server/enums"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/services/insightServices"
	"context"
)

func (r *Resolver) DeviceCreate(ctx context.Context, args insightInputs.DeviceCreateInput) (*insightTypes.DeviceModifiedType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeDevices.String(), enums.PermissionActionTypeWrite.String()); err != nil {
		return nil, err
	}

	device := models.Device{}
	devicesUsingHistory := models.DevicesUsingHistory{}

	deviceCreateService := insightServices.DeviceCreateService{
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

	if err := deviceCreateService.Execute(); err != nil {
		return nil, err
	}

	if err := devicesUsingHistoryCreateService.Execute(); err != nil {
		return nil, err
	}

	return &insightTypes.DeviceModifiedType{
		Device: &globalTypes.DeviceType{
			Device: &device,
		},
	}, nil

}
