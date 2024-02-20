package msInputs

import graphql "github.com/graph-gophers/graphql-go"

// SnippetFavoriteInput represents args for toggle favorite on a snippet
type SnippetFavoriteInput struct {
	Id graphql.ID
}
