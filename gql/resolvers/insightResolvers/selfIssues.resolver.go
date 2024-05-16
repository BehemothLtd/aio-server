package insightResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/repository"
	"context"
)

func (r *Resolver) SelfIssues(ctx context.Context, args insightInputs.IssuesInput) (*insightTypes.IssuesType, error) {
	user, err := auths.AuthInsightUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	var issues []*models.Issue
	issuesQuery, paginationData := args.ToPaginationDataAndQuery()

	repo := repository.NewIssueRepository(
		&ctx,
		r.Db.Preload("IssueAssignees.User.Avatar.AttachmentBlob").
			Preload("Creator.Avatar.AttachmentBlob").
			Preload("IssueStatus"))
	if err := repo.ListByUser(&issues, issuesQuery, &paginationData, user); err != nil {
		return nil, exceptions.NewBadRequestError(err.Error())
	}

	return &insightTypes.IssuesType{
		Collection: r.IssuesSliceToTypes(issues),
		Metadata: &globalTypes.MetadataType{
			Metadata: &paginationData.Metadata,
		},
	}, nil
}
