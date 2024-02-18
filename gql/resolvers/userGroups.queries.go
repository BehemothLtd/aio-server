package gql

import (
	"aio-server/gql/inputs"
	"aio-server/gql/payloads"
	"context"

	// graphql "github.com/graph-gophers/graphql-go"
)


// MmUserGroups resolves the query for retrieving a collection of user groups.
func (r *Resolver) MmUserGroups(ctx context.Context, args inputs.MmUserGroupsInput) (*payloads.MmUserGroupsResolver, error) {
	resolver := payloads.MmUserGroupsResolver{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	if err := resolver.Resolve(); err != nil {
		return nil, err
	}

	return &resolver, nil
}
