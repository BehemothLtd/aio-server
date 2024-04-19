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
	form.AddAttributes(
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "CurrentPassword",
			},
			Value: helpers.GetStringOrDefault(form.CurrentPassword),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "Password",
			},
			Value: helpers.GetStringOrDefault(form.Password),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "PasswordConfirmation",
			},
			Value: helpers.GetStringOrDefault(form.PasswordConfirmation),
		},
	)
}

func (form *UserPasswordForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	} else {
		if password, err := bcrypt.GenerateFromPassword([]byte(*form.Password), 10); err != nil {
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
	passwordField := form.FindAttrByCode("CurrentPassword")

	passwordField.ValidateRequired()

	if passwordField.IsClean() {
		if err := bcrypt.CompareHashAndPassword([]byte(form.User.EncryptedPassword), []byte(*form.CurrentPassword)); err != nil {
			passwordField.AddError("is incorrect")
		}
	}

	return form
}

func (form *UserPasswordForm) validateNewPassword() *UserPasswordForm {
	newPasswordField := form.FindAttrByCode("Password")
	newPasswordConfirmationField := form.FindAttrByCode("PasswordConfirmation")

	newPasswordField.ValidateRequired()
	newPasswordField.ValidateMin(interface{}(int64(8)))
	newPasswordField.ValidateMax(interface{}(int64(20)))

	newPasswordConfirmationField.ValidateRequired()
	newPasswordConfirmationField.ValidateMin(interface{}(int64(8)))
	newPasswordConfirmationField.ValidateMax(interface{}(int64(20)))

	if newPasswordField.IsClean() && newPasswordConfirmationField.IsClean() {
		if *form.Password != *form.PasswordConfirmation {
			newPasswordConfirmationField.AddError("needs to be the same with new password")
		}
	}

	return form
}
