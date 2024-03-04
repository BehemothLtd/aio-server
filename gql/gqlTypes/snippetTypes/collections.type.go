package snippetTypes

import "aio-server/gql/gqlTypes/globalTypes"

type CollectionsType struct {
	Collection *[]*globalTypes.CollectionType
	Metadata   *globalTypes.MetadataType
}
