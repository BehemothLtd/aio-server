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

func (r *Resolver) MmWorkingTimelogs(ctx context.Context, args insightInputs.WorkingTimelogsInput) (*insightTypes.WorkingTimelogsType, error) {
	_, authErr := auths.AuthUserFromCtx(ctx)

	if authErr != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	var workingTimelogs []*models.WorkingTimelog

	workingTimelogQuery, paginationData := args.ToPaginationDataAndQuery()

	repo := repository.NewWorkingTimelogRepository(&ctx, r.Db.Preload("User").Preload("Project").Preload("Issue"))

	err := repo.List(&workingTimelogs, &paginationData, workingTimelogQuery)
	if err != nil {
		return nil, err
	}

	result := make([]*globalTypes.WorkingTimelogType, len(workingTimelogs))

	for i, workingTimelog := range workingTimelogs {
		result[i] = &globalTypes.WorkingTimelogType{
			WorkingTimelog: workingTimelog,
		}
	}

	return &insightTypes.WorkingTimelogsType{
		Collection: &result,
		Metadata: &globalTypes.MetadataType{
			Metadata: &paginationData.Metadata,
		},
	}, nil
}
