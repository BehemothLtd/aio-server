package validators

import (
	"aio-server/database"
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type UserCreateForm struct {
	Form
	insightInputs.UserFormInput
	User *models.User
	Repo *repository.UserRepository
}

func NewUserCreateFormValidator(
	input *insightInputs.UserFormInput,
	repo *repository.UserRepository,
	user *models.User,
) UserCreateForm {
	form := UserCreateForm{
		Form:          Form{},
		UserFormInput: *input,
		User:          user,
		Repo:          repo,
	}

	form.assignAttributes(input)

	return form
}

func (form *UserCreateForm) assignAttributes(input *insightInputs.UserFormInput) {

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
				Code: "SlackId",
			},
			Value: helpers.GetStringOrDefault(input.SlackId),
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
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "Password",
			},
			Value: helpers.GetStringOrDefault(input.Password),
		},
	)
}

func (form *UserCreateForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}
	return form.Repo.Create(form.User)
}

func (form *UserCreateForm) validate() error {
	form.validateFullName().
		validateEmail().
		validateSlackId().
		validateAbout().
		validateAvatarKey().
		validatePassword().
		summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}

	return nil
}

func (form *UserCreateForm) validateAbout() *UserCreateForm {
	about := form.FindAttrByCode("About")
	about.ValidateMax(interface{}(int64(constants.MaxLongTextLength)))

	if about.IsClean() {
		form.User.About = form.About
	}

	return form
}

func (form *UserCreateForm) validateFullName() *UserCreateForm {
	fullName := form.FindAttrByCode("FullName")

	fullName.ValidateRequired()
	fullName.ValidateMin(interface{}(int64(10)))
	fullName.ValidateMax(interface{}(int64(constants.MaxLongTextLength)))

	if fullName.IsClean() {
		form.User.FullName = *form.FullName
	}

	return form
}

func (form *UserCreateForm) validateSlackId() *UserCreateForm {
	slackId := form.FindAttrByCode("SlackId")

	slackId.ValidateRequired()
	slackId.ValidateMin(interface{}(int64(11)))
	slackId.ValidateMax(interface{}(int64(constants.MaxStringLength)))

	if slackId.IsClean() {
		form.User.SlackId = form.SlackId
	}

	return form
}

func (form *UserCreateForm) validateAvatarKey() *UserCreateForm {
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

func (form *UserCreateForm) validateEmail() *UserCreateForm {
	emailField := form.FindAttrByCode("Email")

	emailField.ValidateRequired()
	emailField.ValidateFormat(constants.EmailFormat, "")

	if emailField.IsClean() {
		form.User.Email = *form.Email
	}

	return form
}

func (form *UserCreateForm) validatePassword() *UserCreateForm {
	password := form.FindAttrByCode("Password")

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
