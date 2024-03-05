package validators

// FieldAttribute represents a field attribute with its name, code, and errors.
type FieldAttribute struct {
	Name   string
	Code   string
	Errors []string
}

// FieldAttributeInterface defines methods for working with field attributes.
type FieldAttributeInterface interface {
	AddError(message string)
	GetCode() string
	GetErrors() []string

	// Validators
	ValidateRequired()
	ValidateLimit(min *int, max *int64)
}

// ValidationMessage returns a formatted validation message.
func ValidationMessage(column string, message string) string {
	return message
}
