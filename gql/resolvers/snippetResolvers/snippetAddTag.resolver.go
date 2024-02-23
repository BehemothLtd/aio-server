package snippetResolvers

import (
	"aio-server/gql/inputs/msInputs"
	"aio-server/services/msServices"
	"context"
)

// SnippetAddTag resolves the mutation for adding tag into snippet
func (r *Resolver) SnippetAddTag(ctx context.Context, args msInputs.SnippetAddTagInput) (bool, error) {
	service := msServices.SnippetAddTagService{
		Ctx:  ctx,
		Db:   *r.Db,
		Args: args,
	}

	if result, err := service.Execute(); err != nil {
		return false, err
	} else {
		return result, err
	}
}
