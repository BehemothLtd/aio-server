package insightResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"

	"gorm.io/gorm"
)

func (r *Resolver) MmWorkingTimelog(ctx context.Context, args insightInputs.WorkingTimelogInput) (*insightTypes.WorkingTimelogType, error) {
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

	return &insightTypes.WorkingTimelogType{WorkingTimelog: &workingTimelog}, nil
}
