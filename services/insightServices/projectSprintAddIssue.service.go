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

type ProjectSprintAddIssueService struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args insightInputs.ProjectSprintIssueModifyInput
}

func (service *ProjectSprintAddIssueService) Execute() error {
	if service.Args.ProjectId == "" || service.Args.Id == "" || service.Args.IssueId == "" {
		return exceptions.NewBadRequestError("Invalid ID")
	}

	projectId, err := helpers.GqlIdToInt32(service.Args.ProjectId)
	if err != nil || projectId == 0 {
		return exceptions.NewBadRequestError("Invalid Id")
	}

	sprintId, err := helpers.GqlIdToInt32(service.Args.Id)
	if err != nil || sprintId == 0 {
		return exceptions.NewBadRequestError("Invalid Id")
	}

	issueId, err := helpers.GqlIdToInt32(service.Args.IssueId)
	if err != nil || issueId == 0 {
		return exceptions.NewBadRequestError("Invalid Id")
	}

	project := models.Project{Id: projectId}
	projectRepo := repository.NewProjectRepository(service.Ctx, service.Db)
	if err := projectRepo.Find(&project); err != nil {
		return exceptions.NewBadRequestError("Invalid Project")
	}

	sprint := models.ProjectSprint{Id: sprintId, ProjectId: projectId}
	sprintRepo := repository.NewProjectSprintRepository(service.Ctx, service.Db)
	if err := sprintRepo.Find(&sprint); err != nil {
		return exceptions.NewBadRequestError("Invalid Sprint")
	}

	issue := models.Issue{Id: issueId, ProjectId: projectId}
	issueRepo := repository.NewIssueRepository(service.Ctx, service.Db)
	if err := issueRepo.Find(&issue); err != nil {
		return exceptions.NewBadRequestError("Invalid Issue")
	}

	if err := issueRepo.UpdateSprint(&issue, sprint.Id); err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	return nil
}
