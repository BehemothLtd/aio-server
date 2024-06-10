package insightServices

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"

	"github.com/davecgh/go-spew/spew"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type ChangePositionBoardService struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args insightInputs.ProjectBoardChangePositionInput
}

func (cpbs *ChangePositionBoardService) Execute() error {
	arg := cpbs.Args

	project := models.Project{Id: arg.ProjectId}
	projectRepo := repository.NewProjectRepository(cpbs.Ctx, cpbs.Db.Preload("ProjectIssueStatuses"))
	if err := projectRepo.Find(&project); err != nil {
		return exceptions.NewBadRequestError("Invalid Project")
	}

	issue := models.Issue{Id: arg.Id, ProjectId: arg.ProjectId}
	issueRepo := repository.NewIssueRepository(cpbs.Ctx, cpbs.Db)
	if err := issueRepo.Find(&issue); err != nil {
		return exceptions.NewBadRequestError("Invalid Issue")
	}

	var issueStatus *models.ProjectIssueStatus
	for _, PIS := range project.ProjectIssueStatuses {
		if PIS.IssueStatusId == arg.NewStatusId {
			issueStatus = PIS
			break
		}
	}
	if issueStatus == nil {
		return exceptions.NewBadRequestError("Invalid New Issue Status")
	}
	projectIssueStatusRepo := repository.NewProjectIssueStatusRepository(cpbs.Ctx, cpbs.Db)
	positionStats, err := projectIssueStatusRepo.FindMinAndMaxPositionWithCount(issueStatus)
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}
	spew.Dump(positionStats)

	return nil
}
