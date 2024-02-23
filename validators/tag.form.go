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
	form.validateName().
		validateLockVersion().
		summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}

	return nil
}

// assignAttributes assigns attributes to the tag form.
func (form *TagForm) assignAttributes(input *msInputs.TagFormInput) {
	name := helpers.GetStringOrDefault(input.Name)
	lockVersion := helpers.GetInt32OrDefault(input.LockVersion)

	form.AddAttributes(
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Name: "Name",
				Code: "name",
			},
			Value: name,
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Name: "Lock Version",
				Code: "lockVersion",
			},
			Value: lockVersion,
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

	if form.Name != nil {
		presentedTag := models.Tag{}
		if err := form.Repo.FindByName(&presentedTag, *form.Name); err == nil {
			if presentedTag.Id != 0 && presentedTag.Id != form.Tag.Id {
				name.AddError("is used.")
			}
		}
	}

	return form
}

func (form *TagForm) validateLockVersion() *TagForm {
	newRecord := form.Tag.Id == 0
	currentLockVersion := form.Tag.LockVersion
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
