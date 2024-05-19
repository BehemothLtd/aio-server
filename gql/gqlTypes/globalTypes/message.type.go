package globalTypes

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type MessageType struct {
	Ctx *context.Context
	Db  *gorm.DB

	Message *models.Message
}

func (mt *MessageType) ID(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(mt.Message.Id)
}

func (mt *MessageType) LeaveDayRequestId(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(mt.Message.LeaveDayRequestId)
}

func (mt *MessageType) Content(ctx context.Context) *string {
	return mt.Message.Content
}

func (mt *MessageType) Timestamp(ctx context.Context) *string {
	return mt.Message.Timestamp
}

func (mt *MessageType) CreatedAt(context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&mt.Message.CreatedAt)
}

func (mt *MessageType) UpdatedAt(context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&mt.Message.UpdatedAt)
}
