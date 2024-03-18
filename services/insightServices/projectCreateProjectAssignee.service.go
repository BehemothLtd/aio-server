package insightServices

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"aio-server/validators"
	"fmt"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type ProjectCreateProjectAssigneeService struct {
	Ctx             *context.Context
	Db              *gorm.DB
	Args            insightInputs.ProjectModifyProjectAssigneeInput
	Project         *models.Project
	ProjectAssignee *models.ProjectAssignee
}

func (pcpas *ProjectCreateProjectAssigneeService) Execute() error {
	if pcpas.Args.Id == "" {
		return exceptions.NewBadRequestError("Invalid Project ID")
	}

	projectId, err := helpers.GqlIdToInt32(pcpas.Args.Id)

	if err != nil || projectId == 0 {
		return exceptions.NewBadRequestError("Invalid Project")
	}

	pcpas.Project = &models.Project{Id: projectId}

	projectRepo := repository.NewProjectRepository(pcpas.Ctx, pcpas.Db)
	if err := projectRepo.Find(pcpas.Project); err != nil {
		return exceptions.NewBadRequestError("Invalid Project")
	}

	pcpas.ProjectAssignee = &models.ProjectAssignee{ProjectId: pcpas.Project.Id}

	form := validators.NewProjectAssigneeFormValidator(
		pcpas.Args.Input,
		repository.NewProjectAssigneeRepository(pcpas.Ctx, pcpas.Db),
		*pcpas.Project,
		pcpas.ProjectAssignee,
	)

	if err := form.Save(); err != nil {
		fmt.Println("ERROR", err)
		return err
	}

	return nil
}
