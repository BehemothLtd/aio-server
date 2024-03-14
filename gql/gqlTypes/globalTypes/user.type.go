package globalTypes

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

// UserType resolves self information.
type UserType struct {
	Ctx *context.Context
	Db  *gorm.DB

	User *models.User
}

type UserUpdatedType struct {
	User *UserType
}

// ID returns the ID of the user.
func (ut *UserType) ID(context.Context) *graphql.ID {
	return helpers.GqlIDP(ut.User.Id)
}

// Email returns the email of the user.
func (ut *UserType) Email(context.Context) *string {
	return &ut.User.Email
}

// FullName returns the full name of the user.
func (ut *UserType) FullName(context.Context) *string {
	return &ut.User.FullName
}

// Name returns the name of the user.
func (ut *UserType) Name(context.Context) *string {
	return &ut.User.Name
}

// About returns the about of the user.
func (ut *UserType) About(context.Context) *string {
	return ut.User.About
}

// AvatarURL returns the AvatarURL of the user.
func (ut *UserType) AvatarUrl(context.Context) *string {
	if ut.User.Avatar != nil {
		url := ut.User.Avatar.Url()
		return &url
	}
	return nil
}

// CreatedAt returns the CreatedAt of the user.
func (ut *UserType) CreatedAt(context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&ut.User.CreatedAt)
}

// CompanyLevelId returns the CompanyLevelId of the user.
func (ut *UserType) CompanyLevelId(context.Context) *graphql.ID {
	if ut.User.CompanyLevelId != nil {
		return helpers.GqlIDP(*ut.User.CompanyLevelId)
	} else {
		return nil
	}
}

// Address returns the Address of the user.
func (ut *UserType) Address(context.Context) *string {
	return ut.User.Address
}

// Phone returns the Phone of the user.
func (ut *UserType) Phone(context.Context) *string {
	return ut.User.Phone
}

// TimingActivedAt returns the TimingActivedAt of the user.
func (ut *UserType) TimingActiveAt(context.Context) *graphql.Time {
	timing := ut.User.Timing

	if timing != nil && timing.ActiveAt != "" {
		return helpers.RubyTimeStringToGqlTime(timing.ActiveAt)
	}

	return nil
}

// timingDeactiveAt returns the timingDeactiveAt of the user.
func (ut *UserType) TimingDeactiveAt(context.Context) *graphql.Time {
	timing := ut.User.Timing

	if timing != nil && timing.InactiveAt != "" {
		return helpers.RubyTimeStringToGqlTime(timing.InactiveAt)
	}
	return nil
}

// Gender returns the Gender of the user.
func (ut *UserType) Gender(context.Context) *string {
	if ut.User.Gender != nil {
		gender := ut.User.Gender.String()
		return &gender
	}

	return nil
}

// Birthday returns the Birthday of the user.
func (ut *UserType) Birthday(context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&ut.User.Birthday)
}

// State returns the State of the user.
func (ut *UserType) State(context.Context) string {
	return ut.User.State.String()
}

// SlackId returns the SlackId of the user.
func (ut *UserType) SlackId(context.Context) *string {
	return ut.User.SlackId
}

// LockVersion returns the lock version of the user.
func (sir *UserType) LockVersion(context.Context) int32 {
	return sir.User.LockVersion
}
