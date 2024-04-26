package insightTypes

import "aio-server/gql/gqlTypes/globalTypes"

type IssuesType struct {
	Collection *[]*globalTypes.IssueType
	Metadata   *globalTypes.MetadataType
}
