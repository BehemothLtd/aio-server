package snippetResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/msInputs"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"
	"fmt"
)

func (r *Resolver) CollectionDelete(ctx context.Context, args msInputs.CollectionDeleteInput) (*string, error) {
	user, err := auths.AuthUserFromCtx(ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	collectionId, err := helpers.GqlIdToInt32(args.Id)
	if err != nil {
		return nil, exceptions.NewBadRequestError("Invalid ID")
	}

	collection := models.Collection{
		Id: collectionId,
	}

	repo := repository.NewCollectionRepository(&ctx, r.Db)
	err = repo.FindByUser(&collection, user.Id)

	if err != nil {
		return nil, exceptions.NewRecordNotFoundError()
	}

	if collection.UserId != user.Id {
		return nil, exceptions.NewRecordNotFoundError()
	}

	if err := repo.Delete(&collection); err != nil {
		return nil, exceptions.NewBadRequestError(fmt.Sprintf("Can not delete this collection %s", err.Error()))
	} else {
		message := "Deleted"

		return &message, nil
	}
}
