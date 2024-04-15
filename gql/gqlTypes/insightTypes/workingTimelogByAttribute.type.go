package insightTypes

import "aio-server/gql/gqlTypes/globalTypes"

type WorkingtimelogByAttributeType struct {
	WorkingTimelog *globalTypes.WorkingTimelogType
	DataExisted    bool
}
