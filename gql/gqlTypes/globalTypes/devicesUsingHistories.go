package globalTypes

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type DevicesUsingHistoryType struct {
	Ctx *context.Context
	Db  *gorm.DB

	DevicesUsingHistory *models.DevicesUsingHistory
}

func (duht *DevicesUsingHistoryType) ID(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(duht.DevicesUsingHistory.Id)
}

func (duht *DevicesUsingHistoryType) State(context.Context) *string {
	value := duht.DevicesUsingHistory.State.String()

	return &value
}

func (duht *DevicesUsingHistoryType) User(ctx context.Context) *UserType {
	return &UserType{
		User: &duht.DevicesUsingHistory.User,
	}
}

func (duht *DevicesUsingHistoryType) CreatedAt(context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&duht.DevicesUsingHistory.CreatedAt)
}

func (duht *DevicesUsingHistoryType) UpdatedAt(context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&duht.DevicesUsingHistory.UpdatedAt)
}
