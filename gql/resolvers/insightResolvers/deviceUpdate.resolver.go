package insightResolvers

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"
)

func (r *Resolver) DeviceUpdate(ctx context.Context, args insightInputs.DeviceUpdateInput) (*globalTypes.DeviceUpdatedType, error) {
	_, err := r.Authorize(ctx, enums.PermissionTargetTypeDevices.String(), enums.PermissionActionTypeWrite.String())
	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	deviceId, err := helpers.GqlIdToInt32(args.Id)
	if err != nil {
		return nil, err
	}

	updatedDevice := models.Device{Id: deviceId}
}
