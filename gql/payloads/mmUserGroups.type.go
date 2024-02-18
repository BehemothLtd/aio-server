package payloads

import (
	"aio-server/gql/inputs"
	"aio-server/models"
	"aio-server/repository"
	"context"

	"gorm.io/gorm"
)

// MmUserGroupsResolver resolves the querying of user groups collection.
type MmUserGroupsResolver struct {
	Ctx        *context.Context
	Db         *gorm.DB
	Args       inputs.MmUserGroupsInput
	Collection *[]*MmUserGroupResolver
	Metadata   *MetadataResolver
}

func (mugr *MmUserGroupsResolver) Resolve() error {
	var userGroups []*models.UserGroup
	userGroupsQuery, paginationData := mugr.Args.ToPaginationDataAndUserGroupsQuery()

	repo := repository.NewUserGroupRepository(mugr.Ctx, mugr.Db)

	err := repo.List(&userGroups, &paginationData, &userGroupsQuery)
	if err != nil {
		return err
	}

	mugr.Collection = mugr.fromUserGroups(userGroups)
	mugr.Metadata = &MetadataResolver{Metadata: &paginationData.Metadata}

	return nil
}

// fromUserGroups converts models.UserGroup slice to []*MmUserGroupResolver.
func (mugr *MmUserGroupsResolver) fromUserGroups(userGroups []*models.UserGroup) *[]*MmUserGroupResolver {
	resolvers := make([]*MmUserGroupResolver, len(userGroups))
	for i, s := range userGroups {
		resolvers[i] = &MmUserGroupResolver{UserGroup: s}
	}

	return &resolvers
}
