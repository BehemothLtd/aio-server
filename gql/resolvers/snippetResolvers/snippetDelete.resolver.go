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

func (r *Resolver) SnippetDelete(ctx context.Context, args msInputs.SnippetDeleteInput) (*string, error) {
	user, err := auths.AuthUserFromCtx(ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	snippetId, err := helpers.GqlIdToInt32(args.Id)
	if err != nil {
		return nil, exceptions.NewBadRequestError("Invalid ID")
	}

	snippet := models.Snippet{
		Id: snippetId,
	}

	repo := repository.NewSnippetRepository(&ctx, r.Db)
	err = repo.FindById(&snippet, snippetId)

	if err != nil {
		return nil, exceptions.NewRecordNotFoundError()
	}

	if snippet.UserId != user.Id {
		return nil, exceptions.NewRecordNotFoundError()
	}

	if err := repo.Delete(&snippet); err != nil {
		return nil, exceptions.NewBadRequestError(fmt.Sprintf("Can not delete this snippet %s", err.Error()))
	} else {
		message := "Deleted"

		return &message, nil
	}
}
