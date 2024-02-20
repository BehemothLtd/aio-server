package msTypes

import (
	"aio-server/gql/gqlTypes/globalTypes"
)

type MsSnippetsType struct {
	Collection *[]*globalTypes.SnippetType
	Metadata   *globalTypes.MetadataType
}
