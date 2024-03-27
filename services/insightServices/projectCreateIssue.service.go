package insightServices

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"aio-server/validators"
	"context"

	"gorm.io/gorm"
)

type ProjectCreateIssueService struct {
	Ctx   *context.Context
	Db    *gorm.DB
	Args  insightInputs.ProjectCreateIssueInput
	Issue *models.Issue
}

func (pcis *ProjectCreateIssueService) Execute() error {
	if pcis.Args.Id == "" {
		return exceptions.NewBadRequestError("Invalid Project ID")
	}

	projectId, err := helpers.GqlIdToInt32(pcis.Args.Id)

	if err != nil || projectId == 0 {
		return exceptions.NewBadRequestError("Invalid Project")
	}

	project := models.Project{Id: projectId}
	projectRepo := repository.NewProjectRepository(pcis.Ctx, pcis.Db.Preload("ProjectIssueStatuses").Preload("Issues").Preload("ProjectSprints"))

	if err := projectRepo.Find(&project); err != nil {
		return exceptions.NewBadRequestError("Invalid Project")
	}

	pcis.Issue.ProjectId = project.Id

	form := validators.NewProjectModifyIssueFormValidator(
		&pcis.Args.Input,
		*repository.NewIssueRepository(pcis.Ctx, pcis.Db),
		project,
		pcis.Issue,
	)

	if err := form.Save(); err != nil {
		return err
	}

	return nil
}
