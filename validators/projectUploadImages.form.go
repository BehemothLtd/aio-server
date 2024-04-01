package validators

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"strings"
)

type ProjectUploadImagesForm struct {
	Form
	insightInputs.ProjectUploadImagesFormInput
	Project       *models.Project
	UpdateProject models.Project
	ProjectRepo   repository.ProjectRepository
	Repo          repository.AttachmentBlobRepository
}

func NewProjectUploadImagesFormValidator(
	input *insightInputs.ProjectUploadImagesFormInput,
	repo repository.AttachmentBlobRepository,
	projectRepo repository.ProjectRepository,
	project *models.Project,
) ProjectUploadImagesForm {
	form := ProjectUploadImagesForm{
		Form:                         Form{},
		ProjectUploadImagesFormInput: *input,
		Project:                      project,
		Repo:                         repo,
		ProjectRepo:                  projectRepo,
	}

	form.assignAttributes()

	return form
}

func (form *ProjectUploadImagesForm) assignAttributes() {
	form.AddAttributes(
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "logoKey",
			},
			Value: helpers.GetStringOrDefault(form.LogoKey),
		},
		&SliceAttribute[string]{
			FieldAttribute: FieldAttribute{
				Code: "fileKeys",
			},
			Value: form.FileKeys,
		},
	)
}

func (form *ProjectUploadImagesForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	if err := form.ProjectRepo.UpdateFiles(form.Project); err != nil {
		return exceptions.NewUnprocessableContentError("", exceptions.ResourceModificationError{
			"base": {err.Error()},
		})
	}

	return nil
}

func (form *ProjectUploadImagesForm) validate() error {
	form.validateLogo().validateFiles().summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}

	return nil
}

func (form *ProjectUploadImagesForm) validateLogo() *ProjectUploadImagesForm {
	if form.LogoKey != nil && strings.TrimSpace(*form.LogoKey) != "" {
		field := form.FindAttrByCode("logoKey")
		blob := models.AttachmentBlob{Key: *form.LogoKey}

		if err := form.Repo.Find(&blob); err != nil {
			field.AddError("is invalid")
		} else {
			form.Project.Logo = &models.Attachment{
				AttachmentBlob:   blob,
				AttachmentBlobId: blob.Id,
				Name:             "logo",
			}
		}
	}

	return form
}

func (form *ProjectUploadImagesForm) validateFiles() *ProjectUploadImagesForm {
	if form.FileKeys != nil {
		fieldKey := "fileKeys"
		form.Project.Files = []*models.Attachment{}

		for i, fileKey := range *form.FileKeys {
			blob := models.AttachmentBlob{Key: fileKey}

			if err := form.Repo.Find(&blob); err != nil {
				form.AddErrorDirectlyToField(form.NestedDirectItemFieldKey(fieldKey, i), []interface{}{"is invalid"})
			} else {
				form.Project.Files = append(form.Project.Files, &models.Attachment{
					AttachmentBlob:   blob,
					AttachmentBlobId: blob.Id,
					Name:             "files",
				})
			}
		}
	}

	return form
}
