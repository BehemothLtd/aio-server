package insightTypes

import "aio-server/gql/gqlTypes/globalTypes"

type ProjectExperiencesType struct {
	Collection *[]*globalTypes.ProjectExperienceType
	Metadata   *globalTypes.MetadataType
}
