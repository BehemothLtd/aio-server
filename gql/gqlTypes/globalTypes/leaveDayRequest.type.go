package globalTypes

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/pkg/utilities"
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
	defaultIdValue := helpers.GetInt32OrDefault(lt.LeaveDayRequest.ApproverId)

	return helpers.GqlIDP(defaultIdValue)
}

func (lt *LeaveDayRequestType) User(ctx context.Context) *UserType {
	return &UserType{User: &lt.LeaveDayRequest.User}
}

func (lt *LeaveDayRequestType) Approver(ctx context.Context) *UserType {
	if lt.LeaveDayRequest.ApproverId != nil {
		return &UserType{User: lt.LeaveDayRequest.Approver}
	} else {
		return nil
	}
}

func (lt *LeaveDayRequestType) Message(ctx context.Context) *MessageType {
	if lt.LeaveDayRequest.Message != nil {
		return &MessageType{Message: lt.LeaveDayRequest.Message}
	} else {
		return nil
	}
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

func (lt *LeaveDayRequestType) TimeOff(ctx context.Context) *float64 {
	return helpers.Float64Pointer(lt.LeaveDayRequest.TimeOff)
}

func (lt *LeaveDayRequestType) Reason(context.Context) *string {
	return lt.LeaveDayRequest.Reason
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

func (lt *LeaveDayRequestType) RequestTypeHumanize(ctx context.Context) *string {
	value := utilities.SnakeCaseToHumanize(lt.LeaveDayRequest.RequestType.String())

	return &value
}

func (lt *LeaveDayRequestType) RequestStateHumanize(ctx context.Context) *string {
	value := utilities.SnakeCaseToHumanize(lt.LeaveDayRequest.RequestState.String())

	return &value
}
