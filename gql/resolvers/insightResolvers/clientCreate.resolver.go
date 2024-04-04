package insightResolvers

import (
	"aio-server/enums"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/services/insightServices"
	"context"
)

func (r *Resolver) ClientCreate(ctx context.Context, args insightInputs.ClientCreateInput) (*insightTypes.ClientCreatedType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeClients.String(), enums.PermissionActionTypeWrite.String()); err != nil {
		return nil, err
	}

	Client := models.Client{}

	service := insightServices.ClientCreateService{
		Ctx:    &ctx,
		Db:     r.Db,
		Args:   args,
		Client: &Client,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	} else {
		return &insightTypes.ClientCreatedType{
			Client: &globalTypes.ClientType{
				Client: &Client,
			},
		}, nil
	}
}
