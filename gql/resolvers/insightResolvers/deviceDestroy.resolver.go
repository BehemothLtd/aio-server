package insightResolvers

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)

func (r *Resolver) DeviceDestroy(ctx context.Context, args struct{ Id graphql.ID }) (*string, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeDevices.String(), enums.PermissionActionTypeDelete.String()); err != nil {
		return nil, err
	}

	deviceId, err := helpers.GqlIdToInt32(args.Id)

	if err != nil {
		return nil, exceptions.NewBadRequestError(err.Error())
	}

	device := models.Device{Id: deviceId}
	repo := repository.NewDeviceRepository(&ctx, r.Db.Preload("DeviceType"))

	if err := repo.Find(&device); err != nil {
		return nil, exceptions.NewRecordNotFoundError()
	}

	if device.UserId != 0 {
		return nil, exceptions.NewBadRequestError("This device is being used")
	}

	if err := repo.Delete(&device); err != nil {
		return nil, exceptions.NewBadRequestError(err.Error())
	} else {
		message := "Deleted"
		return &message, nil
	}
}
