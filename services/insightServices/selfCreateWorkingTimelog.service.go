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

type SelfCreateWorkingTimelogService struct {
	Ctx            *context.Context
	Db             *gorm.DB
	Args           insightInputs.SelfWorkingTimelogFormInput
	User           *models.User
	WorkingTimelog *models.WorkingTimelog
	IssueId        *int32
}

func (scwts *SelfCreateWorkingTimelogService) Execute() error {
	IssueAssignee := models.IssueAssignee{}

	checkIssueError := scwts.Db.Preload("Issue").Model(&models.IssueAssignee{}).Where("issue_id = ? AND user_id = ?", scwts.IssueId, scwts.User.Id).First(&IssueAssignee).Error

	if checkIssueError != nil {
		return exceptions.NewUnprocessableContentError("You have no permission for this issue", nil)
	}
	issue := IssueAssignee.Issue

	form := validators.NewWorkingTimelogFormValidator(
		&scwts.Args,
		scwts.User,
		repository.NewWorkingTimelogRepository(scwts.Ctx, scwts.Db),
		scwts.WorkingTimelog,
		&issue,
	)

	if err := form.Create(); err != nil {
		return err
	}

	return nil
}
