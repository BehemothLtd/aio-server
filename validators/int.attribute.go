package validators

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type IntAttribute[T constraints.Signed] struct {
	FieldAttribute
	Value     T
	AllowZero bool
}

// GetCode returns the code of the attribute.
func (attribute *IntAttribute[T]) GetCode() string {
	return attribute.Code
}

// GetErrors returns the errors associated with the attribute.
func (attribute *IntAttribute[T]) GetErrors() []string {
	return attribute.Errors
}

// GetValue returns the value of attribute
func (attribute *IntAttribute[T]) GetValue() interface{} {
	return attribute.Value
}

// AddError adds an error message to the attribute.
func (attribute *IntAttribute[T]) AddError(message string) {
	attribute.Errors = append(attribute.Errors, ValidationMessage(attribute.Name, message))
}

// ValidateRequired validates if the attribute is required.
func (attribute *IntAttribute[T]) ValidateRequired() {
	if attribute.Value == 0 && !attribute.AllowZero {
		attribute.AddError("is required")
	}
}

// ValidateLimit validates if the attribute value is within the specified limits.
func (attribute *IntAttribute[T]) ValidateLimit(min *int, max *int64) {
	value := int64(attribute.Value)

	if min != nil && int64(attribute.Value) < int64(*min) {
		attribute.AddError(fmt.Sprintf("is too small. Min value is %d", *min))
	}

	if max != nil && value > *max {
		attribute.AddError(fmt.Sprintf("is too large. Max value is %d", *max))
	}
}
