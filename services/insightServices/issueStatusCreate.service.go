package insightServices

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"aio-server/validators"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type IssueStatusCreateService struct {
	Ctx         *context.Context
	Db          *gorm.DB
	Args        insightInputs.IssueStatusFormInput
	IssueStatus *models.IssueStatus
}

func (iscs *IssueStatusCreateService) Execute() error {
	// validator
	form := validators.NewIssueStatusFormValidator(
		&iscs.Args,
		repository.NewIssueStatusRepository(iscs.Ctx, iscs.Db),
		iscs.IssueStatus,
	)

	if err := form.Save(); err != nil {
		return err
	}

	return nil
}
