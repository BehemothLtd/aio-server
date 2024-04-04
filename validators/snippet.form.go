package validators

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/inputs/msInputs"
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/pkg/helpers"
	"aio-server/repository"
)

// SnippetForm represents a validator for snippet input.
type SnippetForm struct {
	Form
	msInputs.SnippetFormInput
	Snippet *models.Snippet
	Repo    *repository.SnippetRepository
}

// NewSnippetFormValidator creates a new SnippetForm validator.
func NewSnippetFormValidator(
	input *msInputs.SnippetFormInput,
	repo *repository.SnippetRepository,
	snippet *models.Snippet,
) SnippetForm {
	form := SnippetForm{
		Form:             Form{},
		SnippetFormInput: *input,
		Snippet:          snippet,
		Repo:             repo,
	}
	form.assignAttributes(input)

	return form
}

// Save saves the snippet after validation.
func (form *SnippetForm) Save() error {
	if validationErr := form.validate(); validationErr != nil {
		return validationErr
	}

	passkey := helpers.GetStringOrDefault(form.Passkey)

	if passkey != "" {
		err := form.Snippet.EncryptContent(*form.Passkey)

		if err != nil {
			return err
		}
	}

	if form.Snippet.Id == 0 {
		form.Snippet.Slug = helpers.NewUUID()
		return form.Repo.Create(form.Snippet)
	}
	return form.Repo.Update(form.Snippet)
}

// validate validates the snippet form.
func (form *SnippetForm) validate() error {
	form.validateTitle().
		validateContent().
		validateSnippetType().
		validateSnippetPrivateContent().
		validateLockVersion().
		summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}

	return nil
}

// assignAttributes assigns attributes to the snippet form.
func (form *SnippetForm) assignAttributes(input *msInputs.SnippetFormInput) {
	title := helpers.GetStringOrDefault(input.Title)
	content := helpers.GetStringOrDefault(input.Content)
	snippetType := helpers.GetStringOrDefault(input.SnippetType)
	passkey := helpers.GetStringOrDefault(input.Passkey)
	lockVersion := helpers.GetInt32OrDefault(input.LockVersion)

	form.AddAttributes(
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Name: "Title",
				Code: "title",
			},
			Value: title,
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Name: "Content",
				Code: "content",
			},
			Value: content,
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Name: "Snippet Type",
				Code: "snippetType",
			},
			Value: snippetType,
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Name: "Pass Key",
				Code: "passkey",
			},
			Value: passkey,
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Name: "Lock Version",
				Code: "lockVersion",
			},
			Value: lockVersion,
		},
	)

	form.Snippet.Title = title
	form.Snippet.Content = content
	form.Snippet.SnippetType = enums.SnippetType(snippetType)
}

func (form *SnippetForm) validateSnippetPrivateContent() *SnippetForm {
	snippetType := form.Snippet.SnippetType
	PasskeyAttr := form.FindAttrByCode("passkey")

	if snippetType == enums.SnippetTypePrivate {
		// Private
		PasskeyAttr.ValidateRequired()

		PasskeyAttr.ValidateMin(interface{}(int64(8)))
		PasskeyAttr.ValidateMax(interface{}(int64(32)))
	}

	return form
}

func (form *SnippetForm) validateTitle() *SnippetForm {
	title := form.FindAttrByCode("title")

	title.ValidateRequired()
	title.ValidateMin(interface{}(int64(5)))
	title.ValidateMax(interface{}(int64(constants.MaxStringLength)))

	return form
}

func (form *SnippetForm) validateContent() *SnippetForm {
	content := form.FindAttrByCode("content")

	content.ValidateRequired()
	content.ValidateMin(interface{}(int64(10)))
	content.ValidateMax(interface{}(int64(constants.MaxLongTextLength)))

	return form
}

func (form *SnippetForm) validateSnippetType() *SnippetForm {
	snippetType := form.FindAttrByCode("snippetType")

	snippetType.ValidateRequired()

	if form.SnippetType != nil {
		snippetTypeValue := enums.SnippetType(*form.SnippetType)

		if !snippetTypeValue.IsValid() {
			snippetType.AddError("is invalid")
		}
	}

	return form
}

func (form *SnippetForm) validateLockVersion() *SnippetForm {
	newRecord := form.Snippet.Id == 0
	currentLockVersion := form.Snippet.LockVersion
	newLockVersion := helpers.GetInt32OrDefault(form.LockVersion) + 1
	formAttribute := form.FindAttrByCode("lockVersion")

	formAttribute.ValidateMin(interface{}(int64(currentLockVersion)))

	if newRecord {
		return form
	}

	if currentLockVersion >= newLockVersion {
		formAttribute.AddError("Attempted to update stale object")
	}

	return form
}
