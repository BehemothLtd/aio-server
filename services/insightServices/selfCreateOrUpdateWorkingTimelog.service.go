package insightServices

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"aio-server/validators"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type SelfCreateOrUpdateWorkingTimelogService struct {
	Ctx            *context.Context
	Db             *gorm.DB
	Args           insightInputs.SelfWorkingTimelogCreateInput
	User           *models.User
	WorkingTimelog *models.WorkingTimelog
}

func (scwts *SelfCreateOrUpdateWorkingTimelogService) Execute() error {
	issueAssignee := models.IssueAssignee{IssueId: *scwts.Args.IssueId, UserId: scwts.User.Id}
	issueAssigneeRepo := repository.NewIssueAssigneeRepository(scwts.Ctx, scwts.Db.Preload("Issue"))

	checkIssueError := issueAssigneeRepo.FindByAttr(&issueAssignee)

	if checkIssueError != nil {
		return exceptions.NewUnprocessableContentError("You have no permission for this issue", nil)
	}
	issue := issueAssignee.Issue

	form := validators.NewWorkingTimelogFormValidator(
		scwts.Args.Input,
		scwts.User,
		repository.NewWorkingTimelogRepository(scwts.Ctx, scwts.Db),
		scwts.WorkingTimelog,
		&issue,
	)

	if err := form.Save(); err != nil {
		return err
	}

	return nil
}
