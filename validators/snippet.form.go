package validators

import (
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
			Valid: true,
			Repo:  repo,
		},
		Input:   input,
		Snippet: snippet,
	}

	form.AddAttributes(
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Name: "Title",
				Code: "title",
			},
			Value: *input.Title,
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Name: "Content",
				Code: "content",
			},
			Value: *input.Content,
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Name: "Snippet Type",
				Code: "snippetType",
			},
			Value:     *input.SnippetType,
			AllowZero: false,
		},
	)

	form.assignSnippet()

	return form
}

func (form *SnippetForm) Validate() {
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
}

func (form *SnippetForm) assignSnippet() {
	form.Snippet.Title = *form.Input.Title
	form.Snippet.Content = *form.Input.Content
	form.Snippet.SnippetType = int(*form.Input.SnippetType)
}

func (form *SnippetForm) Create() error {
	return form.Repo.CreateSnippet(form.Snippet)
}

func (form *SnippetForm) Update() error {
	return form.Repo.UpdateSnippet(form.Snippet)
}
