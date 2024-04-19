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
	"strings"
)

type UserProfileForm struct {
	Form
	insightInputs.SelfUpdateFormInput
	User    *models.User
	updates map[string]interface{}
	Repo    *repository.UserRepository
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
		updates:             map[string]interface{}{},
		Repo:                repo,
	}
	form.assignAttributes(input)

	return form
}

func (form *UserProfileForm) assignAttributes(input *insightInputs.SelfUpdateFormInput) {
	form.AddAttributes(
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "About",
			},
			Value: helpers.GetStringOrDefault(input.About),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "SlackId",
			},
			Value: helpers.GetStringOrDefault(input.SlackId),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "Gender",
			},
			Value: helpers.GetStringOrDefault(input.Gender),
		},
		&TimeAttribute{
			FieldAttribute: FieldAttribute{
				Code: "Birthday",
			},
			Value: helpers.GetStringOrDefault(input.Birthday),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "Phone",
			},
			Value: helpers.GetStringOrDefault(input.Phone),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "Address",
			},
			Value: helpers.GetStringOrDefault(input.Address),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "AvatarKey",
			},
			Value: helpers.GetStringOrDefault(input.AvatarKey),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "FullName",
			},
			Value: helpers.GetStringOrDefault(input.FullName),
		},
	)
}

func (form *UserProfileForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	return form.Repo.UpdateProfile(form.User, form.updates)
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
		validateFullName().
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
			form.updates["Gender"] = &genderValue
		} else {
			gender.AddError("is invalid")
		}
	}

	return form
}

func (form *UserProfileForm) validateBirthday() *UserProfileForm {
	key := "Birthday"

	if form.Birthday != nil {
		field := form.FindAttrByCode(key)

		field.ValidateFormat(constants.YYYYMMDD_DateFormat, constants.HUMAN_YYYYMMDD_DateFormat)

		if field.IsClean() {
			form.updates["Birthday"] = field.Time()
		}
	}

	return form
}

func (form *UserProfileForm) validateAbout() *UserProfileForm {
	key := "About"
	about := form.FindAttrByCode(key)
	about.ValidateMax(interface{}(int64(constants.MaxLongTextLength)))

	if about.IsClean() {
		form.updates[key] = form.About
	}

	return form
}

func (form *UserProfileForm) validateFullName() *UserProfileForm {
	key := "FullName"
	fullName := form.FindAttrByCode(key)

	fullName.ValidateRequired()
	fullName.ValidateMin(interface{}(int64(10)))
	fullName.ValidateMax(interface{}(int64(constants.MaxLongTextLength)))

	if fullName.IsClean() {
		form.updates[key] = *form.FullName
	}

	return form
}

func (form *UserProfileForm) validateSlackId() *UserProfileForm {
	key := "SlackId"
	slackId := form.FindAttrByCode(key)

	slackId.ValidateRequired()
	slackId.ValidateMin(interface{}(int64(11)))
	slackId.ValidateMax(interface{}(int64(constants.MaxStringLength)))

	if slackId.IsClean() {
		form.updates[key] = form.SlackId
	}

	return form
}

func (form *UserProfileForm) validatePhone() *UserProfileForm {
	key := "Phone"

	if form.Phone != nil {
		phone := form.FindAttrByCode(key)

		phone.ValidateMin(interface{}(int64(10)))
		phone.ValidateMax(interface{}(int64(13)))

		if phone.IsClean() {
			form.updates[key] = form.Phone
		}
	}

	return form
}

func (form *UserProfileForm) validateAddress() *UserProfileForm {
	key := "Address"

	if form.Address != nil {
		address := form.FindAttrByCode(key)

		address.ValidateMin(interface{}(int64(20)))
		address.ValidateMax(interface{}(int64(constants.MaxLongTextLength)))

		if address.IsClean() {
			form.updates[key] = form.Address
		}
	}

	return form
}

func (form *UserProfileForm) validateAvatarKey() *UserProfileForm {
	avatar := form.FindAttrByCode("AvatarKey")

	if form.AvatarKey != nil && strings.TrimSpace(*form.AvatarKey) != "" {
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
	}
	return form
}
