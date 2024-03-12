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

func (r *Resolver) Users(ctx context.Context, args insightInputs.UsersInput) (*insightTypes.UsersType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeUsers.String(), enums.PermissionActionTypeRead.String()); err != nil {
		return nil, err
	}

	var users []*models.User

	usersQuery, paginationData := args.ToPaginationDataAndQuery()

	repo := repository.NewUserRepository(&ctx, r.Db)

	err := repo.List(&users, &paginationData, usersQuery)
	if err != nil {
		return nil, err
	}

	return &insightTypes.UsersType{
		Collection: r.UsersSliceToTypes(users),
		Metadata: &globalTypes.MetadataType{
			Metadata: &paginationData.Metadata,
		},
	}, nil
}
