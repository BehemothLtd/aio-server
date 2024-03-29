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

type ProjectUpdateService struct {
	Ctx     *context.Context
	Db      *gorm.DB
	Args    insightInputs.ProjectUpdateInput
	Project *models.Project
}

func (pus *ProjectUpdateService) Execute() error {
	repo := repository.NewProjectRepository(
		pus.Ctx,
		pus.Db.Preload("ProjectIssueStatuses.IssueStatus").Preload("ProjectAssignees"),
	)

	if err := repo.Find(pus.Project); err != nil {
		return exceptions.NewRecordNotFoundError()
	}

	form := validators.NewProjectUpdateFormValidator(
		&pus.Args.Input,
		repo,
		pus.Project,
	)

	if err := form.Save(); err != nil {
		return err
	}

	return nil
}
