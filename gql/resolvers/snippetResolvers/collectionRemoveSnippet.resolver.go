package snippetResolvers

import (
	"aio-server/gql/inputs/msInputs"
	"aio-server/services/msServices"
	"context"
)

func (r *Resolver) CollectionRemoveSnippet(ctx context.Context, args msInputs.CollectionRemoveSnippetInput) (bool, error) {
	service := msServices.CollectionRemoveSnippetService{
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
