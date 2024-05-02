package insightServices

import (
	"aio-server/exceptions"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"fmt"
	"slices"

	"github.com/graph-gophers/graphql-go"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type ProjectCreateProjectIssueStatusService struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args struct {
		ProjectId graphql.ID
		Id        graphql.ID
	}
}

func (service *ProjectCreateProjectIssueStatusService) Execute() error {
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

	issueStatusId, err := helpers.GqlIdToInt32(service.Args.Id)

	if err != nil || issueStatusId == 0 {
		return exceptions.NewBadRequestError("Invalid Issue Status ID")
	}

	project := models.Project{Id: projectId}
	projectRepo := repository.NewProjectRepository(service.Ctx, service.Db)

	if err := projectRepo.Find(&project); err != nil {
		return exceptions.NewRecordNotFoundError()
	}

	issueStatus := models.IssueStatus{Id: issueStatusId}
	repo := repository.NewIssueStatusRepository(service.Ctx, service.Db)

	if err := repo.Find(&issueStatus); err != nil {
		return exceptions.NewRecordNotFoundError()
	}

	var currentProjectStatusIds []int32
	if err := repo.FetchIdsOnProject(projectId, &currentProjectStatusIds); err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	if slices.Contains(currentProjectStatusIds, issueStatusId) {
		return exceptions.NewUnprocessableContentError("This status is already existed", nil)
	}

	projectIssueStatusRepo := repository.NewProjectIssueStatusRepository(service.Ctx, service.Db)

	if err := projectIssueStatusRepo.Create(projectId, issueStatusId); err != nil {
		return exceptions.NewUnprocessableContentError(fmt.Sprintf("Error happened %s", err.Error()), nil)
	}

	return nil
}
