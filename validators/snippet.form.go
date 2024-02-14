package validators

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs"
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/repository"
)

type SnippetForm struct {
	Form
	Input   *inputs.MsSnippetInput
	Snippet *models.Snippet
}

func NewSnippetFormValidator(input *inputs.MsSnippetInput, repo *repository.Repository, snippet *models.Snippet) SnippetForm {
	form := SnippetForm{
		Form: Form{
			Repo: repo,
		},
		Input:   input,
		Snippet: snippet,
	}

	form.assignAttributes(input)

	return form
}

func (form *SnippetForm) Save() error {
	validationErr := form.validate()

	if validationErr != nil {
		return validationErr
	}

	var saveErr error

	if form.Snippet.Id == 0 {
		// Create
		saveErr = form.Repo.CreateSnippet(form.Snippet)
	} else {
		// Update
		saveErr = form.Repo.UpdateSnippet(form.Snippet)
	}

	return saveErr
}

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

	form.SummaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError(nil, &form.Errors)
	} else {
		return nil
	}
}

func (form *SnippetForm) assignAttributes(input *inputs.MsSnippetInput) {
	title := input.GetTitle()
	content := input.GetContent()
	snippetType := input.GetSnippetType()

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
	)

	form.Snippet.Title = title
	form.Snippet.Content = content
	form.Snippet.SnippetType = int(snippetType)
}
