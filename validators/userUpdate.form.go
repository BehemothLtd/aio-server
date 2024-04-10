package validators

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
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
	form.summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}

	return nil
}
