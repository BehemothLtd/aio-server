package insightResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func (r *Resolver) SelfWorkingTimeLogDestroy(ctx context.Context, args insightInputs.WorkingTimelogInput) (*string, error) {
	user, err := auths.AuthInsightUserFromCtx(ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	if args.Id == "" {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}

	workingTimeLogId, err := helpers.GqlIdToInt32(args.Id)

	if err != nil {
		return nil, err
	}

	workingTimelog := models.WorkingTimelog{}

	repo := repository.NewWorkingTimelogRepository(&ctx, r.Db)

	if err := repo.FindById(&workingTimelog, workingTimeLogId); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exceptions.NewRecordNotFoundError()
		}
		return nil, err
	}

	if workingTimelog.UserId != user.Id {
		return nil, exceptions.NewBadRequestError("Cannot delete.This timelog is not yours.")
	}

	if !workingTimelog.CreatedAt.After(time.Now().Add(-2 * time.Hour)) {
		return nil, exceptions.NewBadRequestError("It's overtime to delete this timelog.")
	}

	if err = repo.Destroy(&workingTimelog); err != nil {
		return nil, exceptions.NewBadRequestError(fmt.Sprintf("Can not delete this working timelog %s", err.Error()))
	} else {
		message := "Deleted"

		return &message, nil
	}
}
