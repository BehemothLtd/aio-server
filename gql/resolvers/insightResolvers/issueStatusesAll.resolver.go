package insightResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/repository"
	"context"
)

func (r *Resolver) IssueStatusesAll(ctx context.Context) (*[]*globalTypes.IssueStatusType, error) {
	_, err := auths.AuthInsightUserFromCtx(ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	issueStatuses := []*models.IssueStatus{}
	repo := repository.NewIssueStatusRepository(&ctx, r.Db)

	if err := repo.All(&issueStatuses); err != nil {
		return nil, err
	} else {
		result := make([]*globalTypes.IssueStatusType, len(issueStatuses))

		for i, issueStatus := range issueStatuses {
			result[i] = &globalTypes.IssueStatusType{
				IssueStatus: issueStatus,
			}
		}
		return &result, nil
	}
}
