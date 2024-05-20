package insightResolvers

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"

	"github.com/graph-gophers/graphql-go"
)

func (r *Resolver) ProjectIssue(ctx context.Context, args struct {
	ProjectId graphql.ID
	Id        graphql.ID
}) (*globalTypes.IssueType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeProjects.String(), enums.PermissionActionTypeRead.String()); err != nil {
		return nil, err
	}

	if args.ProjectId == "" || args.Id == "" {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}

	projectId, err := helpers.GqlIdToInt32(args.ProjectId)
	if err != nil || projectId == 0 {
		return nil, exceptions.NewBadRequestError("Invalid Project")
	}

	id, err := helpers.GqlIdToInt32(args.Id)
	if err != nil || id == 0 {
		return nil, exceptions.NewBadRequestError("Invalid Issue")
	}

	project := models.Project{Id: projectId}
	projectRepo := repository.NewProjectRepository(&ctx, r.Db)

	if err := projectRepo.Find(&project); err != nil {
		return nil, exceptions.NewRecordNotFoundError()
	}

	issue := models.Issue{Id: id, ProjectId: projectId}
	issueRepo := repository.NewIssueRepository(
		&ctx,
		r.Db.Preload("IssueAssignees.User.Avatar.AttachmentBlob").
			Preload("Creator.Avatar.AttachmentBlob").
			Preload("IssueStatus"),
	)

	if err := issueRepo.Find(&issue); err != nil {
		return nil, exceptions.NewRecordNotFoundError()
	}

	return &globalTypes.IssueType{
		Issue: &issue,
	}, nil
}
