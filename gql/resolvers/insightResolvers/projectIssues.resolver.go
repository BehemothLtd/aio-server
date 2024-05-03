package insightResolvers

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"
)

func (r *Resolver) ProjectIssues(ctx context.Context, args insightInputs.ProjectIssuesInput) (*insightTypes.IssuesType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeProjects.String(), enums.PermissionActionTypeRead.String()); err != nil {
		return nil, err
	}

	if args.Id == "" {
		return nil, exceptions.NewBadRequestError("Invalid Project Id")
	}

	projectId, err := helpers.GqlIdToInt32(args.Id)

	if err != nil || projectId == 0 {
		return nil, exceptions.NewBadRequestError("Invalid Project")
	}

	project := models.Project{Id: projectId}
	projectRepo := repository.NewProjectRepository(&ctx, r.Db)

	if err := projectRepo.Find(&project); err != nil {
		return nil, exceptions.NewRecordNotFoundError()
	}

	var issues []*models.Issue
	issuesQuery, paginationData := args.ToPaginationDataAndQuery()
	issuesQuery.ProjectIdEq = &projectId

	repo := repository.NewIssueRepository(
		&ctx,
		r.Db.Preload("IssueAssignees.User.Avatar.AttachmentBlob").
			Preload("Creator.Avatar.AttachmentBlob").
			Preload("IssueStatus"),
	)

	if err := repo.List(&issues, issuesQuery, &paginationData); err != nil {
		return nil, exceptions.NewBadRequestError(err.Error())
	}

	return &insightTypes.IssuesType{
		Collection: r.IssuesSliceToTypes(issues),
		Metadata: &globalTypes.MetadataType{
			Metadata: &paginationData.Metadata,
		},
	}, nil
}
