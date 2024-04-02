package insightResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/services/insightServices"
	"context"
)

func (r *Resolver) MmSelfWorkingTimelogCreateOrUpdate(ctx context.Context, args insightInputs.SelfWorkingTimelogCreateInput) (*insightTypes.WorkingtimelogMutationType, error) {
	user, err := auths.AuthInsightUserFromCtx(ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	workingTimelog := models.WorkingTimelog{}

	service := insightServices.SelfCreateOrUpdateWorkingTimelogService{
		Ctx:            &ctx,
		Db:             r.Db,
		Args:           args,
		User:           &user,
		WorkingTimelog: &workingTimelog,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	}

	return &insightTypes.WorkingtimelogMutationType{
		WorkingTimelog: &globalTypes.WorkingTimelogType{
			WorkingTimelog: &workingTimelog,
		},
	}, nil
}
