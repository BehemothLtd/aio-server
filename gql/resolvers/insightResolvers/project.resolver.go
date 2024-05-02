package insightResolvers

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

func (r *Resolver) Project(ctx context.Context, args struct{ Id graphql.ID }) (*globalTypes.ProjectType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeProjects.String(), enums.PermissionActionTypeRead.String()); err != nil {
		return nil, err
	}

	projectId, err := helpers.GqlIdToInt32(args.Id)

	if err != nil {
		return nil, exceptions.NewBadRequestError(err.Error())
	}

	project := models.Project{Id: projectId}

	repo := repository.NewProjectRepository(
		&ctx,
		r.Db.Preload("Client").
			Preload("ProjectIssueStatuses", func(db *gorm.DB) *gorm.DB {
				return db.Order("project_issue_statuses.position ASC")
			}).
			Preload("ProjectIssueStatuses.IssueStatus").
			Preload("ProjectAssignees.User").
			Preload("ProjectAssignees.User.Avatar.AttachmentBlob").
			Preload("Logo", "name = 'logo'").Preload("Logo.AttachmentBlob").
			Preload("Files", "name = 'files'").Preload("Files.AttachmentBlob"),
	)

	if err := repo.Find(&project); err != nil {
		return nil, exceptions.NewRecordNotFoundError()
	}

	return &globalTypes.ProjectType{Project: &project}, nil
}
