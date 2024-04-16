package insightServices

import (
	"aio-server/gql/inputs/insightInputs"

	"gorm.io/gorm"
)

type FetchSelectOptionsService struct {
	Db   *gorm.DB
	Args insightInputs.SelectOptionsInput
}

func (fsos *FetchSelectOptionsService) Execute() error {
	return nil
}
