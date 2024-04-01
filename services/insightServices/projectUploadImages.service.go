package insightServices

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"aio-server/validators"
	"context"

	"gorm.io/gorm"
)

type ProjectUploadImagesService struct {
	Ctx     *context.Context
	Db      *gorm.DB
	Args    insightInputs.ProjectUploadImagesInput
	Project *models.Project
}

func (puis *ProjectUploadImagesService) Execute() error {
	if puis.Args.Id == "" {
		return exceptions.NewBadRequestError("Invalid Project ID")
	}

	projectId, err := helpers.GqlIdToInt32(puis.Args.Id)

	if err != nil || projectId == 0 {
		return exceptions.NewBadRequestError("Invalid Project")
	}

	puis.Project.Id = projectId
	projectRepo := repository.NewProjectRepository(puis.Ctx, puis.Db)

	if err := projectRepo.Find(puis.Project); err != nil {
		return exceptions.NewRecordNotFoundError()
	}

	form := validators.NewProjectUploadImagesFormValidator(
		&puis.Args.Input,
		*repository.NewAttachmentBlobRepository(puis.Ctx, puis.Db),
		*projectRepo,
		puis.Project,
	)

	if err := form.Save(); err != nil {
		return err
	}

	return nil
}
