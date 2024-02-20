package insightResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"

	"gorm.io/gorm"
)

func (r *Resolver) UserGroup(ctx context.Context, args insightInputs.UserGroupInput) (*globalTypes.UserGroupType, error) {
	if args.Id == "" {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}

	userGroupId, err := helpers.GqlIdToInt32(args.Id)

	if err != nil {
		return nil, err
	}

	userGroup := models.UserGroup{}
	repo := repository.NewUserGroupRepository(&ctx, r.Db)
	err = repo.FindById(&userGroup, userGroupId)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exceptions.NewRecordNotFoundError()
		}
		return nil, err
	}

	return &globalTypes.UserGroupType{UserGroup: &userGroup}, nil
}
