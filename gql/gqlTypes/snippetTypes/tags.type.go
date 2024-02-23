package snippetTypes

import (
	"aio-server/gql/gqlTypes/globalTypes"
)

type TagsType struct {
	Collection *[]*globalTypes.TagType
	Metadata   *globalTypes.MetadataType
}
