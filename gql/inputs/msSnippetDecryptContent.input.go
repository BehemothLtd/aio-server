package inputs

import graphql "github.com/graph-gophers/graphql-go"

type MsSnippetDecryptContentInput struct {
	Id      graphql.ID
	Passkey string
}
