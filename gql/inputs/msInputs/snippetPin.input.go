package msInputs

import graphql "github.com/graph-gophers/graphql-go"

// SnippetPinInput represents args for toggle pin on a snippet
type SnippetPinInput struct {
	Id graphql.ID
}
