package validators

import "time"

// FieldAttribute represents a field attribute with its name, code, and errors.
type FieldAttribute struct {
	Name   string
	Code   string
	Errors []interface{}
}

// FieldAttributeInterface defines methods for working with field attributes.
type FieldAttributeInterface interface {
	AddError(message interface{})
	GetCode() string
	GetErrors() []interface{}
	Time() *time.Time
	IsClean() bool

	// Validators
	ValidateRequired()
	ValidateMin(min interface{})
	ValidateMax(max interface{})
	ValidateFormat(formatter string, formatterRemind string)
}

// ValidationMessage returns a formatted validation message.
func ValidationMessage(column string, message interface{}) interface{} {
	return message
}
