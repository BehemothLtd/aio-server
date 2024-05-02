package insightServices

import (
	"aio-server/exceptions"
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"fmt"
	"slices"

	"github.com/graph-gophers/graphql-go"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type ProjectDeleteProjectIssueStatusService struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args struct {
		ProjectId graphql.ID
		Id        graphql.ID
	}
}

func (service *ProjectDeleteProjectIssueStatusService) Execute() error {
	if service.Args.ProjectId == "" {
		return exceptions.NewBadRequestError("Invalid Project Id")
	}

	if service.Args.Id == "" {
		return exceptions.NewBadRequestError("Invalid Id")
	}

	projectId, err := helpers.GqlIdToInt32(service.Args.ProjectId)

	if err != nil || projectId == 0 {
		return exceptions.NewBadRequestError("Invalid Project")
	}

	projectIssueStatusId, err := helpers.GqlIdToInt32(service.Args.Id)

	if err != nil || projectIssueStatusId == 0 {
		return exceptions.NewBadRequestError("Invalid Project Issue Status ID")
	}

	project := models.Project{Id: projectId}
	projectRepo := repository.NewProjectRepository(service.Ctx, service.Db)

	if err := projectRepo.Find(&project); err != nil {
		return exceptions.NewRecordNotFoundError()
	}

	projectIssueStatus := models.ProjectIssueStatus{Id: projectIssueStatusId, ProjectId: projectId}
	repo := repository.NewProjectIssueStatusRepository(service.Ctx, service.Db)

	if err := repo.Find(&projectIssueStatus); err != nil {
		return exceptions.NewRecordNotFoundError()
	}

	var requiredIssueStatusIds []int32

	if project.IsKanban() {
		requiredIssueStatusIds = constants.RequiredIssueStatusIdsForKanbanProject()
	} else {
		requiredIssueStatusIds = constants.RequiredIssueStatusIdsForScrumProject()
	}

	if slices.Contains(requiredIssueStatusIds, projectIssueStatus.IssueStatusId) {
		return exceptions.NewUnprocessableContentError("Cant delete. This status is required for this kind of project", nil)
	}

	var thisStatusIssueCount int32
	issueRepo := repository.NewIssueRepository(service.Ctx, service.Db)

	if err := issueRepo.CountByProjectAndIssueStatus(projectId, projectIssueStatusId, &thisStatusIssueCount); err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	if thisStatusIssueCount > 0 {
		return exceptions.NewUnprocessableContentError("cant delete this status since there are issues using this status", nil)
	}

	if err := repo.Delete(projectId, projectIssueStatusId); err != nil {
		return exceptions.NewUnprocessableContentError(fmt.Sprintf("cant delete this status, error: %s", err.Error()), nil)
	}

	return nil
}
