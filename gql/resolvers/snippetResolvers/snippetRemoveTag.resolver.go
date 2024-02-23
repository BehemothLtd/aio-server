package snippetResolvers

import (
	"aio-server/gql/inputs/msInputs"
	"aio-server/services/msServices"
	"context"
)

// SnippetAddTag resolves the mutation for removing tag into snippet
func (r *Resolver) SnippetRemoveTag(ctx context.Context, args msInputs.SnippetRemoveTagInput) (bool, error) {
	service := msServices.SnippetRemoveTagService{
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
