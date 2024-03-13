package insightTypes

import "aio-server/gql/gqlTypes/globalTypes"

type LeaveDayRequestsType struct {
	Collection *[]*globalTypes.LeaveDayRequestType
	Metadata   *globalTypes.MetadataType
}
