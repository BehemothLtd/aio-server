package insightServices

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"aio-server/validators"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type ProjectSprintCreateService struct {
	Ctx           *context.Context
	Db            *gorm.DB
	Args          insightInputs.ProjectSprintCreateInput
	ProjectSprint *models.ProjectSprint
}

func (pscs *ProjectSprintCreateService) Execute() error {

	form := validators.NewProjectSprintFormValidator(
		&pscs.Args.Input,
		repository.NewProjectSprintRepository(pscs.Ctx, pscs.Db),
		pscs.ProjectSprint,
	)

	if err := form.Save(); err != nil {
		return err
	}

	return nil
}
