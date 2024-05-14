package insightServices

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type ProjectSprintArchiveService struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args insightInputs.ProjectSprintArchiveInput
}

func (service *ProjectSprintArchiveService) Execute() error {
	// Validate Project
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
	// End validate Project

	// Validate Sprint
	sprintId, err := helpers.GqlIdToInt32(service.Args.Id)
	if err != nil || sprintId == 0 {
		return exceptions.NewBadRequestError("Invalid Id")
	}

	sprintRepo := repository.NewProjectSprintRepository(service.Ctx, service.Db)

	sprint := models.ProjectSprint{Id: sprintId, ProjectId: projectId}
	if err := sprintRepo.Find(&sprint); err != nil {
		return exceptions.NewRecordNotFoundError()
	}
	// End Validate Sprint

	// Validate MoveToSprint
	moveToId, err := helpers.GqlIdToInt32(service.Args.MoveToId)
	if err != nil || moveToId == 0 {
		return exceptions.NewBadRequestError("Invalid Id")
	}

	moveToSprint := models.ProjectSprint{Id: moveToId, ProjectId: projectId}
	if err := sprintRepo.Find(&moveToSprint); err != nil {
		return exceptions.NewRecordNotFoundError()
	}
	// End Validate MoveToSprint

	if err := projectRepo.ChangeActiveSprint(&project, &sprint, moveToSprint); err != nil {
		return exceptions.NewUnprocessableContentError(err.Error(), nil)
	}

	return nil
}
