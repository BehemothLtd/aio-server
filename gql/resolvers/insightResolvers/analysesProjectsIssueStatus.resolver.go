package insightResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/pkg/auths"
	"aio-server/services/insightServices"
	"context"
)

func (r *Resolver) AnalysesProjectsIssueStatus(ctx context.Context) (*insightTypes.AnalysesProjectsIssueStatus, error) {
	if _, err := auths.AuthInsightUserFromCtx(ctx); err != nil {
		return nil, err
	}

	data := insightTypes.AnalysesProjectsIssueStatus{}
	service := insightServices.AnalysesProjectsIssueStatusService{
		Ctx:  ctx,
		Db:   *r.Db,
		Data: &data,
	}

	if err := service.Execute(); err != nil {
		return nil, exceptions.NewBadRequestError(err.Error())
	}

	return &insightTypes.AnalysesProjectsIssueStatus{
		Categories: service.Data.Categories,
		Series:     service.Data.Series,
	}, nil
}
