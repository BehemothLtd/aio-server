package insightInputs

import (
	"aio-server/gql/inputs/globalInputs"
	"aio-server/models"
)

// UserGroupsInput represents input for querying user groups collection.
type UserGroupsInput struct {
	Input *globalInputs.PagyInput
	Query *UserGroupsQueryInput
}

// ToPaginationDataAndUserGroupsQuery converts MmUserGroupsInputUserGroupsInput to UserGroupsQueryInput and models.PaginationData.
func (ugi *UserGroupsInput) ToPaginationDataAndQuery() (UserGroupsQueryInput, models.PaginationData) {
	paginationData := ugi.Input.ToPaginationInput()
	query := UserGroupsQueryInput{}

	if ugi.Query != nil && ugi.Query.TitleCont != nil {
		query.TitleCont = ugi.Query.TitleCont
	}

	return query, paginationData
}
