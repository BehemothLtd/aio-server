package insightInputs

import "github.com/graph-gophers/graphql-go"

type UserFormInput struct {
	FullName       *string
	Email          *string
	Phone          *string
	Address        *string
	Birthday       *string
	Gender         *string
	SlackId        *string
	State          *string
	CompanyLevelId *graphql.ID
	Password       *string
	About          *string
	AvatarKey      *string
	LockVersion    *int32
}
