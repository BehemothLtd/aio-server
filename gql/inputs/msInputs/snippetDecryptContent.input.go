package msInputs

import graphql "github.com/graph-gophers/graphql-go"

type SnippetDecryptContentInput struct {
	Id      graphql.ID
	Passkey string
}
