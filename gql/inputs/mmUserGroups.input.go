package inputs

import "aio-server/models"

// MmUserGroupsInput represents input for querying user groups collection.
type MmUserGroupsInput struct {
	Input *PagyInput
	Query *MmUserGroupQueryInput
}

// ToPaginationDataAndUserGroupsQuery converts MmUserGroupsInput to models.UserGroupsQuery and models.PaginationData.
func (mugi *MmUserGroupsInput) ToPaginationDataAndUserGroupsQuery() (models.UserGroupsQuery, models.PaginationData) {
	paginationData := mugi.Input.ToPaginationInput()

	var titleCont string
	if mugi.Query != nil && mugi.Query.TitleCont != nil {
		titleCont = *mugi.Query.TitleCont
	}

	return models.UserGroupsQuery{TitleCont: titleCont}, paginationData
}
