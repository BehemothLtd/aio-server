package insightResolvers

import (
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/pkg/constants"
	"aio-server/repository"
	"context"
)

func (r *Resolver) SelfThisWeekIssuesState(ctx context.Context) (*insightTypes.SelfThisWeekIssuesStateType, error) {
	if user, err := auths.AuthInsightUserFromCtx(ctx); err != nil {
		return nil, err
	} else {
		labels := []string{}
		series := insightTypes.SelfThisWeekIssuesStateSeriesType{}

		issueDateBaseState := []models.IssuesDeadlineBaseState{}
		repo := repository.NewIssueRepository(&ctx, r.Db)
		if err := repo.UserAllWeekIssuesState(user, &issueDateBaseState); err != nil {
			return nil, err
		}

		for _, issueDateBaseData := range issueDateBaseState {
			labels = append(labels, issueDateBaseData.Date.Format(constants.MMDD_DateFormatForChart))
			series.Done = append(series.Done, issueDateBaseData.Done)
			series.NotDone = append(series.Done, issueDateBaseData.NotDone)
		}

		return &insightTypes.SelfThisWeekIssuesStateType{
			Labels: labels,
			Series: series,
		}, nil
	}
}
