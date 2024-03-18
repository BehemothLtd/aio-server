package validators

import (
	"fmt"
	"time"

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
func (attribute *IntAttribute[T]) GetErrors() []interface{} {
	return attribute.Errors
}

// AddError adds an error message to the attribute.
func (attribute *IntAttribute[T]) AddError(message interface{}) {
	attribute.Errors = append(attribute.Errors, ValidationMessage(attribute.Name, message))
}

// ValidateRequired validates if the attribute is required.
func (attribute *IntAttribute[T]) ValidateRequired() {
	if attribute.Value == 0 && !attribute.AllowZero {
		attribute.AddError("is required")
	}
}

func (attribute *IntAttribute[T]) ValidateFormat(formatter string, formatterRemind string) {
	// No need to implement yet
}

func (attribute *IntAttribute[T]) Time() *time.Time {
	return nil
}

func (attribute *IntAttribute[T]) IsClean() bool {
	return len(attribute.Errors) == 0
}

func (attribute *IntAttribute[T]) ValidateMin(min interface{}) {
	switch v := min.(type) {
	case int64:
		if int64(attribute.Value) < v {
			attribute.AddError(fmt.Sprintf("is too small. Min value is %d", min))
		}
	default:
		panic("Need to provide int64 interface{} as params")
	}
}

func (attribute *IntAttribute[T]) ValidateMax(max interface{}) {
	switch v := max.(type) {
	case int64:
		if v < int64(attribute.Value) {
			attribute.AddError(fmt.Sprintf("is too large. Max value is %d", max))
		}
	default:
		panic("Need to provide int64 interface{} as params")
	}
}
