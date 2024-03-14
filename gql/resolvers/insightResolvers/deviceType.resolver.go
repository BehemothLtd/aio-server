package insightResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"

	"gorm.io/gorm"
)

func (r *Resolver) DeviceType(ctx context.Context, args insightInputs.DeviceTypeInput) (*globalTypes.DeviceTypeType, error) {
	if args.Id == "" {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}

	deviceTypeId, err := helpers.GqlIdToInt32(args.Id)

	if err != nil {
		return nil, err
	}

	deviceType := models.DeviceType{}
	repo := repository.NewDeviceTypeRepository(&ctx, r.Db)
	err = repo.FindById(&deviceType, deviceTypeId)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exceptions.NewRecordNotFoundError()
		}
		return nil, err
	}

	return &globalTypes.DeviceTypeType{DeviceType: &deviceType}, nil
}
