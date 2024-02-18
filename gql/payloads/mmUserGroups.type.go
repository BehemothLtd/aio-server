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

func (mug *MmUserGroupsResolver) Resolve() error {
	var userGroups []*models.UserGroup
	userGroupsQuery, paginationData := mug.Args.ToPaginationDataAndUserGroupsQuery()

	repo := repository.NewUserGroupRepository(mug.Ctx, mug.Db)

	err := repo.List(&userGroups, &paginationData, &userGroupsQuery)
	if err != nil {
		return err
	}

	mug.Collection = mug.fromUserGroups(userGroups)
	mug.Metadata = &MetadataResolver{Metadata: &paginationData.Metadata}

	return nil
}

// fromUserGroups converts models.UserGroup slice to []*MmUserGroupsResolver.
func (mug *MmUserGroupsResolver) fromUserGroups(userGroups []*models.UserGroup) *[]*MmUserGroupResolver {
	resolvers := make([]*MmUserGroupResolver, len(userGroups))
	for i, s := range userGroups {
		resolvers[i] = &MmUserGroupResolver{UserGroup: s}
	}

	return &resolvers
}
