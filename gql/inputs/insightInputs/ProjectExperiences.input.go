package insightInputs

import (
	"aio-server/gql/inputs/globalInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"

	"github.com/graph-gophers/graphql-go"
)

type ProjectExperiencesInput struct {
	Input *globalInputs.PagyInput
	Query *ProjectExperiencesFrontQueryInput
}

type ProjectExperiencesFrontQueryInput struct {
	ProjectIdEq *graphql.ID
}

type ProjectExperiencesQueryInput struct {
	ProjectIdEq *int32
}

func (pei *ProjectExperiencesInput) ToPaginationAndQueryData() (ProjectExperiencesQueryInput, models.PaginationData) {
	paginationData := pei.Input.ToPaginationInput()
	query := ProjectExperiencesQueryInput{}

	if pei.Query != nil {
		if pei.Query.ProjectIdEq != nil {
			if ProjectIdEq, err := helpers.GqlIdToInt32(*pei.Query.ProjectIdEq); err != nil {
				query.ProjectIdEq = nil
			} else {
				query.ProjectIdEq = &ProjectIdEq
			}
		}
	}

	return query, paginationData
}
