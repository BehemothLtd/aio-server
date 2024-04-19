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
				Code: "slackId",
			},
			Value: helpers.GetStringOrDefault(input.SlackId),
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
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "password",
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
	about := form.FindAttrByCode("about")
	about.ValidateMax(interface{}(int64(constants.MaxLongTextLength)))

	if about.IsClean() {
		form.User.About = form.About
	}

	return form
}

func (form *UserCreateForm) validateFullName() *UserCreateForm {
	fullName := form.FindAttrByCode("fullName")

	fullName.ValidateRequired()
	fullName.ValidateMin(interface{}(int64(10)))
	fullName.ValidateMax(interface{}(int64(constants.MaxLongTextLength)))

	if fullName.IsClean() {
		form.User.FullName = *form.FullName
	}

	return form
}

func (form *UserCreateForm) validateSlackId() *UserCreateForm {
	slackId := form.FindAttrByCode("slackId")

	slackId.ValidateRequired()
	slackId.ValidateMin(interface{}(int64(11)))
	slackId.ValidateMax(interface{}(int64(constants.MaxStringLength)))

	if slackId.IsClean() {
		form.User.SlackId = form.SlackId
	}

	return form
}

func (form *UserCreateForm) validateAvatarKey() *UserCreateForm {
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

func (form *UserCreateForm) validateEmail() *UserCreateForm {
	emailField := form.FindAttrByCode("email")

	emailField.ValidateRequired()
	emailField.ValidateFormat(constants.EmailFormat, "")

	if emailField.IsClean() {
		form.User.Email = *form.Email
	}

	return form
}

func (form *UserCreateForm) validatePassword() *UserCreateForm {
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
