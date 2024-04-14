package insightTypes

import "aio-server/gql/gqlTypes/globalTypes"

type TimeSheet struct {
	ProjectData *globalTypes.ProjectType
	UserData    []*globalTypes.UserType
}
