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

func (r *Resolver) TagDelete(ctx context.Context, args msInputs.TagDeleteInput) (*string, error) {
	user, err := auths.AuthUserFromCtx(ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	tagId, err := helpers.GqlIdToInt32(args.Id)
	if err != nil {
		return nil, exceptions.NewBadRequestError("Invalid ID")
	}

	tag := models.Tag{}
	repo := repository.NewTagRepository(&ctx, r.Db)
	err = repo.FindById(&tag, tagId)

	if err != nil {
		return nil, exceptions.NewRecordNotFoundError()
	}

	if tag.UserId != user.Id {
		return nil, exceptions.NewRecordNotFoundError()
	}

	if err := repo.Delete(&tag); err != nil {
		return nil, exceptions.NewBadRequestError(fmt.Sprintf("Cant delete this tag %s", err.Error()))
	} else {
		message := "Deleted"

		return &message, nil
	}
}
