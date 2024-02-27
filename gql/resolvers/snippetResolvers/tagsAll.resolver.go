package snippetResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/repository"
	"context"
	"fmt"
)

func (r *Resolver) TagsAll(ctx context.Context) (*[]*globalTypes.TagType, error) {
	// Authenticate the user
	_, err := auths.AuthUserFromCtx(ctx)
	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	tags := []*models.Tag{}
	repo := repository.NewTagRepository(&ctx, r.Db.Preload("Snippets"))

	if err := repo.ListAll(&tags); err != nil {
		return nil, exceptions.NewBadRequestError(fmt.Sprintf("Error happened %s", err.Error()))
	} else {
		// tagTypes := []globalTypes.TagType{}
		result := make([]*globalTypes.TagType, len(tags))
		for i, t := range tags {
			result[i] = &globalTypes.TagType{Tag: t}
		}

		return &result, nil
	}
}
