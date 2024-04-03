package insightResolvers

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"
	"fmt"
)

func (r *Resolver) ClientDelete(ctx context.Context, args insightInputs.ClientInput) (*string, error) {
	_, err := r.Authorize(ctx, string(enums.PermissionTargetTypeClients), string(enums.PermissionActionTypeWrite))

	if err != nil {
		return nil, err
	}
	if args.Id == "" {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}

	requestId, err := helpers.GqlIdToInt32(args.Id)
	if err != nil {
		return nil, err
	}

	client := models.Client{
		Id:     requestId,
	}
	repo := repository.NewClientRepository(&ctx, r.Db)

	if err := repo.Find(&client); err != nil {
		return nil, exceptions.NewRecordNotFoundError()
	}

	if err = repo.Destroy(&client); err != nil {
		return nil, exceptions.NewBadRequestError(fmt.Sprintf("Can not delete this client %s", err.Error()))
	} else {
		message := "Deleted"

		return &message, nil
	}
}
