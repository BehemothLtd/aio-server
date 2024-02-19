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
	Input   *inputs.MsSnippetInput
	Snippet *models.Snippet
	Repo    *repository.SnippetRepository
}

// NewSnippetFormValidator creates a new SnippetForm validator.
func NewSnippetFormValidator(input *inputs.MsSnippetInput, repo *repository.SnippetRepository, snippet *models.Snippet) SnippetForm {
	form := SnippetForm{
		Form:    Form{},
		Input:   input,
		Snippet: snippet,
		Repo:    repo,
	}
	form.assignAttributes(input)
	return form
}

// Save saves the snippet after validation.
func (form *SnippetForm) Save() error {
	if validationErr := form.validate(); validationErr != nil {
		return validationErr
	}

	err := form.Snippet.EncryptContent(*form.Input.Passkey)

	if err != nil {
		return err
	}

	if form.Snippet.Id == 0 {
		form.Snippet.Slug = helpers.NewUUID()
		return form.Repo.Create(form.Snippet)
	}
	return form.Repo.Update(form.Snippet)
}

// validate validates the snippet form.
func (form *SnippetForm) validate() error {
	title := form.FindAttrByCode("title")
	if title != nil {
		minTitleLength := 5
		maxTitleLength := int64(constants.MaxStringLength)

		title.ValidateRequired()
		title.ValidateLimit(&minTitleLength, &maxTitleLength)
	}

	content := form.FindAttrByCode("content")
	if content != nil {
		minContentLength := 10
		maxContentLength := int64(constants.MaxLongTextLength)

		content.ValidateRequired()
		content.ValidateLimit(&minContentLength, &maxContentLength)
	}

	snippetType := form.FindAttrByCode("snippetType")
	if snippetType != nil {
		minSnippetType := 1
		maxSnippetType := int64(2)

		snippetType.ValidateRequired()
		snippetType.ValidateLimit(&minSnippetType, &maxSnippetType)
	}

	form.validateSnippetPrivateContent()
	form.SummaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}
	return nil
}

// assignAttributes assigns attributes to the snippet form.
func (form *SnippetForm) assignAttributes(input *inputs.MsSnippetInput) {
	title := helpers.GetStringOrDefault(input.Title)
	content := helpers.GetStringOrDefault(input.Content)
	snippetType := helpers.GetInt32OrDefault(input.SnippetType)
	Passkey := helpers.GetStringOrDefault(input.Passkey)

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
				Code: "Passkey",
			},
			Value: Passkey,
		},
	)

	form.Snippet.Title = title
	form.Snippet.Content = content
	form.Snippet.SnippetType = int(snippetType)
}

func (form *SnippetForm) validateSnippetPrivateContent() {
	snippetType := form.Snippet.SnippetType
	PasskeyAttr := form.FindAttrByCode("Passkey")

	if snippetType == 2 {
		// Private
		PasskeyAttr.ValidateRequired()

		min := 8
		max := int64(32)
		PasskeyAttr.ValidateLimit(&min, &max)
	} else {
		return
	}
}
