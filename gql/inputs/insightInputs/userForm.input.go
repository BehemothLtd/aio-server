package insightInputs

type UserFormInput struct {
	FullName       *string
	Password       *string
	Name           *string
	Email          *string
	Phone          *string
	Birthday       *string
	SlackId        *string
	About          *string
	AvatarKey      *string
	Address        *string
	Gender         *string
	State          *string
	LockVersion    *int32
	Active         *bool
	CompanyLevelId *int32
}
