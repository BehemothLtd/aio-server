package globalTypes

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)

type AttendanceType struct {
	Attendance *models.Attendance
}

func (at *AttendanceType) ID(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(at.Attendance.Id)
}

func (at *AttendanceType) UserId(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(at.Attendance.UserId)
}

func (at *AttendanceType) User(ctx context.Context) *UserType {
	return &UserType{
		User: &at.Attendance.User,
	}
}

func (at *AttendanceType) CreatedUserId(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(at.Attendance.CreatedUserId)
}

func (at *AttendanceType) CreatedUser(ctx context.Context) *UserType {
	return &UserType{
		User: &at.Attendance.CreatedUser,
	}
}

func (at *AttendanceType) CheckinAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&at.Attendance.CheckinAt)
}

func (at *AttendanceType) CheckoutAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&at.Attendance.CheckoutAt)
}

func (at *AttendanceType) CreatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&at.Attendance.CreatedAt)
}

func (at *AttendanceType) UpdatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&at.Attendance.UpdatedAt)
}
