package insightInputs

import (
	"aio-server/gql/inputs/globalInputs"
	"aio-server/models"
	"strings"
)

// ProjectInput represents input for querying user groups collection.
type ProjectInput struct {
	Input *globalInputs.PagyInput
	Query *ProjectQueryInput
}

type ProjectQueryInput struct {
	NameCont        *string
	DescriptionCont *string
	ProjectTypeCont *string
	StateCont       *string
}

func (pi *ProjectInput) ToPaginationDataAndProjectQuery() (ProjectQueryInput, models.PaginationData) {

	paginationData := pi.Input.ToPaginationInput()
	query := ProjectQueryInput{}

	if pi.Query != nil && pi.Query.NameCont != nil && strings.TrimSpace(*pi.Query.NameCont) != "" {
		query.NameCont = pi.Query.NameCont
	}

	if pi.Query != nil && pi.Query.StateCont != nil && strings.TrimSpace(*pi.Query.StateCont) != "" {
		query.StateCont = pi.Query.StateCont
	}

	if pi.Query != nil && pi.Query.DescriptionCont != nil && strings.TrimSpace(*pi.Query.DescriptionCont) != "" {
		query.DescriptionCont = pi.Query.DescriptionCont
	}

	if pi.Query != nil && pi.Query.ProjectTypeCont != nil && strings.TrimSpace(*pi.Query.ProjectTypeCont) != "" {
		query.ProjectTypeCont = pi.Query.ProjectTypeCont
	}

	return query, paginationData
}
