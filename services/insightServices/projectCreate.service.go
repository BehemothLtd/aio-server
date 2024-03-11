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

func (pcs *ProjectCreateService) Execute() (*models.Project, error) {
	project := models.Project{}

	form := validators.NewProjectCreateFormValidator(
		&pcs.Args.Input,
		repository.NewProjectRepository(pcs.Ctx, pcs.Db),
		&project,
	)

	if err := form.Save(); err != nil {
		return nil, err
	}

	repo := repository.NewProjectRepository(pcs.Ctx, pcs.Db.Preload("ProjectIssueStatuses"))
	createdProject := models.Project{Id: project.Id}
	repo.Find(&createdProject)

	return &project, nil
}
