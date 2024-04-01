package insightTypes

import "aio-server/gql/gqlTypes/globalTypes"

type WorkingTimelogsType struct {
	Collection *[]*globalTypes.WorkingTimelogType
	Metadata   *globalTypes.MetadataType
}
