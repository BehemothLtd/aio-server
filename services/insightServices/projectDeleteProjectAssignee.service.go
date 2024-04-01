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

type ProjectDeleteProjectAssigneeService struct {
	Ctx             *context.Context
	Db              *gorm.DB
	Args            insightInputs.ProjectDeleteProjectAssigneeInput
	Project         *models.Project
	ProjectAssignee *models.ProjectAssignee
}

func (pdpas *ProjectDeleteProjectAssigneeService) Execute() error {
	if pdpas.Args.ProjectId == "" {
		return exceptions.NewBadRequestError("Invalid Project ID")
	}

	if pdpas.Args.Id == "" {
		return exceptions.NewBadRequestError("Invalid Id")
	}

	projectId, err := helpers.GqlIdToInt32(pdpas.Args.ProjectId)

	if err != nil || projectId == 0 {
		return exceptions.NewBadRequestError("Invalid Project")
	}

	projectAssigneeId, err := helpers.GqlIdToInt32(pdpas.Args.Id)
	if err != nil || projectAssigneeId == 0 {
		return exceptions.NewBadRequestError("Invalid Project Assignee Id")
	}

	pdpas.Project = &models.Project{Id: projectId}

	projectRepo := repository.NewProjectRepository(pdpas.Ctx, pdpas.Db)
	if err := projectRepo.Find(pdpas.Project); err != nil {
		return exceptions.NewBadRequestError("Invalid Project")
	}

	pdpas.ProjectAssignee = &models.ProjectAssignee{Id: projectAssigneeId, ProjectId: pdpas.Project.Id}
	repo := repository.NewProjectAssigneeRepository(pdpas.Ctx, pdpas.Db)

	if err := repo.Find(pdpas.ProjectAssignee); err != nil {
		return exceptions.NewBadRequestError("Invalid Project Assignee")
	}

	issueAssigneeRepo := repository.NewIssueAssigneeRepository(pdpas.Ctx, pdpas.Db)

	if countIssueAssignee := issueAssigneeRepo.CountByProjectAssignee(*pdpas.ProjectAssignee); countIssueAssignee > 0 {
		return exceptions.NewBadRequestError("This member cant be delete because some issues in this project might has assignment on him")
	}

	if err := repo.Delete(pdpas.ProjectAssignee); err != nil {
		return err
	}

	return nil
}
