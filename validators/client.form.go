package validators

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/pkg/helpers"
	"aio-server/repository"
)

type ClientForm struct {
	Form
	insightInputs.ClientFormInput
	Client *models.Client
	Repo   *repository.ClientRepository
}

func NewClientFormValidator(
	input *insightInputs.ClientFormInput,
	repo *repository.ClientRepository,
	client *models.Client,
) ClientForm {
	form := ClientForm{
		Form:            Form{},
		ClientFormInput: *input,
		Client:          client,
		Repo:            repo,
	}

	form.assignAttributes()

	return form
}

func (form *ClientForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	if form.Client.Id == 0 {
		return form.Repo.Create(form.Client)
	}

	return nil
}

func (form *ClientForm) assignAttributes() {
	form.AddAttributes(
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "name",
			},
			Value: helpers.GetStringOrDefault(form.Name),
		},
		&BoolAttribute{
			FieldAttribute: FieldAttribute{
				Code: "showOnHomepage",
			},
			Value: helpers.GetBoolOrDefault(form.ShowOnHomePage),
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "lockVersion",
			},
			Value: helpers.GetInt32OrDefault(form.LockVersion),
		},
	)
}

func (form *ClientForm) validate() error {
	form.validateName().validateShowOnHomePage().validateLockVersion()

	if form.Client.Id != 0 {
		form.validateLockVersion()
	}

	form.summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}
	return nil
}

func (form *ClientForm) validateName() *ClientForm {
	nameField := form.FindAttrByCode("name")

	nameField.ValidateRequired()

	nameField.ValidateMin(interface{}(int64(2)))
	nameField.ValidateMax(interface{}(int64(constants.MaxStringLength)))

	if nameField.IsClean() {
		form.Client.Name = *form.Name
	}

	return form
}

func (form *ClientForm) validateShowOnHomePage() *ClientForm {
	form.Client.ShowOnHomePage = helpers.GetBoolOrDefault(form.ShowOnHomePage)

	return form
}

func (form *ClientForm) validateLockVersion() *ClientForm {
	currentLockVersion := form.Client.LockVersion

	field := form.FindAttrByCode("lockVersion")

	field.ValidateRequired()

	if field.IsClean() {
		field.ValidateMin(interface{}(int64(currentLockVersion)))
	}

	return form
}
