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

type ProjectSprintUpdateService struct {
	Ctx           *context.Context
	Db            *gorm.DB
	Args          insightInputs.ProjectSprintUpdateInput
	ProjectSprint *models.ProjectSprint
}

func (psus *ProjectSprintUpdateService) Execute() error {
	repo := repository.NewProjectSprintRepository(psus.Ctx, psus.Db)

	if err := repo.Find(psus.ProjectSprint); err != nil {
		return exceptions.NewRecordNotFoundError()
	}

	form := validators.NewProjectSprintFormValidator(
		&psus.Args.Input,
		repo,
		psus.ProjectSprint,
	)

	if err := form.Save(); err != nil {
		return err
	}

	return nil
}
