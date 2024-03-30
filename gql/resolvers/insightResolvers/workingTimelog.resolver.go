package insightResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"

	"gorm.io/gorm"
)

func (r *Resolver) MmWorkingTimelog(ctx context.Context, args insightInputs.WorkingTimelogInput) (*globalTypes.WorkingTimelogType, error) {
	_, err := auths.AuthUserFromCtx(ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	if args.Id == "" {
		return nil, exceptions.NewBadRequestError(("Invalid ID"))
	}

	workingTimeLogId, err := helpers.GqlIdToInt32(args.Id)

	if err != nil {
		return nil, err
	}

	workingTimelog := models.WorkingTimelog{}
	repo := repository.NewWorkingTimelogRepository(&ctx, r.Db.Preload("User").Preload("Project").Preload("Issue"))
	err = repo.FindById(&workingTimelog, workingTimeLogId)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exceptions.NewRecordNotFoundError()
		}
		return nil, err
	}

	return &globalTypes.WorkingTimelogType{WorkingTimelog: &workingTimelog}, nil
}
