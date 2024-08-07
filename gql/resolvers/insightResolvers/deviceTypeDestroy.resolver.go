package insightResolvers

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"
	"fmt"
)

func (r *Resolver) DeviceTypeDestroy(ctx context.Context, args insightInputs.DeviceTypeInput) (*string, error) {
	_, err := r.Authorize(ctx, string(enums.PermissionTargetTypeDevices), string(enums.PermissionActionTypeDelete))

	if err != nil {
		return nil, err
	}

	deviceTypeId, err := helpers.GqlIdToInt32(args.Id)

	if err != nil {
		return nil, exceptions.NewBadRequestError("Invalid ID")
	}

	deviceType := models.DeviceType{Id: deviceTypeId}

	repo := repository.NewDeviceTypeRepository(&ctx, r.Db.Preload("Devices"))

	if err := repo.Find(&deviceType); err != nil {
		return nil, exceptions.NewRecordNotFoundError()
	}

	devicesCount := len(deviceType.Devices)

	if devicesCount > 0 {
		return nil, exceptions.NewBadRequestError("This device type is being used")
	}

	if err := repo.Destroy(&deviceType); err != nil {
		return nil, exceptions.NewBadRequestError(fmt.Sprintf("Cant delete this device type %s", err.Error()))
	} else {
		message := "Deleted"

		return &message, nil
	}
}
