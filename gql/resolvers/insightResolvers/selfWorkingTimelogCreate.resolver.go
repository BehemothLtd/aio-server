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

func (r *Resolver) MmSelfWorkingTimelogCreate(ctx context.Context, args insightInputs.SelfWorkingTimelogCreateInput) (*insightTypes.WorkingtimelogCreateReturnType, error) {
	user, err := auths.AuthUserFromCtx(ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	workingTimelog := models.WorkingTimelog{}

	service := insightServices.SelfCreateWorkingTimelogService{
		Ctx:            &ctx,
		Db:             r.Db,
		Args:           *args.Input,
		User:           &user,
		IssueId:        args.IssueId,
		WorkingTimelog: &workingTimelog,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	}

	return &insightTypes.WorkingtimelogCreateReturnType{
		WorkingTimelog: &globalTypes.WorkingTimelogType{
			WorkingTimelog: &workingTimelog,
		},
	}, nil
}
