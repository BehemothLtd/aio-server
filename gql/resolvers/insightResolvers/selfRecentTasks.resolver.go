package insightResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/repository"
	"context"
)

func (r *Resolver) SelfRecentTasks(ctx context.Context) ([]*globalTypes.IssueType, error) {
	user, err := auths.AuthInsightUserFromCtx(ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	recentTasks := []models.Issue{}

	repo := repository.NewIssueRepository(&ctx, r.Db)
	repo.FindRecentTasksByUser(&recentTasks, user.Id)

	result := make([]*globalTypes.IssueType, len(recentTasks))

	for i, task := range recentTasks {
		result[i] = &globalTypes.IssueType{
			Issue: &task,
		}
	}

	return result, nil
}
