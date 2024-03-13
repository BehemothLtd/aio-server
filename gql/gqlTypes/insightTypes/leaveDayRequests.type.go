package insightTypes

import "aio-server/gql/gqlTypes/globalTypes"

type LeaveDayRequestsType struct {
	Collection *[]*globalTypes.LeaveDayRequestType
	Metadate   *globalTypes.MetadataType
}
