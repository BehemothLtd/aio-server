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

type IssueStatusUpdateService struct {
	Ctx         *context.Context
	Db          *gorm.DB
	Args        insightInputs.IssueStatusUpdateInput
	IssueStatus *models.IssueStatus
}

func (isus *IssueStatusUpdateService) Execute() error {
	repo := repository.NewIssueStatusRepository(isus.Ctx, isus.Db)

	if err := repo.Find(isus.IssueStatus); err != nil {
		return exceptions.NewRecordNotFoundError()
	}

	form := validators.NewIssueStatusUpdateFormValidator(
		&isus.Args.Input,
		repo,
		isus.IssueStatus,
	)

	if err := form.Save(); err != nil {
		return err
	}

	return nil
}
