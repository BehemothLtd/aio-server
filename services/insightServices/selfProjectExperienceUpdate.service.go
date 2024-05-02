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

type SelfProjectExperienceUpdateService struct {
	Ctx               *context.Context
	Db                *gorm.DB
	Args              insightInputs.ProjectExperienceUpdateInput
	ProjectExperience *models.ProjectExperience
}

func (speus *SelfProjectExperienceUpdateService) Execute() error {
	repo := repository.NewProjectExperienceRepository(speus.Ctx, speus.Db)

	if err := repo.Find(speus.ProjectExperience); err != nil {
		return exceptions.NewRecordNotFoundError()
	}

	form := validators.NewProjectExperienceFormValidator(
		&speus.Args.Input,
		repository.NewProjectExperienceRepository(speus.Ctx, speus.Db),
		speus.ProjectExperience,
	)

	if err := form.Save(); err != nil {
		return err
	}

	return nil
}
