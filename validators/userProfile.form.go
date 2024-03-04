package validators

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/pkg/helpers"
	"aio-server/repository"
)

type UserProfileForm struct {
	Form
	insightInputs.SelfsUpdateFormInput
	User *models.User
	Repo *repository.UserRepository
}

// NewUserProfileFormValidator creates a new UserProfileForm validator.
func NewUserProfileFormValidator(
	input *insightInputs.SelfsUpdateFormInput,
	repo *repository.UserRepository,
	user *models.User,
) UserProfileForm {
	form := UserProfileForm{
		Form:                 Form{},
		SelfsUpdateFormInput: *input,
		User:                 user,
		Repo:                 repo,
	}
	form.assignAttributes(input)

	return form
}

func (form *UserProfileForm) assignAttributes(input *insightInputs.SelfsUpdateFormInput) {
	about := helpers.GetStringOrDefault(input.About)
	slackId := helpers.GetStringOrDefault(input.SlackId)
	gender := helpers.GetStringOrDefault(input.Gender)

	form.AddAttributes(
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Name: "About",
				Code: "about",
			},
			Value: about,
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Name: "SlackId",
				Code: "slackId",
			},
			Value: slackId,
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Name: "Gender",
				Code: "gender",
			},
			Value: gender,
		},
	)

	form.User.About = about
	form.User.SlackId = &slackId
}

func (form *UserProfileForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	return form.Repo.Update(form.User)
}

// validate validates the snippet form.
func (form *UserProfileForm) validate() error {
	form.validateAbout().validateSlackId().validateGender().summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}

	return nil
}

func (form *UserProfileForm) assignGender(gender string) {
	if genderValue, err := enums.ParseUserGenderType(gender); err != nil {
		form.User.Gender = nil
	} else {
		form.User.Gender = &genderValue
	}
}

func (form *UserProfileForm) validateGender() *UserProfileForm {
	if form.Gender != nil {
		gender := form.FindAttrByCode("gender")

		genderValue := enums.UserGenderType(*form.Gender)

		if genderValue.IsValid() {
			form.assignGender(*form.Gender)
		} else {
			gender.AddError("is invalid")
		}
	}

	return form
}

func (form *UserProfileForm) validateAbout() *UserProfileForm {
	about := form.FindAttrByCode("about")
	maxTitleLength := int64(constants.MaxLongTextLength)

	about.ValidateLimit(nil, &maxTitleLength)

	return form
}

func (form *UserProfileForm) validateSlackId() *UserProfileForm {
	slackId := form.FindAttrByCode("slackId")
	minLength := 11
	maxLength := int64(constants.MaxStringLength)

	slackId.ValidateRequired()
	slackId.ValidateLimit(&minLength, &maxLength)

	return form
}