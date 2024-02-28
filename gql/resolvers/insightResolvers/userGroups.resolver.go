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

// UserGroups resolves the query for retrieving a collection of userGroups.
func (r *Resolver) UserGroups(ctx context.Context, args insightInputs.UserGroupsInput) (*insightTypes.UserGroupsType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeUserGroups.String(), enums.PermissionActionTypeRead.String()); err != nil {
		return nil, err
	}

	var userGroups []*models.UserGroup
	userGroupsQuery, paginationData := args.ToPaginationDataAndUserGroupsQuery()

	repo := repository.NewUserGroupRepository(&ctx, r.Db)

	err := repo.List(&userGroups, &paginationData, &userGroupsQuery)
	if err != nil {
		return nil, err
	}

	return &insightTypes.UserGroupsType{
		Collection: r.UserGroupSliceToTypes(userGroups),
		Metadata: &globalTypes.MetadataType{
			Metadata: &paginationData.Metadata,
		},
	}, nil
}
