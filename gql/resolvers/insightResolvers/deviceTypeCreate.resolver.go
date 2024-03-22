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

func (r *Resolver) DeviceTypeCreate(ctx context.Context, args insightInputs.DeviceTypeCreateInput) (*insightTypes.DeviceTypeCreatedType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeDevices.String(), enums.PermissionActionTypeWrite.String()); err != nil {
		return nil, err
	}

	deviceType := models.DeviceType{}
	service := insightServices.DeviceTypeCreateService{
		Ctx:        &ctx,
		Db:         r.Db,
		Args:       args,
		DeviceType: &deviceType,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	} else {
		return &insightTypes.DeviceTypeCreatedType{
			DeviceType: &globalTypes.DeviceTypeType{
				DeviceType: &deviceType,
			},
		}, nil
	}
}
