package globalTypes

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type DeviceType struct {
	Ctx *context.Context
	Db  *gorm.DB

	Device *models.Device
}

type DeviceUpdatedType struct {
	Device *DeviceType
}

func (dt *DeviceType) ID(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(dt.Device.Id)
}

func (dt *DeviceType) UserId(context.Context) *graphql.ID {
	return helpers.GqlIDP(dt.Device.UserId)
}

func (dt *DeviceType) User(ctx context.Context) *UserType {
	return &UserType{
		User: &dt.Device.User,
	}
}

func (dt *DeviceType) Name(context.Context) *string {
	return &dt.Device.Name
}

func (dt *DeviceType) Code(context.Context) *string {
	return &dt.Device.Code
}

func (dt *DeviceType) Description(context.Context) *string {
	return &dt.Device.Description
}

func (dt *DeviceType) State(context.Context) *string {
	value := dt.Device.State.String()

	return &value
}

func (dt *DeviceType) DeviceTypeId(context.Context) *graphql.ID {
	return helpers.GqlIDP(dt.Device.DeviceTypeId)
}

func (dt *DeviceType) Seller(context.Context) *string {
	return &dt.Device.Seller
}

func (dt *DeviceType) BuyAt(context.Context) *graphql.Time {
	return helpers.GqlTimePointer(dt.Device.BuyAt)
}

func (dt *DeviceType) CreatedAt(context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&dt.Device.CreatedAt)
}

func (dt *DeviceType) UpdatedAt(context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&dt.Device.UpdatedAt)
}

func (dt *DeviceType) DeviceType(ctx context.Context) *DeviceTypeType {
	return &DeviceTypeType{
		DeviceType: &dt.Device.DeviceType,
	}
}

func (dt *DeviceType) DevicesUsingHistories(ctx context.Context) *[]*DevicesUsingHistoryType {
	result := make([]*DevicesUsingHistoryType, len(dt.Device.DevicesUsingHistories))

	for i, devicesUsingHistory := range dt.Device.DevicesUsingHistories {
		result[i] = &DevicesUsingHistoryType{
			DevicesUsingHistory: devicesUsingHistory,
		}
	}

	return &result
}

func (dt *DeviceType) LockVersion(context.Context) *int32 {
	return &dt.Device.LockVersion
}
