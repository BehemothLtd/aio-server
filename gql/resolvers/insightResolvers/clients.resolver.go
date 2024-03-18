package insightResolvers

import (
	"aio-server/enums"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"context"
)

// Clients resolves the query for retrieving a collection of clients.
func (r *Resolver) Clients(ctx context.Context, args insightInputs.ClientsInput) (*insightTypes.ClientsType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeClients.String(), enums.PermissionActionTypeRead.String()); err != nil {
		return nil, err
	}

	var clients []*models.Client
	clientQuery, paginationData := args.ToPaginationDataAndQuery()

	repo := repository.NewClientRepository(&ctx, r.Db)

	err := repo.List(&clients, &paginationData, clientQuery)
	if err != nil {
		return nil, err
	}

	return &insightTypes.ClientsType{
		Collection: r.ClientSliceToTypes(clients),
		Metadata: &globalTypes.MetadataType{
			Metadata: &paginationData.Metadata,
		},
	}, nil
}
