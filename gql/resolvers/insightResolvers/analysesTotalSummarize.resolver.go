package insightResolvers

import (
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/pkg/auths"
	"context"
)

func (r *Resolver) AnalysesTotalSummarize(ctx context.Context) (*insightTypes.AnalysesTotalSummarize, error) {
	if _, err := auths.AuthInsightUserFromCtx(ctx); err != nil {
		return nil, err
	}

	var issueCount int32
	var memberCount int32
	var projectCount int32
	var workingTimeCount int32

	r.Db.Table("issues").Select("Count(*)").Scan(&issueCount)
	r.Db.Table("users").Select("Count(*)").Scan(&memberCount)
	r.Db.Table("projects").Select("Count(*)").Scan(&projectCount)
	r.Db.Table("working_timelogs").Select("SUM(minutes)").Scan(&workingTimeCount)

	return &insightTypes.AnalysesTotalSummarize{
		IssueCount:       issueCount,
		MemberCount:      memberCount,
		ProjectCount:     projectCount,
		WorkingTimeCount: workingTimeCount,
	}, nil
}
