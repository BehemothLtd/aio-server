package insightServices

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"aio-server/validators"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type ProjectSprintCreateService struct {
	Ctx           *context.Context
	Db            *gorm.DB
	Args          insightInputs.ProjectSprintCreateInput
	ProjectSprint *models.ProjectSprint
}

func (service *ProjectSprintCreateService) Execute() error {
	if service.Args.ProjectId == "" {
		return exceptions.NewBadRequestError("Invalid Project")
	}

	projectId, err := helpers.GqlIdToInt32(service.Args.ProjectId)
	if err != nil || projectId == 0 {
		return exceptions.NewBadRequestError("Invalid Id")
	}

	project := models.Project{Id: projectId}
	projectRepo := repository.NewProjectRepository(service.Ctx, service.Db)
	if err := projectRepo.Find(&project); err != nil {
		return exceptions.NewBadRequestError("Invalid Project")
	}

	service.ProjectSprint.ProjectId = project.Id

	form := validators.NewProjectSprintFormValidator(
		&service.Args.Input,
		repository.NewProjectSprintRepository(service.Ctx, service.Db),
		service.ProjectSprint,
		project,
	)

	if err := form.Save(); err != nil {
		return err
	}

	return nil
}
