package validators

import (
	"aio-server/exceptions"
	"fmt"
)

type FieldAttribute struct {
	Name   string
	Code   string
	Errors []string
}

type FieldAttributeInterface interface {
	AddError(message string)
	GetCode() string
	GetErrors() exceptions.ResourceModifyErrors
	ValidateRequired()
	ValidateLimit(min *int, max *int64)
}

func ValidationMessage(column string, message string) string {
	return fmt.Sprintf("%s %s", column, message)
}
