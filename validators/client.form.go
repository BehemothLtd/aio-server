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
		// &StringAttribute{
		// 	FieldAttribute: FieldAttribute{
	// 		Code: "ShowOnHomepage",
	// 	},
	// 	Value: helpers.GetStringOrDefault(form.ShowOnHomePage),
	// },
	// &StringAttribute{
	// 	FieldAttribute: FieldAttribute{
	// 		Code: "LockVersion",
	// 	},
	// 	Value: helpers.GetStringOrDefault(form.LockVersion),
	// },
	)
}

func (form *ClientForm) validate() error {
	form.validateName()

	if form.Client.Id != 0 {
		form.validateLockVersion()
	}

	form.summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}
	return nil
}

func (form *ClientForm) validateColor() *ClientForm {
	colorField := form.FindAttrByCode("color")

	colorField.ValidateRequired()

	if colorField.IsClean() {
		form.Client.Name = *form.Name
	}

	return form
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

// func (form *ClientForm) validateClientType() *ClientForm {
// 	typeField := form.FindAttrByCode("statusType")

// 	typeField.ValidateRequired()

// 	if typeField.IsClean() {
// 		if statusType, err := enums.ParseIssueStatusStatusType(*form.StatusType); err != nil {
// 			typeField.AddError("is invalid issue status type")
// 		} else {
// 			form.IssueStatus.StatusType = statusType
// 		}
// 	}

// 	return form
// }

func (form *ClientForm) validateLockVersion() *ClientForm {
	currentLockVersion := form.Client.LockVersion

	field := form.FindAttrByCode("lockVersion")

	field.ValidateRequired()

	if field.IsClean() {
		field.ValidateMin(interface{}(int64(currentLockVersion)))
	}

	return form
}
