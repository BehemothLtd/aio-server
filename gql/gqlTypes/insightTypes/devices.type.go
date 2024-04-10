package insightTypes

import "aio-server/gql/gqlTypes/globalTypes"

type DevicesType struct {
	Collection *[]*globalTypes.DeviceType
	Metadata   *globalTypes.MetadataType
}
