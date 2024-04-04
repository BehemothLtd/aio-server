package insightResolvers

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/services/insightServices"
	"context"
)

func (r *Resolver) ClientUpdate(ctx context.Context, args insightInputs.ClientUpdateInput) (*insightTypes.ClientUpdatedType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeClients.String(), enums.PermissionActionTypeWrite.String()); err != nil {
		return nil, err
	}

	if args.Id == "" {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}

	clientId, err := helpers.GqlIdToInt32(args.Id)

	if err != nil {
		return nil, err
	}

	client := models.Client{Id: clientId}
	service := insightServices.ClientUpdateService{
		Ctx:         &ctx,
		Db:          r.Db,
		Args:        args,
		Client: &client,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	} else {
		return &insightTypes.ClientUpdatedType{
			Client: &globalTypes.ClientType{
				Client: &client,
			},
		}, nil
	}
}
