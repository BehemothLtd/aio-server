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

type ClientUpdateForm struct {
	Form
	insightInputs.ClientFormInput
	Client             *models.Client
	UpdateClient       models.Client
	Repo               *repository.ClientRepository
	AttachmentBlobRepo repository.AttachmentBlobRepository
}

func NewClientUpdateFormValidator(
	input *insightInputs.ClientFormInput,
	repo *repository.ClientRepository,
	client *models.Client,
	attachmentBlobRepo repository.AttachmentBlobRepository,
) ClientUpdateForm {
	form := ClientUpdateForm{
		Form:               Form{},
		ClientFormInput:    *input,
		UpdateClient:       models.Client{},
		Client:             client,
		Repo:               repo,
		AttachmentBlobRepo: attachmentBlobRepo,
	}

	form.assignAttributes()

	return form
}

func (form *ClientUpdateForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	if err := form.Repo.Update(form.Client, form.UpdateClient); err != nil {
		return exceptions.NewUnprocessableContentError("", exceptions.ResourceModificationError{
			"base": {err.Error()},
		})
	}

	return nil
}

func (form *ClientUpdateForm) assignAttributes() {
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
			Value:     helpers.GetInt32OrDefault(form.LockVersion),
			AllowZero: true,
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "avatarKey",
			},
			Value: helpers.GetStringOrDefault(form.AvatarKey),
		},
	)
}

func (form *ClientUpdateForm) validate() error {
	form.validateName().validateShowOnHomePage().validateLockVersion().validateAvatar()

	form.summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}

	return nil
}

func (form *ClientUpdateForm) validateName() *ClientUpdateForm {
	nameField := form.FindAttrByCode("name")

	nameField.ValidateRequired()

	nameField.ValidateMin(interface{}(int64(2)))
	nameField.ValidateMax(interface{}(int64(constants.MaxStringLength)))

	if nameField.IsClean() {
		form.UpdateClient.Name = *form.Name
	}

	return form
}

func (form *ClientUpdateForm) validateAvatar() *ClientUpdateForm {
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

func (form *ClientUpdateForm) validateShowOnHomePage() *ClientUpdateForm {
	form.UpdateClient.ShowOnHomePage = helpers.GetBoolOrDefault(form.ShowOnHomePage)

	return form
}

func (form *ClientUpdateForm) validateLockVersion() *ClientUpdateForm {
	currentLockVersion := form.UpdateClient.LockVersion

	field := form.FindAttrByCode("lockVersion")

	field.ValidateRequired()

	if field.IsClean() {
		field.ValidateMin(interface{}(int64(currentLockVersion)))
	}

	return form
}
