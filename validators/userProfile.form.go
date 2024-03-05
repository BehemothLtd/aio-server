package validators

import (
	"aio-server/database"
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
	insightInputs.SelfUpdateFormInput
	User *models.User
	Repo *repository.UserRepository
}

// NewUserProfileFormValidator creates a new UserProfileForm validator.
func NewUserProfileFormValidator(
	input *insightInputs.SelfUpdateFormInput,
	repo *repository.UserRepository,
	user *models.User,
) UserProfileForm {
	form := UserProfileForm{
		Form:                Form{},
		SelfUpdateFormInput: *input,
		User:                user,
		Repo:                repo,
	}
	form.assignAttributes(input)

	return form
}

func (form *UserProfileForm) assignAttributes(input *insightInputs.SelfUpdateFormInput) {
	about := helpers.GetStringOrDefault(input.About)
	slackId := helpers.GetStringOrDefault(input.SlackId)
	gender := helpers.GetStringOrDefault(input.Gender)
	birthday := helpers.GetStringOrDefault(input.Birthday)
	phone := helpers.GetStringOrDefault(input.Phone)
	address := helpers.GetStringOrDefault(input.Address)
	avatarKey := helpers.GetStringOrDefault(input.AvatarKey)

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
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Name: "AvatarKey",
				Code: "avatarKey",
			},
			Value: avatarKey,
		},
	)
}

func (form *UserProfileForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	} else {
		form.User.About = form.About
		form.User.SlackId = form.SlackId
		form.User.Phone = form.Phone
		form.User.Address = form.Address
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
		validateAvatarKey().
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

func (form *UserProfileForm) validateAvatarKey() *UserProfileForm {
	if form.AvatarKey != nil {
		if *form.AvatarKey != "" {
			blob := models.AttachmentBlob{Key: *form.AvatarKey}

			repo := repository.NewAttachmentBlobRepository(nil, database.Db)
			if err := repo.Find(&blob); err != nil {
				avatar := form.FindAttrByCode("avatarKey")
				avatar.AddError("is invalid")
			} else {
				if form.User.Avatar == nil {
					form.User.Avatar = &models.Attachment{
						AttachmentBlob: &blob,
					}
				} else {
					form.User.Avatar.AttachmentBlob = &blob
				}
			}
		} else {
			avatar := form.FindAttrByCode("avatarKey")
			avatar.AddError("is invalid")
		}
	}
	return form
}
