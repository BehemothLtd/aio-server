package validators

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"strings"
)

type ClientCreateForm struct {
	Form
	insightInputs.ClientFormInput
	Client             *models.Client
	Repo               *repository.ClientRepository
	AttachmentBlobRepo repository.AttachmentBlobRepository
}

func NewClientCreateFormValidator(
	input *insightInputs.ClientFormInput,
	repo *repository.ClientRepository,
	client *models.Client,
	attachmentBlobRepo repository.AttachmentBlobRepository,
) ClientCreateForm {
	form := ClientCreateForm{
		Form:               Form{},
		ClientFormInput:    *input,
		Client:             client,
		Repo:               repo,
		AttachmentBlobRepo: attachmentBlobRepo,
	}

	form.assignAttributes()

	return form
}

func (form *ClientCreateForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	return form.Repo.Create(form.Client)
}

func (form *ClientCreateForm) assignAttributes() {
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
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "avatarKey",
			},
			Value: helpers.GetStringOrDefault(form.AvatarKey),
		},
	)
}

func (form *ClientCreateForm) validate() error {
	form.validateName().validateShowOnHomePage().validateAvatar()

	form.summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}
	return nil
}

func (form *ClientCreateForm) validateName() *ClientCreateForm {
	nameField := form.FindAttrByCode("name")

	nameField.ValidateRequired()

	nameField.ValidateMin(interface{}(int64(2)))
	nameField.ValidateMax(interface{}(int64(constants.MaxStringLength)))

	if nameField.IsClean() {
		form.Client.Name = *form.Name
	}

	return form
}

func (form *ClientCreateForm) validateShowOnHomePage() *ClientCreateForm {
	form.Client.ShowOnHomePage = helpers.GetBoolOrDefault(form.ShowOnHomePage)

	return form
}

func (form *ClientCreateForm) validateAvatar() *ClientCreateForm {
	if form.AvatarKey != nil && strings.TrimSpace(*form.AvatarKey) != "" {
		field := form.FindAttrByCode("avatarKey")
		blob := models.AttachmentBlob{Key: *form.AvatarKey}

		if err := form.AttachmentBlobRepo.Find(&blob); err != nil {
			field.AddError("is invalid")
		} else {
			form.Client.Avatar = &models.Attachment{
				AttachmentBlob:   blob,
				AttachmentBlobId: blob.Id,
				Name:             "avatar",
			}
		}
	}

	return form
}
