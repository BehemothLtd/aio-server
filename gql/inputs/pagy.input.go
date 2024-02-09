package inputs

import (
	"aio-server/models"
)

type PagyInput struct {
	PerPage *int32
	Page    *int32
}

func (input *PagyInput) ToPaginationInput() models.PaginationData {
	paginationInput := models.PaginationData{}

	if input != nil {
		if input.Page != nil {
			paginationInput.Metadata.Page = int(*input.Page)
		} else {
			paginationInput.Metadata.Page = 1
		}

		if input.PerPage != nil {
			paginationInput.Metadata.PerPage = int(*input.PerPage)
		} else {
			paginationInput.Metadata.PerPage = 10
		}
	}

	return paginationInput
}
