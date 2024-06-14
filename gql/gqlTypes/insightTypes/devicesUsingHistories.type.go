package insightTypes

import "aio-server/gql/gqlTypes/globalTypes"

type DevicesUsingHistoriesType struct {
	Collection *[]*globalTypes.DevicesUsingHistoryType
	Metadata   *globalTypes.MetadataType
}
