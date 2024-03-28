package globalTypes

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	"github.com/graph-gophers/graphql-go"
)

type DeviceTypeType struct {
	DeviceType *models.DeviceType
}

func (dtt *DeviceTypeType) ID(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(dtt.DeviceType.Id)
}

func (dtt *DeviceTypeType) Name(context.Context) *string {
	return &dtt.DeviceType.Name
}

func (dtt *DeviceTypeType) DevicesCount(context.Context) *int32 {
	return helpers.Int32Pointer(int32(len(dtt.DeviceType.Devices)))
}

func (dtt *DeviceTypeType) CreatedAt(context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&dtt.DeviceType.CreatedAt)
}

func (dtt *DeviceTypeType) UpdatedAt(context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&dtt.DeviceType.UpdatedAt)
}
