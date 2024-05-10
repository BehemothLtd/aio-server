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
	User         *models.User
	UpdatedAttrs map[string]interface{}
	Repo         *repository.UserRepository
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
		UpdatedAttrs:  map[string]interface{}{},
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
				Code: "FullName",
			},
			Value: helpers.GetStringOrDefault(input.FullName),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "Email",
			},
			Value: helpers.GetStringOrDefault(input.Email),
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
		&TimeAttribute{
			FieldAttribute: FieldAttribute{
				Code: "Birthday",
			},
			Value: helpers.GetStringOrDefault(input.Birthday),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "Gender",
			},
			Value: helpers.GetStringOrDefault(input.Gender),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "SlackId",
			},
			Value: helpers.GetStringOrDefault(input.SlackId),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "State",
			},
			Value: helpers.GetStringOrDefault(input.State),
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "CompanyLevelId",
			},
			Value: helpers.GetInt32OrDefault(&companyLevelId),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "Password",
			},
			Value: helpers.GetStringOrDefault(input.Password),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "About",
			},
			Value: helpers.GetStringOrDefault(input.About),
		},

		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "AvatarKey",
			},
			Value: helpers.GetStringOrDefault(input.AvatarKey),
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "LockVersion",
			},
			Value: helpers.GetInt32OrDefault(form.LockVersion),
		},
	)
}

func (form *UserUpdateForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	return form.Repo.UpdateUser(form.User, form.UpdatedAttrs)
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
		key := "Gender"
		gender := form.FindAttrByCode(key)

		genderValue := enums.UserGenderType(*form.Gender)

		if genderValue.IsValid() {
			form.UpdatedAttrs[key] = &genderValue
		} else {
			gender.AddError("is invalid")
		}
	}

	return form
}

func (form *UserUpdateForm) validateBirthday() *UserUpdateForm {
	if form.Birthday != nil {
		key := "Birthday"
		field := form.FindAttrByCode(key)

		field.ValidateFormat(constants.YYYYMMDD_DateFormat, constants.HUMAN_YYYYMMDD_DateFormat)

		if field.IsClean() {
			form.UpdatedAttrs[key] = field.Time()
		}
	}

	return form
}

func (form *UserUpdateForm) validateAbout() *UserUpdateForm {
	key := "About"
	about := form.FindAttrByCode(key)
	about.ValidateMax(interface{}(int64(constants.MaxLongTextLength)))

	if about.IsClean() {
		form.UpdatedAttrs[key] = form.About
	}

	return form
}

func (form *UserUpdateForm) validateFullName() *UserUpdateForm {
	key := "FullName"
	fullName := form.FindAttrByCode(key)

	fullName.ValidateRequired()
	fullName.ValidateMin(interface{}(int64(10)))
	fullName.ValidateMax(interface{}(int64(constants.MaxLongTextLength)))

	if fullName.IsClean() {
		form.UpdatedAttrs[key] = *form.FullName
	}

	return form
}

func (form *UserUpdateForm) validateSlackId() *UserUpdateForm {
	key := "SlackId"
	slackId := form.FindAttrByCode(key)

	slackId.ValidateRequired()
	slackId.ValidateMin(interface{}(int64(11)))
	slackId.ValidateMax(interface{}(int64(constants.MaxStringLength)))

	if slackId.IsClean() {
		form.UpdatedAttrs[key] = form.SlackId
	}

	return form
}

func (form *UserUpdateForm) validatePhone() *UserUpdateForm {
	if form.Phone != nil {
		key := "Phone"
		phone := form.FindAttrByCode(key)

		phone.ValidateMin(interface{}(int64(10)))
		phone.ValidateMax(interface{}(int64(13)))

		if phone.IsClean() {
			form.UpdatedAttrs[key] = form.Phone
		}
	}

	return form
}

func (form *UserUpdateForm) validateAddress() *UserUpdateForm {
	if form.Address != nil {
		key := "Address"
		address := form.FindAttrByCode(key)

		address.ValidateMin(interface{}(int64(20)))
		address.ValidateMax(interface{}(int64(constants.MaxLongTextLength)))

		if address.IsClean() {
			form.UpdatedAttrs[key] = form.Address
		}
	}

	return form
}

func (form *UserUpdateForm) validateAvatarKey() *UserUpdateForm {
	avatar := form.FindAttrByCode("AvatarKey")

	if form.AvatarKey != nil && strings.TrimSpace(*form.AvatarKey) != "" {
		blob := models.AttachmentBlob{Key: *form.AvatarKey}

		repo := repository.NewAttachmentBlobRepository(nil, database.Db)
		if err := repo.Find(&blob); err != nil {

			avatar.AddError("is invalid")
		} else {
			form.UpdatedAttrs["Avatar"] = &models.Attachment{
				AttachmentBlob:   blob,
				AttachmentBlobId: blob.Id,
				Name:             "avatar",
			}
		}
	}
	return form
}

func (form *UserUpdateForm) validateEmail() *UserUpdateForm {
	key := "Email"
	emailField := form.FindAttrByCode(key)

	emailField.ValidateRequired()
	emailField.ValidateFormat(constants.EmailFormat, "")

	if emailField.IsClean() {
		form.UpdatedAttrs[key] = *form.Email
	}

	return form
}

func (form *UserUpdateForm) validateState() *UserUpdateForm {
	key := "State"
	userState := form.FindAttrByCode(key)
	userState.ValidateRequired()

	if userState.IsClean() {
		if userStateEnum, err := enums.ParseUserState(*form.State); err != nil {
			userState.AddError("is invalid")
		} else {
			if userStateEnum == enums.UserStateInactive {
				userState.AddError("State is invalid")
			} else {
				form.UpdatedAttrs[key] = userStateEnum
			}
		}
	}

	return form
}

func (form *UserUpdateForm) validateCompanyLevelId() *UserUpdateForm {
	key := "CompanyLevelId"
	level := form.FindAttrByCode(key)

	if form.CompanyLevelId != nil {
		level.ValidateMin(interface{}(int64(1)))
		level.ValidateMax(interface{}(int64(4)))

		companyLevelId, err := helpers.GqlIdToInt32(*form.CompanyLevelId)
		if err != nil {
			level.AddError("is invalid")
		}

		if level.IsClean() {
			form.UpdatedAttrs[key] = &companyLevelId
		}
	}

	return form
}

func (form *UserUpdateForm) validatePassword() *UserUpdateForm {
	password := form.FindAttrByCode("Password")

	if form.Password != nil {
		password.ValidateMin(interface{}(int64(6)))
		password.ValidateMax(interface{}(int64(20)))

		if encryptPassword, err := bcrypt.GenerateFromPassword([]byte(*form.Password), 10); err != nil {
			password.AddError(err)
		} else {
			form.UpdatedAttrs["EncryptedPassword"] = string(encryptPassword)
		}
	}

	return form
}
