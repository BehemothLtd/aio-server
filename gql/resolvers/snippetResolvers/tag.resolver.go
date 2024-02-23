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

func (r *Resolver) Tag(ctx context.Context, args msInputs.TagInput) (*globalTypes.TagType, error) {
	if args.Id == "" {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}

	// Authenticate the user
	user, err := auths.AuthUserFromCtx(ctx)
	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	tagId, err := helpers.GqlIdToInt32(args.Id)
	if err != nil {
		return nil, err
	}

	tag := models.Tag{
		Id: tagId,
	}

	repo := repository.NewTagRepository(&ctx, r.Db)
	err = repo.FindById(&tag, tagId)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exceptions.NewRecordNotFoundError()
		}

		return nil, err
	}

	if tag.UserId != user.Id {
		return nil, exceptions.NewRecordNotFoundError()
	}

	return &globalTypes.TagType{Tag: &tag}, nil
}
