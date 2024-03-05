package insightTypes

import "aio-server/gql/gqlTypes/globalTypes"

type UsersType struct {
	Collection *[]*globalTypes.UserType
	Metadata   *globalTypes.MetadataType
}
