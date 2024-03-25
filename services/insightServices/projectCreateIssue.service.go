package insightServices

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
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
	return nil
}
