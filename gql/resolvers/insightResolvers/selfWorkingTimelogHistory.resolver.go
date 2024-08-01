package insightResolvers

import (
	"aio-server/enums"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"context"
)

func (r *Resolver) SelfWorkingTimelogHistory(ctx context.Context, args insightInputs.SelfWorkingTimelogHistoryInput) (*[]*globalTypes.WorkingTimelogHistoryType, error) {
	currentUser, err := r.Authorize(ctx, enums.PermissionTargetTypeClients.String(), enums.PermissionActionTypeWrite.String())
	if err != nil {
		return nil, err
	}

	var history []*models.WorkingTimelogHistory

	selfWorkingTimelogQuery := args.ToQuery()

	repo := repository.NewWorkingTimelogRepository(&ctx, r.Db)

	err1 := repo.SelfWorkingTimelogHistory(&history, currentUser.Id, selfWorkingTimelogQuery)

	if err1 != nil {
		return nil, err
	}

	return r.SelfWorkingTimelogHistorySlideToTypes(history), nil

}
