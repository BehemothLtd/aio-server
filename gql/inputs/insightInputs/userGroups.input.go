package insightInputs

import (
	"aio-server/gql/inputs/globalInputs"
	"aio-server/models"
)

// UserGroupsInput represents input for querying user groups collection.
type UserGroupsInput struct {
	Input *globalInputs.PagyInput
	Query *UserGroupQueryInput
}

// ToPaginationDataAndUserGroupsQuery converts MmUserGroupsInputUserGroupsInput to models.UserGroupsQuery and models.PaginationData.
func (ugi *UserGroupsInput) ToPaginationDataAndUserGroupsQuery() (models.UserGroupsQuery, models.PaginationData) {
	paginationData := ugi.Input.ToPaginationInput()

	var titleCont string
	if ugi.Query != nil && ugi.Query.TitleCont != nil {
		titleCont = *ugi.Query.TitleCont
	}

	return models.UserGroupsQuery{TitleCont: titleCont}, paginationData
}
