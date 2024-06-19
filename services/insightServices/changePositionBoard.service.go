package insightServices

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"

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
	minPosition, maxPosition, count, err := issueRepo.FindMinAndMaxPositionWithCount(issueStatus)
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	newPosition := cpbs.calculateIssuesPosition(*minPosition, *maxPosition, *count, int(arg.NewIndex))

	listAffectedIssues, err := issueRepo.FindIssueOfProjectByStatusAndPosition(issueStatus, newPosition)
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	if err := cpbs.changeIssuesPosition(issue, int32(newPosition), arg.NewStatusId, listAffectedIssues); err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	return nil
}

func (cpbs *ChangePositionBoardService) calculateIssuesPosition(minPosition int, maxPosition int, count int, newIndex int) int {
	var position int
	switch count {
	// Case: Project Issue Status has no issue
	case 0:
		position = minPosition
	// Case: When move issue to bottom
	case newIndex:
		position = maxPosition + 1
	default:
		position = newIndex
	}
	return position
}

func (cpbs *ChangePositionBoardService) changeIssuesPosition(issue models.Issue, newPosition int32, newIssueStatusID int32, listAffectedIssues []models.Issue) error {
	issueRepo := repository.NewIssueRepository(cpbs.Ctx, cpbs.Db)
	if err := cpbs.Db.Transaction(func(tx *gorm.DB) error {
		for _, issue := range listAffectedIssues {
			if err := issueRepo.UpdatePosition(&issue, issue.Position+1, newIssueStatusID); err != nil {
				return err
			}
		}

		if err := issueRepo.UpdatePosition(&issue, newPosition, newIssueStatusID); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
