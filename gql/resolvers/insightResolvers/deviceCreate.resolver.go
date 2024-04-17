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

func (r *Resolver) DeviceCreate(ctx context.Context, args insightInputs.DeviceCreateInput) (*insightTypes.DeviceCreatedType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeDevices.String(), enums.PermissionActionTypeWrite.String()); err != nil {
		return nil, err
	}

	Device := models.Device{}
	service := insightServices.DeviceService{
		Ctx:    &ctx,
		Db:     r.Db,
		Args:   args,
		Device: &Device,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	} else {
		return &insightTypes.DeviceCreatedType{
			Device: &globalTypes.DeviceType{
				Device: &Device,
			},
		}, nil
	}
}
