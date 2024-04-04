package insightTypes

import "aio-server/gql/gqlTypes/globalTypes"

type AttendancesType struct {
	Collection *[]*globalTypes.AttendanceType
	Metadata   *globalTypes.MetadataType
}
