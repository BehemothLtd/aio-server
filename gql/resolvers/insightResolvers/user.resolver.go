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

func (r *Resolver) User(ctx context.Context, args insightInputs.UserInput) (*globalTypes.UserType, error) {
	if args.Id == "" {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}

	userId, err := helpers.GqlIdToInt32(args.Id)

	if err != nil {
		return nil, err
	}

	user := models.User{Id: userId}
	repo := repository.NewUserRepository(&ctx, r.Db)
	err = repo.Find(&user)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exceptions.NewRecordNotFoundError()
		}
		return nil, err
	}

	return &globalTypes.UserType{User: &user}, nil
}
