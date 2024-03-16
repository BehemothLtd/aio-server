package insightTypes

import "aio-server/gql/gqlTypes/globalTypes"

type IssueStatusesType struct {
	Collection *[]*globalTypes.IssueStatusType
	Metadata   *globalTypes.MetadataType
}
