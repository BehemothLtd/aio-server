package snippetResolvers

import (
	snippetTypes "aio-server/gql/gqlTypes/snippetTypes"
	"aio-server/gql/inputs/msInputs"
	"context"
)

func (sr *Resolver) TagCreate(ctx context.Context, args msInputs.TagFormInput) (*snippetTypes.TagCreatedType, error) {

}
