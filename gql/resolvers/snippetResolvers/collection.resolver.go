package snippetResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/inputs/msInputs"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/pkg/helpers"
	"aio-server/repository"

	"context"

	"gorm.io/gorm"
)

func (r *Resolver) Collection(ctx context.Context, args msInputs.CollectionInput) (*globalTypes.CollectionType, error) {
	if args.Id == "" {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}

	user, err := auths.AuthUserFromCtx(ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	collectionId, err := helpers.GqlIdToInt32(args.Id)
	if err != nil {
		return nil, err
	}

	collection := models.Collection{
		Id: collectionId,
	}

	repo := repository.NewCollectionRepository(&ctx, r.Db.Preload("Snippets"))

	err = repo.FindByUser(&collection, user.Id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exceptions.NewRecordNotFoundError()
		}

		return nil, err
	}

	if collection.UserId != user.Id {
		return nil, exceptions.NewRecordNotFoundError()
	}

	return &globalTypes.CollectionType{Collection: &collection}, nil
}
