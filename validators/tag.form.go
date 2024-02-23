package validators

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/msInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
)

// TagForm represents a validator for tag input.
type TagForm struct {
	Form
	msInputs.TagFormInput
	Tag  *models.Tag
	Repo *repository.TagRepository
}

func NewTagFormValidator(
	input *msInputs.TagFormInput,
	repo *repository.TagRepository,
	tag *models.Tag,
) TagForm {
	form := TagForm{
		Form:         Form{},
		TagFormInput: *input,
		Tag:          tag,
		Repo:         repo,
	}

	form.assignAttributes(input)

	return form
}

func (form *TagForm) Save() error {
	if validationErr := form.validate(); validationErr != nil {
		return validationErr
	}

	if form.Tag.Id == 0 {
		return form.Repo.Create(form.Tag)
	}
	return form.Repo.Update(form.Tag)
}

// validate validates the snippet form.
func (form *TagForm) validate() error {
	form.validateName().summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}

	return nil
}

// assignAttributes assigns attributes to the tag form.
func (form *TagForm) assignAttributes(input *msInputs.TagFormInput) {
	name := helpers.GetStringOrDefault(input.Name)

	form.AddAttributes(
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Name: "Name",
				Code: "name",
			},
			Value: name,
		},
	)

	form.Tag.Name = name
}

func (form *TagForm) validateName() *TagForm {
	name := form.FindAttrByCode("name")

	minNameLength := 2
	maxNameLength := int64(20)

	name.ValidateRequired()
	name.ValidateLimit(&minNameLength, &maxNameLength)

	presentedTag := models.Tag{}
	if err := form.Repo.FindByName(&presentedTag, *form.Name); err == nil {
		if presentedTag.Id != 0 {
			name.AddError("is used.")
		}
	}

	return form
}
