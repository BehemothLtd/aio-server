package inputs

import graphql "github.com/graph-gophers/graphql-go"

// MsSnippetFavoriteInput represents args for toggle favorite on a snippet
type MsSnippetFavoriteInput struct {
	Id graphql.ID
}
