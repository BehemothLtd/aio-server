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
	deviceUsingHistory := models.DevicesUsingHistory{}

	deviceCreateService := insightServices.DeviceCreateService{
		Ctx:    &ctx,
		Db:     r.Db,
		Args:   args,
		Device: &device,
	}

	deviceUsingHistoryInput := insightInputs.DeviceUsingHistoryCreateInput{
		Input: insightInputs.DeviceUsingHistoryCreateFormInput{
			UserId:   args.Input.UserId,
			DeviceId: &device.Id,
			State:    args.Input.State,
		},
	}

	deviceUsingHistoryCreateService := insightServices.DeviceUsingHistoryCreateService{
		Ctx:                &ctx,
		Db:                 r.Db,
		Args:               deviceUsingHistoryInput,
		DeviceUsingHistory: &deviceUsingHistory,
	}

	if err := deviceCreateService.Execute(); err != nil {
		return nil, err
	}

	if err := deviceUsingHistoryCreateService.Execute(); err != nil {
		return nil, err
	}

	return &insightTypes.DeviceModifiedType{
		Device: &globalTypes.DeviceType{
			Device: &device,
		},
	}, nil

}
