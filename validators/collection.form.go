package validators

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/msInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
)

type CollectionForm struct {
	Form
	msInputs.CollectionFormInput
	Collection *models.Collection
	Repo       *repository.CollectionRepository
}

func NewCollectionFormValidator(
	input *msInputs.CollectionFormInput,
	repo *repository.CollectionRepository,
	collection *models.Collection,
) CollectionForm {
	form := CollectionForm{
		Form:                Form{},
		CollectionFormInput: *input,
		Collection:          collection,
		Repo:                repo,
	}

	form.assignAttributes(input)

	return form
}

func (form *CollectionForm) Save() error {

	if validationErr := form.validate(); validationErr != nil {
		return validationErr
	}

	if form.Collection.Id == 0 {
		return form.Repo.Create(form.Collection)
	}

	return form.Repo.Update(form.Collection)
}

func (form *CollectionForm) validate() error {
	form.validateTitle().summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}

	return nil
}

func (form *CollectionForm) validateTitle() *CollectionForm {
	title := form.FindAttrByCode("title")

	minTitleLength := 2
	maxTitleLength := int64(30)

	title.ValidateRequired()
	title.ValidateLimit(&minTitleLength, &maxTitleLength)

	return form
}

func (form *CollectionForm) assignAttributes(input *msInputs.CollectionFormInput) {
	title := helpers.GetStringOrDefault(input.Title)

	form.AddAttributes(
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Name: "Title",
				Code: "title",
			},
			Value: title,
		},
	)

	form.Collection.Title = title
}
