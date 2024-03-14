package insightTypes

import "aio-server/gql/gqlTypes/globalTypes"

type WorkingTimelogsType struct {
	Collection *[]*WorkingTimelogType
	Metadata   *globalTypes.MetadataType
}
