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
				Code: "about",
			},
			Value: about,
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "slackId",
			},
			Value: slackId,
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "gender",
			},
			Value: gender,
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "birthday",
			},
			Value: birthday,
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "phone",
			},
			Value: phone,
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "address",
			},
			Value: address,
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
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

	return form.Repo.Update(form.User, []string{
		"FullName", "Phone", "Birthday", "SlackId",
		"About", "Address", "Gender", "Avatar"},
	)
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
	about.ValidateMax(interface{}(int64(constants.MaxLongTextLength)))

	return form
}

func (form *UserProfileForm) validateSlackId() *UserProfileForm {
	slackId := form.FindAttrByCode("slackId")

	slackId.ValidateRequired()
	slackId.ValidateMin(interface{}(int64(11)))
	slackId.ValidateMax(interface{}(int64(constants.MaxStringLength)))

	return form
}

func (form *UserProfileForm) validatePhone() *UserProfileForm {
	if form.Phone != nil {
		phone := form.FindAttrByCode("phone")

		phone.ValidateMin(interface{}(int64(10)))
		phone.ValidateMax(interface{}(int64(13)))
	}

	return form
}

func (form *UserProfileForm) validateAddress() *UserProfileForm {
	if form.Address != nil {
		address := form.FindAttrByCode("address")

		address.ValidateMin(interface{}(int64(20)))
		address.ValidateMax(interface{}(int64(constants.MaxLongTextLength)))
	}

	return form
}

func (form *UserProfileForm) validateAvatarKey() *UserProfileForm {
	avatar := form.FindAttrByCode("avatarKey")

	if form.AvatarKey != nil {
		if *form.AvatarKey != "" {
			blob := models.AttachmentBlob{Key: *form.AvatarKey}

			repo := repository.NewAttachmentBlobRepository(nil, database.Db)
			if err := repo.Find(&blob); err != nil {

				avatar.AddError("is invalid")
			} else {
				form.User.Avatar = &models.Attachment{
					AttachmentBlob:   blob,
					AttachmentBlobId: blob.Id,
					Name:             "avatar",
				}
			}
		} else {
			avatar.AddError("is invalid")
		}
	}
	return form
}
