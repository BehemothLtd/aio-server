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

type ProjectUpdateProjectAssigneeService struct {
	Ctx             *context.Context
	Db              *gorm.DB
	Args            insightInputs.ProjectUpdateProjectAssigneeInput
	Project         *models.Project
	ProjectAssignee *models.ProjectAssignee
}

func (pupas *ProjectUpdateProjectAssigneeService) Execute() error {
	if pupas.Args.ProjectId == "" {
		return exceptions.NewBadRequestError("Invalid Project ID")
	}

	if pupas.Args.Id == "" {
		return exceptions.NewBadRequestError("Invalid Id")
	}

	projectId, err := helpers.GqlIdToInt32(pupas.Args.ProjectId)

	if err != nil || projectId == 0 {
		return exceptions.NewBadRequestError("Invalid Project")
	}

	projectAssigneeId, err := helpers.GqlIdToInt32(pupas.Args.Id)
	if err != nil || projectAssigneeId == 0 {
		return exceptions.NewBadRequestError("Invalid Project Assignee Id")
	}

	pupas.Project = &models.Project{Id: projectId}

	projectRepo := repository.NewProjectRepository(pupas.Ctx, pupas.Db)
	if err := projectRepo.Find(pupas.Project); err != nil {
		return exceptions.NewBadRequestError("Invalid Project")
	}

	pupas.ProjectAssignee = &models.ProjectAssignee{Id: projectAssigneeId, ProjectId: pupas.Project.Id}
	repo := repository.NewProjectAssigneeRepository(pupas.Ctx, pupas.Db)

	if err := repo.Find(pupas.ProjectAssignee); err != nil {
		return exceptions.NewBadRequestError("Invalid Project Assignee")
	}

	form := validators.NewProjectAssigneeFormValidator(
		pupas.Args.Input,
		repo,
		*pupas.Project,
		pupas.ProjectAssignee,
	)

	if err := form.Save(); err != nil {
		return err
	}

	return nil
}
