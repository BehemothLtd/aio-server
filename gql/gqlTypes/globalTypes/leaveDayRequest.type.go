package globalTypes

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type LeaveDayRequestType struct {
	Ctx *context.Context
	Db  *gorm.DB

	LeaveDayRequest *models.LeaveDayRequest
}

func (lt *LeaveDayRequestType) ID(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(lt.LeaveDayRequest.Id)
}

func (lt *LeaveDayRequestType) UserId(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(lt.LeaveDayRequest.UserId)
}

func (lt *LeaveDayRequestType) ApproverId(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(lt.LeaveDayRequest.ApproverId)
}

func (lt *LeaveDayRequestType) User(ctx context.Context) *UserType {
	return &UserType{User: &lt.LeaveDayRequest.User}
}

// TODO: Handle approver relation
func (lt *LeaveDayRequestType) Approver(ctx context.Context) *UserType {
	return &UserType{User: &lt.LeaveDayRequest.Approver}
}

func (lt *LeaveDayRequestType) From(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&lt.LeaveDayRequest.From)
}

func (lt *LeaveDayRequestType) To(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&lt.LeaveDayRequest.To)
}

func (lt *LeaveDayRequestType) RequestType(ctx context.Context) *string {
	value := lt.LeaveDayRequest.RequestType.String()

	return &value
}

func (lt *LeaveDayRequestType) RequestState(context.Context) *string {
	value := lt.LeaveDayRequest.RequestState.String()

	return &value
}

// TODO: handle  float poiter
// func (lt *LeaveDayRequestType) TimeOff(context.Context) *float32 {
// 	return &lt.LeaveDayRequest.TimeOff
// }

func (lt *LeaveDayRequestType) Reason(context.Context) *string {
	return &lt.LeaveDayRequest.Reason
}

func (lt *LeaveDayRequestType) CreatedAt(context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&lt.LeaveDayRequest.CreatedAt)
}

func (lt *LeaveDayRequestType) UpdatedAt(context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&lt.LeaveDayRequest.UpdatedAt)
}

func (lt *LeaveDayRequestType) LockVersion(ctx context.Context) int32 {
	return lt.LeaveDayRequest.LockVersion
}
