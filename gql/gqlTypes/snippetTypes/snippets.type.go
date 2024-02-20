package snippetTypes

import (
	"aio-server/gql/gqlTypes/globalTypes"
)

type SnippetsType struct {
	Collection *[]*globalTypes.SnippetType
	Metadata   *globalTypes.MetadataType
}
