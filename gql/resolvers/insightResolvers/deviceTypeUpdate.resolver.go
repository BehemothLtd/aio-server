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

func (r *Resolver) DeviceTypeUpdate(ctx context.Context, args insightInputs.DeviceTypeUpdateInput) (*insightTypes.DeviceTypeType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeDevices.String(), enums.PermissionActionTypeWrite.String()); err != nil {
		return nil, err
	}

	if args.Id == "" {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}

	deviceTypeId, err := helpers.GqlIdToInt32(args.Id)

	if err != nil {
		return nil, err
	}

	deviceType := models.DeviceType{Id: deviceTypeId}
	service := insightServices.DeviceTypeUpdateService{
		Ctx:        &ctx,
		Db:         r.Db,
		Args:       args,
		DeviceType: &deviceType,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	} else {
		return &insightTypes.DeviceTypeType{
			DeviceType: &globalTypes.DeviceTypeType{
				DeviceType: &deviceType,
			},
		}, nil
	}
}
