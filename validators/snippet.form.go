package validators

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs"
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/pkg/helpers"
	"aio-server/repository"
)

// SnippetForm represents a validator for snippet input.
type SnippetForm struct {
	Form
	inputs.MsSnippetFormInput
	Snippet *models.Snippet
	Repo    *repository.SnippetRepository
}

// NewSnippetFormValidator creates a new SnippetForm validator.
func NewSnippetFormValidator(input *inputs.MsSnippetFormInput, repo *repository.SnippetRepository, snippet *models.Snippet) SnippetForm {
	form := SnippetForm{
		Form:               Form{},
		MsSnippetFormInput: *input,
		Snippet:            snippet,
		Repo:               repo,
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
func (form *SnippetForm) assignAttributes(input *inputs.MsSnippetFormInput) {
	title := helpers.GetStringOrDefault(input.Title)
	content := helpers.GetStringOrDefault(input.Content)
	snippetType := helpers.GetInt32OrDefault(input.SnippetType)
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
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Name: "Snippet Type",
				Code: "snippetType",
			},
			Value:     snippetType,
			AllowZero: false,
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
	form.Snippet.SnippetType = int(snippetType)
}

func (form *SnippetForm) validateSnippetPrivateContent() *SnippetForm {
	snippetType := form.Snippet.SnippetType
	PasskeyAttr := form.FindAttrByCode("Passkey")

	if snippetType == 2 && PasskeyAttr != nil {
		// Private
		PasskeyAttr.ValidateRequired()

		min := 8
		max := int64(32)
		PasskeyAttr.ValidateLimit(&min, &max)
	}

	return form
}

func (form *SnippetForm) validateTitle() *SnippetForm {
	title := form.FindAttrByCode("title")
	minTitleLength := 5
	maxTitleLength := int64(constants.MaxStringLength)

	title.ValidateRequired()
	title.ValidateLimit(&minTitleLength, &maxTitleLength)

	return form
}

func (form *SnippetForm) validateContent() *SnippetForm {
	content := form.FindAttrByCode("content")
	minContentLength := 10
	maxContentLength := int64(constants.MaxLongTextLength)

	content.ValidateRequired()
	content.ValidateLimit(&minContentLength, &maxContentLength)

	return form
}

func (form *SnippetForm) validateSnippetType() *SnippetForm {
	snippetType := form.FindAttrByCode("snippetType")
	minSnippetType := 1
	maxSnippetType := int64(2)

	snippetType.ValidateRequired()
	snippetType.ValidateLimit(&minSnippetType, &maxSnippetType)

	return form
}

func (form *SnippetForm) validateLockVersion() *SnippetForm {
	newRecord := form.Snippet.Id == 0
	currentLockVersion := form.Snippet.LockVersion
	newLockVersion := helpers.GetInt32OrDefault(form.LockVersion) + 1
	formAttribute := form.FindAttrByCode("lockVersion")

	min := int(currentLockVersion)
	formAttribute.ValidateLimit(&min, nil)

	if newRecord {
		return form
	}

	if currentLockVersion >= newLockVersion {
		formAttribute.AddError("Attempted to update stale object")
	}

	return form
}
