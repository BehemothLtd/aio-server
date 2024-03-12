package validators

import "aio-server/pkg/specialTypes"

// FieldAttribute represents a field attribute with its name, code, and errors.
type FieldAttribute struct {
	Name   string
	Code   string
	Errors specialTypes.FieldAttributeErrorType
}

// FieldAttributeInterface defines methods for working with field attributes.
type FieldAttributeInterface interface {
	AddError(message string)
	AddItemsError(index int, field string, message string)
	GetCode() string
	GetErrors() specialTypes.FieldAttributeErrorType

	// Validators
	ValidateRequired()
	ValidateLimit(min *int, max *int64)
}

// ValidationMessage returns a formatted validation message.
func ValidationMessage(column string, message string) string {
	return message
}
