package insightTypes

import "aio-server/gql/gqlTypes/globalTypes"

type UserGroupsType struct {
	Collection *[]*globalTypes.UserGroupType
	Metadata   *globalTypes.MetadataType
}
