package insightResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"
)

func (r *Resolver) MmSelfWorkingTimelogDelete(ctx context.Context, args insightInputs.SelfWorkingTimelogDeleteInput) (*string, error) {
	user, err := auths.AuthUserFromCtx(ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	workingTimeLogId, idError := helpers.GqlIdToInt32(args.Id)

	if idError != nil {
		return nil, exceptions.NewRecordNotFoundError()
	}

	workingTimelog := models.WorkingTimelog{Id: workingTimeLogId, UserId: user.Id, IssueId: *args.IssueId}

	repo := repository.NewWorkingTimelogRepository(&ctx, r.Db)

	err = repo.Delete(&workingTimelog)

	if err != nil {
		return nil, exceptions.NewRecordNotFoundError()
	} else {
		successMessage := "Deleted"
		return &successMessage, nil
	}

}
