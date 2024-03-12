package insightTypes

import "aio-server/gql/gqlTypes/globalTypes"

type DeviceTypesType struct {
	Collection *[]*globalTypes.DeviceTypeType
	Metadata   *globalTypes.MetadataType
}
