package globalTypes

import (
	"context"

	"gorm.io/gorm"
)

type OptionType struct {
	Label *string
	Value interface{}
}

type SelectOptionType struct {
	Ctx *context.Context
	Db  *gorm.DB

	SelectOption *OptionType
}

func (sot *SelectOptionType) Label(ctx context.Context) *string {
	return sot.SelectOption.Label
}

func (sot *SelectOptionType) Value(ctx context.Context) interface{} {
	return sot.SelectOption.Value
}
