package validators

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserPasswordForm struct {
	Form
	insightInputs.SelfUpdatePasswordFormInput
	User *models.User
	Repo *repository.UserRepository
}

func NewUserPasswordFormValidator(
	input *insightInputs.SelfUpdatePasswordFormInput,
	repo *repository.UserRepository,
	user *models.User,
) UserPasswordForm {
	form := UserPasswordForm{
		Form:                        Form{},
		SelfUpdatePasswordFormInput: *input,
		User:                        user,
		Repo:                        repo,
	}

	form.assignAttributes(input)

	return form
}

func (form *UserPasswordForm) assignAttributes(input *insightInputs.SelfUpdatePasswordFormInput) {
	currentPassword := helpers.GetStringOrDefault(&input.CurrentPassword)
	password := helpers.GetStringOrDefault(&input.Password)
	passwordConfirmation := helpers.GetStringOrDefault(&input.PasswordConfirmation)

	form.AddAttributes(
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "currentPassword",
			},
			Value: currentPassword,
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "password",
			},
			Value: password,
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "passwordConfirmation",
			},
			Value: passwordConfirmation,
		},
	)
}

func (form *UserPasswordForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	} else {
		if password, err := bcrypt.GenerateFromPassword([]byte(form.Password), 10); err != nil {
			return err
		} else {
			form.User.EncryptedPassword = string(password)

			return form.Repo.Update(form.User, []string{"EncryptedPassword"})
		}
	}
}

func (form *UserPasswordForm) validate() error {
	form.validatePassword().
		validateNewPassword().
		summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}

	return nil
}

func (form *UserPasswordForm) validatePassword() *UserPasswordForm {
	passwordField := form.FindAttrByCode("currentPassword")

	passwordField.ValidateRequired()

	if err := bcrypt.CompareHashAndPassword([]byte(form.User.EncryptedPassword), []byte(form.CurrentPassword)); err != nil {
		passwordField.AddError("is incorrect")
	}

	return form
}

func (form *UserPasswordForm) validateNewPassword() *UserPasswordForm {
	newPasswordField := form.FindAttrByCode("password")
	newPasswordConfirmationField := form.FindAttrByCode("passwordConfirmation")

	min := 8
	max := int64(20)

	newPasswordField.ValidateRequired()
	newPasswordField.ValidateLimit(&min, &max)

	newPasswordConfirmationField.ValidateRequired()
	newPasswordConfirmationField.ValidateLimit(&min, &max)

	if form.Password != form.PasswordConfirmation {
		newPasswordConfirmationField.AddError("needs to be the same with new password")
	}

	return form
}
