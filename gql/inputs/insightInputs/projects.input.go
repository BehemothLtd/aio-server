package insightInputs

import (
	"aio-server/enums"
	"aio-server/gql/inputs/globalInputs"
	"aio-server/models"
	"strings"
)

type ProjectsQueryInput struct {
	NameCont        *string
	DescriptionCont *string
	ProjectTypeEq   *string
	StateEq         *string
}

type ProjectsInput struct {
	Input *globalInputs.PagyInput
	Query *ProjectsQueryInput
}

func (pi *ProjectsInput) ToPaginationDataAndQuery() (ProjectsQueryInput, models.PaginationData) {
	paginationData := pi.Input.ToPaginationInput()
	query := ProjectsQueryInput{}

	if pi.Query != nil {
		if pi.Query.NameCont != nil && strings.TrimSpace(*pi.Query.NameCont) != "" {
			query.NameCont = pi.Query.NameCont
		}

		if pi.Query.DescriptionCont != nil && strings.TrimSpace(*pi.Query.DescriptionCont) != "" {
			query.DescriptionCont = pi.Query.DescriptionCont
		}

		if pi.Query.ProjectTypeEq != nil && strings.TrimSpace(*pi.Query.ProjectTypeEq) != "" {
			if projectTypeEq, err := enums.ParseProjectType(*pi.Query.ProjectTypeEq); err == nil {
				query.ProjectTypeEq = (*string)(&projectTypeEq)
			}
		}

		if pi.Query.StateEq != nil && strings.TrimSpace(*pi.Query.StateEq) != "" {
			if stateEq, err := enums.ParseProjectState(*pi.Query.StateEq); err == nil {
				query.StateEq = (*string)(&stateEq)
			}
		}
	}

	return query, paginationData
}
