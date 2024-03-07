package insightTypes

import "aio-server/gql/gqlTypes/globalTypes"

type ProjectsType struct {
	Collection *[]*globalTypes.ProjectType
	Metadata   *globalTypes.MetadataType
}
