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

type ProjectUpdateIssueService struct {
	Ctx   *context.Context
	Db    *gorm.DB
	Args  insightInputs.ProjectUpdateIssueInput
	Issue *models.Issue
}

func (puis *ProjectUpdateIssueService) Execute() error {
	if puis.Args.ProjectId == "" {
		return exceptions.NewBadRequestError("Invalid Project ID")
	}

	projectId, err := helpers.GqlIdToInt32(puis.Args.ProjectId)

	if err != nil || projectId == 0 {
		return exceptions.NewBadRequestError("Invalid Project")
	}

	project := models.Project{Id: projectId}
	projectRepo := repository.NewProjectRepository(puis.Ctx, puis.Db.Preload("ProjectIssueStatuses").Preload("Issues").Preload("ProjectSprints"))

	if err := projectRepo.Find(&project); err != nil {
		return exceptions.NewBadRequestError("Invalid Project")
	}

	id, err := helpers.GqlIdToInt32(puis.Args.Id)

	if err != nil || id == 0 {
		return exceptions.NewBadRequestError("Invalid Issue")
	}

	puis.Issue.Id = id
	puis.Issue.ProjectId = projectId

	repo := repository.NewIssueRepository(nil, puis.Db)
	if err := repo.Find(puis.Issue); err != nil {
		return exceptions.NewBadRequestError("Invalid Issue")
	}

	return nil
}
