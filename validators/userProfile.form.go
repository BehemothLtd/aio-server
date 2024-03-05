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
		Form: Form{
			SkipAttributes: []string{"AvatarKey"},
			UpdateMap:      map[string]interface{}{},
		},
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
				Code: "About",
			},
			Value: about,
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Name: "Slack Id",
				Code: "SlackId",
			},
			Value: slackId,
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Name: "Gender",
				Code: "Gender",
			},
			Value: gender,
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Name: "Birthday",
				Code: "Birthday",
			},
			Value: birthday,
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Name: "Phone",
				Code: "Phone",
			},
			Value: phone,
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Name: "Address",
				Code: "Address",
			},
			Value: address,
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Name: "AvatarKey",
				Code: "AvatarKey",
			},
			Value: avatarKey,
		},
	)
}

func (form *UserProfileForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	return form.Repo.UpdateFields(form.User, form.UpdateMap)
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
		gender := form.FindAttrByCode("Gender")

		genderValue := enums.UserGenderType(*form.Gender)

		if !genderValue.IsValid() {
			gender.AddError("is invalid")
		}
	}

	return form
}

func (form *UserProfileForm) validateBirthday() *UserProfileForm {
	if form.Birthday != nil {
		if _, error := time.Parse(time.DateOnly, *form.Birthday); error != nil {
			birthdayField := form.FindAttrByCode("Birthday")
			birthdayField.AddError("is invalid date")
		}
	}

	return form
}

func (form *UserProfileForm) validateAbout() *UserProfileForm {
	about := form.FindAttrByCode("About")
	maxTitleLength := int64(constants.MaxLongTextLength)

	about.ValidateLimit(nil, &maxTitleLength)

	return form
}

func (form *UserProfileForm) validateSlackId() *UserProfileForm {
	slackId := form.FindAttrByCode("SlackId")
	minLength := 11
	maxLength := int64(constants.MaxStringLength)

	slackId.ValidateRequired()
	slackId.ValidateLimit(&minLength, &maxLength)

	return form
}

func (form *UserProfileForm) validatePhone() *UserProfileForm {
	if form.Phone != nil {
		phone := form.FindAttrByCode("Phone")
		minLength := 10
		maxLength := int64(13)

		phone.ValidateLimit(&minLength, &maxLength)
	}

	return form
}

func (form *UserProfileForm) validateAddress() *UserProfileForm {
	if form.Address != nil {
		address := form.FindAttrByCode("Address")
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
				avatar := form.FindAttrByCode("AvatarKey")
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
			avatar := form.FindAttrByCode("AvatarKey")
			avatar.AddError("is invalid")
		}
	}
	return form
}
