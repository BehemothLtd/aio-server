package validators

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

	// Validators
	ValidateRequired()
	ValidateLimit(min *int, max *int64)
}

// ValidationMessage returns a formatted validation message.
func ValidationMessage(column string, message interface{}) interface{} {
	return message
}
