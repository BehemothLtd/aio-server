package globalInputs

import (
	"aio-server/models"
)

// PagyInput represents input for pagination.
type PagyInput struct {
	PerPage *int32
	Page    *int32
}

// ToPaginationInput converts PagyInput to models.PaginationData.
func (input *PagyInput) ToPaginationInput() models.PaginationData {
	paginationInput := models.PaginationData{
		Metadata: models.Metadata{
			Page:    1,
			PerPage: 10,
		},
	}

	if input != nil {
		if input.Page != nil {
			paginationInput.Metadata.Page = int(*input.Page)
		}

		if input.PerPage != nil {
			paginationInput.Metadata.PerPage = int(*input.PerPage)
		}
	}

	return paginationInput
}
