package validators

import (
	"aio-server/gql/inputs"
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/repository"
)

type SnippetForm struct {
	Form
	Snippet *models.Snippet
}

func NewSnippetFormValidator(input *inputs.MsSnippetInput, repo *repository.Repository, snippet *models.Snippet) SnippetForm {
	form := SnippetForm{
		Form: Form{
			Valid: true,
			Repo:  repo,
		},
		Snippet: snippet,
	}

	form.assignAttributes(input)

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

func (form *SnippetForm) Create() error {
	return form.Repo.CreateSnippet(form.Snippet)
}

func (form *SnippetForm) Update() error {
	return form.Repo.UpdateSnippet(form.Snippet)
}

func (form *SnippetForm) assignAttributes(input *inputs.MsSnippetInput) {
	formInput := input.ToFormInput()

	form.AddAttributes(
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Name: "Title",
				Code: "title",
			},
			Value: formInput.Title,
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Name: "Content",
				Code: "content",
			},
			Value: formInput.Content,
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Name: "Snippet Type",
				Code: "snippetType",
			},
			Value:     formInput.SnippetType,
			AllowZero: false,
		},
	)

	form.Snippet.Title = formInput.Title
	form.Snippet.Content = formInput.Content
	form.Snippet.SnippetType = int(formInput.SnippetType)
}
