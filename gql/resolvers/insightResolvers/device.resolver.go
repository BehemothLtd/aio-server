package insightResolvers

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"

	"gorm.io/gorm"
)

func (r *Resolver) Device(ctx context.Context, args insightInputs.DeviceInput) (*globalTypes.DeviceType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeDevices.String(), enums.PermissionActionTypeAll.String()); err != nil {
		return nil, err
	}

	if args.ID == "" {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}

	deviceId, err := helpers.GqlIdToInt32(args.ID)

	if err != nil {
		return nil, err
	}

	device := models.Device{Id: deviceId}
	repo := repository.NewDeviceRepository(&ctx, r.Db.Preload("DeviceType"))
	err = repo.Find(&device)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exceptions.NewRecordNotFoundError()
		}
		return nil, err
	}

	return &globalTypes.DeviceType{Device: &device}, nil
}
