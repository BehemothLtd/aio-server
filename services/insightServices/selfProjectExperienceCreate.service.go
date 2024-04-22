package insightServices

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"aio-server/validators"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type SelfProjectExperienceCreateService struct {
	Ctx               *context.Context
	Db                *gorm.DB
	Args              insightInputs.ProjectExperienceCreateInput
	ProjectExperience *models.ProjectExperience
}

func (specs *SelfProjectExperienceCreateService) Execute() error {
	form := validators.NewProjectExperienceFormValidator(
		&specs.Args.Input,
		repository.NewProjectExperienceRepository(specs.Ctx, specs.Db),
		specs.ProjectExperience,
	)

	if err := form.Save(); err != nil {
		return err
	}

	return nil
}
