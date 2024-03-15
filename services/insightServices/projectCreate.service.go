package insightServices

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"aio-server/validators"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type ProjectCreateService struct {
	Ctx     *context.Context
	Db      *gorm.DB
	Args    insightInputs.ProjectCreateInput
	Project *models.Project
}

func (pcs *ProjectCreateService) Execute() error {
	form := validators.NewProjectCreateFormValidator(
		&pcs.Args.Input,
		repository.NewProjectRepository(pcs.Ctx, pcs.Db),
		pcs.Project,
	)

	if err := form.Save(); err != nil {
		return err
	}

	return nil
}
