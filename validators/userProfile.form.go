package validators

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"time"
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
	birthday := helpers.GetStringOrDefault(input.Birthday)
	phone := helpers.GetStringOrDefault(input.Phone)
	address := helpers.GetStringOrDefault(input.Address)

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
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Name: "Birthday",
				Code: "birthday",
			},
			Value: birthday,
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Name: "Phone",
				Code: "phone",
			},
			Value: phone,
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Name: "Address",
				Code: "address",
			},
			Value: address,
		},
	)
}

func (form *UserProfileForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	} else {
		if form.About != nil {
			form.User.About = form.About
		}
		if form.SlackId != nil {
			form.User.SlackId = form.SlackId
		}
		form.User.Phone = form.Phone
		if form.Address != nil {
			form.User.Address = form.Address
		}
	}

	return form.Repo.Update(form.User)
}

// validate validates the snippet form.
func (form *UserProfileForm) validate() error {
	form.validateAbout().
		validateSlackId().
		validateGender().
		validateBirthday().
		validatePhone().
		validateAddress().
		summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}

	return nil
}

func (form *UserProfileForm) validateGender() *UserProfileForm {
	if form.Gender != nil {
		gender := form.FindAttrByCode("gender")

		genderValue := enums.UserGenderType(*form.Gender)

		if genderValue.IsValid() {
			form.User.Gender = &genderValue
		} else {
			gender.AddError("is invalid")
		}
	}

	return form
}

func (form *UserProfileForm) validateBirthday() *UserProfileForm {
	if form.Birthday != nil {
		if birthday, error := time.Parse(time.DateOnly, *form.Birthday); error != nil {
			birthdayField := form.FindAttrByCode("birthday")
			birthdayField.AddError("is invalid date")
		} else {
			form.User.Birthday = birthday
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

func (form *UserProfileForm) validatePhone() *UserProfileForm {
	if form.Phone != nil {
		phone := form.FindAttrByCode("phone")
		minLength := 10
		maxLength := int64(13)

		phone.ValidateLimit(&minLength, &maxLength)
	}

	return form
}

func (form *UserProfileForm) validateAddress() *UserProfileForm {
	if form.Address != nil {
		address := form.FindAttrByCode("address")
		minLength := 20
		maxLength := int64(constants.MaxLongTextLength)

		address.ValidateLimit(&minLength, &maxLength)
	}

	return form
}
