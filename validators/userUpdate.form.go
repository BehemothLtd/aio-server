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

type UserUpdateForm struct {
	Form
	insightInputs.UserFormInput
	User        *models.User
	UpdatedUser models.User
	Repo        *repository.UserRepository
}

func NewUserUpdateFormValidator(
	input *insightInputs.UserFormInput,
	repo *repository.UserRepository,
	user *models.User,
) UserUpdateForm {
	form := UserUpdateForm{
		Form:          Form{},
		UserFormInput: *input,
		User:          user,
		UpdatedUser:   models.User{},
		Repo:          repo,
	}

	form.assignAttributes(input)
	return form
}

func (form *UserUpdateForm) assignAttributes(input *insightInputs.UserFormInput) {
	form.AddAttributes(
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "about",
			},
			Value: helpers.GetStringOrDefault(input.About),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "slackId",
			},
			Value: helpers.GetStringOrDefault(input.SlackId),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "gender",
			},
			Value: helpers.GetStringOrDefault(input.Gender),
		},
		&TimeAttribute{
			FieldAttribute: FieldAttribute{
				Code: "birthday",
			},
			Value: helpers.GetStringOrDefault(input.Birthday),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "phone",
			},
			Value: helpers.GetStringOrDefault(input.Phone),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "address",
			},
			Value: helpers.GetStringOrDefault(input.Address),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "avatarKey",
			},
			Value: helpers.GetStringOrDefault(input.AvatarKey),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "fullName",
			},
			Value: helpers.GetStringOrDefault(input.FullName),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "password",
			},
			Value: helpers.GetStringOrDefault(input.Password),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "email",
			},
			Value: helpers.GetStringOrDefault(input.Email),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "name",
			},
			Value: helpers.GetStringOrDefault(input.Name),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "state",
			},
			Value: helpers.GetStringOrDefault(input.State),
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "companyLevelId",
			},
			Value: helpers.GetInt32OrDefault(form.CompanyLevelId),
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "lockVersion",
			},
			Value: helpers.GetInt32OrDefault(form.LockVersion),
		},
	)
}

func (form *UserUpdateForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	return form.Repo.UpdateProfile(form.User, form.UpdatedUser)
}

func (form *UserUpdateForm) validate() error {
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

func (form *UserUpdateForm) validateGender() *UserUpdateForm {
	if form.Gender != nil {
		gender := form.FindAttrByCode("gender")

		genderValue := enums.UserGenderType(*form.Gender)

		if genderValue.IsValid() {
			form.UpdatedUser.Gender = &genderValue
		} else {
			gender.AddError("is invalid")
		}
	}

	return form
}

func (form *UserUpdateForm) validateBirthday() *UserUpdateForm {
	if form.Birthday != nil {
		field := form.FindAttrByCode("birthday")

		field.ValidateFormat(constants.YYYYMMDD_DateFormat, constants.HUMAN_YYYYMMDD_DateFormat)

		if field.IsClean() {
			form.UpdatedUser.Birthday = field.Time()
		}
	}

	return form
}

func (form *UserUpdateForm) validateAbout() *UserUpdateForm {
	about := form.FindAttrByCode("about")
	about.ValidateMax(interface{}(int64(constants.MaxLongTextLength)))

	if about.IsClean() {
		form.UpdatedUser.About = form.About
	}

	return form
}

func (form *UserUpdateForm) validateFullName() *UserUpdateForm {
	fullName := form.FindAttrByCode("fullName")

	fullName.ValidateRequired()
	fullName.ValidateMin(interface{}(int64(10)))
	fullName.ValidateMax(interface{}(int64(constants.MaxLongTextLength)))

	if fullName.IsClean() {
		form.UpdatedUser.FullName = *form.FullName
	}

	return form
}

func (form *UserUpdateForm) validateSlackId() *UserUpdateForm {
	slackId := form.FindAttrByCode("slackId")

	slackId.ValidateRequired()
	slackId.ValidateMin(interface{}(int64(11)))
	slackId.ValidateMax(interface{}(int64(constants.MaxStringLength)))

	if slackId.IsClean() {
		form.UpdatedUser.SlackId = form.SlackId
	}

	return form
}

func (form *UserUpdateForm) validatePhone() *UserUpdateForm {
	if form.Phone != nil {
		phone := form.FindAttrByCode("phone")

		phone.ValidateMin(interface{}(int64(10)))
		phone.ValidateMax(interface{}(int64(13)))

		if phone.IsClean() {
			form.UpdatedUser.Phone = form.Phone
		}
	}

	return form
}

func (form *UserUpdateForm) validateAddress() *UserUpdateForm {
	if form.Address != nil {
		address := form.FindAttrByCode("address")

		address.ValidateMin(interface{}(int64(20)))
		address.ValidateMax(interface{}(int64(constants.MaxLongTextLength)))

		if address.IsClean() {
			form.UpdatedUser.Address = form.Address
		}
	}

	return form
}

func (form *UserUpdateForm) validateAvatarKey() *UserUpdateForm {
	avatar := form.FindAttrByCode("avatarKey")

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
