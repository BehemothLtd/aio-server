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

	service := insightServices.DeviceCreateService{
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
