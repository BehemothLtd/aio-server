package insightResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/pkg/helpers"
	"aio-server/services/insightServices"
	"context"
)

func (r *Resolver) MmSelfWorkingTimelogUpdate(ctx context.Context, args insightInputs.SelfWorkingTimelogUpdateInput) (*insightTypes.WorkingtimelogMutationType, error) {
	user, err := auths.AuthUserFromCtx(ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	workingTimeLogId, idError := helpers.GqlIdToInt32(args.Id)

	if idError != nil {
		return nil, exceptions.NewRecordNotFoundError()
	}

	workingTimelog := models.WorkingTimelog{Id: workingTimeLogId, UserId: user.Id}

	service := insightServices.SelfUpdateWorkingTimelogService{
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

	return &insightTypes.WorkingtimelogMutationType{
		WorkingTimelog: &globalTypes.WorkingTimelogType{
			WorkingTimelog: &workingTimelog,
		},
	}, nil
}
