package insightInputs

import (
	"aio-server/gql/inputs/globalInputs"
	"aio-server/models"
)

type DeviceTypesInput struct {
	Input *globalInputs.PagyInput
}

func (dti *DeviceTypesInput) ToPaginationData() models.PaginationData {
	paginationData := dti.Input.ToPaginationInput()

	return paginationData
}
