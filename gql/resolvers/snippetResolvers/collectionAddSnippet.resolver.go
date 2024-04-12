package snippetResolvers

import (
	"aio-server/gql/inputs/msInputs"
	"aio-server/services/msServices"
	"context"
)

func (r *Resolver) CollectionAddSnippet(ctx context.Context, args msInputs.CollectionAddSnippetInput) (bool, error) {
	service := msServices.CollectionAddSnippetService{
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
