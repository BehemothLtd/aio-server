package insightInputs

import (
	"aio-server/gql/inputs/globalInputs"
	"aio-server/models"
)

type UsersInput struct {
	Input *globalInputs.PagyInput
	Query *UserQueryInput
}

func (ui *UsersInput) ToPaginationDataAndUsersQuery() (models.UsersQuery, models.PaginationData) {
	paginationData := ui.Input.ToPaginationInput()
	query := models.UsersQuery{}

	if ui.Query != nil && ui.Query.NameCont != nil {
		query.NameCont = ui.Query.NameCont
	}

	if ui.Query != nil && ui.Query.FullNameCont != nil {
		query.FullNameCont = ui.Query.FullNameCont
	}

	if ui.Query != nil && ui.Query.EmailCont != nil {
		query.EmailCont = ui.Query.EmailCont
	}

	if ui.Query != nil && ui.Query.SlackIdCont != nil {
		query.SlackIdCont = ui.Query.SlackIdCont
	}

	if ui.Query != nil && ui.Query.StateEq != nil {
		query.StateEq = ui.Query.StateEq
	}

	return query, paginationData
}
