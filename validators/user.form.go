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

	"golang.org/x/crypto/bcrypt"
)

type UserUpdateForm struct {
	Form
	insightInputs.UserFormInput
	User *models.User
	Repo *repository.UserRepository
}

func NewUserFormValidator(
	input *insightInputs.UserFormInput,
	repo *repository.UserRepository,
	user *models.User,
) UserUpdateForm {
	form := UserUpdateForm{
		Form:          Form{},
		UserFormInput: *input,
		User:          user,
		Repo:          repo,
	}

	form.assignAttributes(input)
	return form
}

func (form *UserUpdateForm) assignAttributes(input *insightInputs.UserFormInput) {
	var companyLevelId int32
	if input.CompanyLevelId != nil {
		companyLevelId, _ = helpers.GqlIdToInt32(*input.CompanyLevelId)
	}

	form.AddAttributes(
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "fullName",
			},
			Value: helpers.GetStringOrDefault(input.FullName),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "email",
			},
			Value: helpers.GetStringOrDefault(input.Email),
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
		&TimeAttribute{
			FieldAttribute: FieldAttribute{
				Code: "birthday",
			},
			Value: helpers.GetStringOrDefault(input.Birthday),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "gender",
			},
			Value: helpers.GetStringOrDefault(input.Gender),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "slackId",
			},
			Value: helpers.GetStringOrDefault(input.SlackId),
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
			Value: helpers.GetInt32OrDefault(&companyLevelId),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "password",
			},
			Value: helpers.GetStringOrDefault(input.Password),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "about",
			},
			Value: helpers.GetStringOrDefault(input.About),
		},

		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "avatarKey",
			},
			Value: helpers.GetStringOrDefault(input.AvatarKey),
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

	if form.User.Id == 0 {
		return form.Repo.Create(form.User)
	}

	return form.Repo.UpdateUser(form.User)
}

func (form *UserUpdateForm) validate() error {
	form.validateFullName().
		validateEmail().
		validatePhone().
		validateAddress().
		validateBirthday().
		validateGender().
		validateSlackId().
		validateState().
		validateCompanyLevelId().
		validatePassword().
		validateAbout().
		validateAvatarKey().
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
			form.User.Gender = &genderValue
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
			form.User.Birthday = field.Time()
		}
	}

	return form
}

func (form *UserUpdateForm) validateAbout() *UserUpdateForm {
	about := form.FindAttrByCode("about")
	about.ValidateMax(interface{}(int64(constants.MaxLongTextLength)))

	if about.IsClean() {
		form.User.About = form.About
	}

	return form
}

func (form *UserUpdateForm) validateFullName() *UserUpdateForm {
	fullName := form.FindAttrByCode("fullName")

	fullName.ValidateRequired()
	fullName.ValidateMin(interface{}(int64(10)))
	fullName.ValidateMax(interface{}(int64(constants.MaxLongTextLength)))

	if fullName.IsClean() {
		form.User.FullName = *form.FullName
	}

	return form
}

func (form *UserUpdateForm) validateSlackId() *UserUpdateForm {
	slackId := form.FindAttrByCode("slackId")

	slackId.ValidateRequired()
	slackId.ValidateMin(interface{}(int64(11)))
	slackId.ValidateMax(interface{}(int64(constants.MaxStringLength)))

	if slackId.IsClean() {
		form.User.SlackId = form.SlackId
	}

	return form
}

func (form *UserUpdateForm) validatePhone() *UserUpdateForm {
	if form.Phone != nil {
		phone := form.FindAttrByCode("phone")

		phone.ValidateMin(interface{}(int64(10)))
		phone.ValidateMax(interface{}(int64(13)))

		if phone.IsClean() {
			form.User.Phone = form.Phone
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
			form.User.Address = form.Address
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

func (form *UserUpdateForm) validateEmail() *UserUpdateForm {
	emailField := form.FindAttrByCode("email")
	emailField.ValidateRequired()

	emailFormat := `\A([^@\s]+)@((?:[-a-z0-9]+\.)+[a-z]{2,})\z`
	emailField.ValidateFormat(emailFormat, "")

	if emailField.IsClean() {
		form.User.Email = *form.Email
	}

	return form
}

func (form *UserUpdateForm) validateState() *UserUpdateForm {
	userState := form.FindAttrByCode("state")

	if form.User.Id == 0 {
		form.User.State = enums.UserStateActive

		return form
	}

	userState.ValidateRequired()

	if userState.IsClean() {
		if userStateEnum, err := enums.ParseUserState(*form.State); err != nil {
			userState.AddError("is invalid")
		} else {
			if userStateEnum == enums.UserStateInactive && !form.User.Inactiveable() {
				userState.AddError("is invalid")
			} else {
				form.User.State = userStateEnum
			}
		}
	}

	return form
}

func (form *UserUpdateForm) validateCompanyLevelId() *UserUpdateForm {
	level := form.FindAttrByCode("companyLevelId")

	if form.CompanyLevelId != nil {
		level.ValidateMin(interface{}(int64(1)))
		level.ValidateMax(interface{}(int64(4)))

		companyLevelId, err := helpers.GqlIdToInt32(*form.CompanyLevelId)
		if err != nil {
			level.AddError("is invalid")
		}

		if level.IsClean() {
			form.User.CompanyLevelId = &companyLevelId
		}
	}

	return form
}

func (form *UserUpdateForm) validatePassword() *UserUpdateForm {
	password := form.FindAttrByCode("password")

	if form.Password != nil {
		password.ValidateMin(interface{}(int64(6)))
		password.ValidateMax(interface{}(int64(20)))

		if encryptPassword, err := bcrypt.GenerateFromPassword([]byte(*form.Password), 10); err != nil {
			password.AddError(err)
		} else {
			form.User.EncryptedPassword = string(encryptPassword)
		}
	}

	return form
}
