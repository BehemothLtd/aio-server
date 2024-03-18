package insightTypes

import "aio-server/gql/gqlTypes/globalTypes"

type ClientsType struct {
	Collection *[]*globalTypes.ClientType
	Metadata   *globalTypes.MetadataType
}
