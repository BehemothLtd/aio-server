package insightServices

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
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

	project := models.Project{Id: projectId}
	projectRepo := repository.NewProjectRepository(puis.Ctx, puis.Db.Preload("ProjectIssueStatuses").Preload("Issues").Preload("ProjectSprints"))

	if err := projectRepo.Find(&project); err != nil {
		return exceptions.NewRecordNotFoundError()
	}

	puis.Project.Id = projectId

	return nil
}
