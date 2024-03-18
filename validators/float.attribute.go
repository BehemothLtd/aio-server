package validators

import (
	"fmt"
	"time"

	"golang.org/x/exp/constraints"
)

type FloatAttribute[T constraints.Float] struct {
	FieldAttribute
	Value     T
	AllowZero bool
}

// GetCode returns the code of the attribute.
func (attribute *FloatAttribute[T]) GetCode() string {
	return attribute.Code
}

// GetErrors returns the errors associated with the attribute.
func (attribute *FloatAttribute[T]) GetErrors() []interface{} {
	return attribute.Errors
}

// AddError adds an error message to the attribute.
func (attribute *FloatAttribute[T]) AddError(message interface{}) {
	attribute.Errors = append(attribute.Errors, ValidationMessage(attribute.Name, message))
}

// ValidateRequired validates if the attribute is required.
func (attribute *FloatAttribute[T]) ValidateRequired() {
	if attribute.Value == 0.0 && !attribute.AllowZero {
		attribute.AddError("is required")
	}
}

func (attribute *FloatAttribute[T]) ValidateFormat(formatter string, formatterRemind string) {
	// No need to implement yet
}

func (attribute *FloatAttribute[T]) Time() *time.Time {
	return nil
}

func (attribute *FloatAttribute[T]) IsClean() bool {
	return len(attribute.Errors) == 0
}

func (attribute *FloatAttribute[T]) ValidateMin(min interface{}) {
	switch v := min.(type) {
	case float64:
		if float64(attribute.Value) < v {
			attribute.AddError(fmt.Sprintf("is too small. Min value is %f", min))
		} else {
			attribute.AddError(fmt.Sprintf("is invalid min value %f", min))
		}
	default:
		panic("Need to provide float interface{} as params")
	}
}

func (attribute *FloatAttribute[T]) ValidateMax(max interface{}) {
	switch v := max.(type) {
	case float64:
		if v < float64(attribute.Value) {
			attribute.AddError(fmt.Sprintf("is too large. Max value is %d", max))
		} else {
			attribute.AddError(fmt.Sprintf("is invalid max value %d", max))
		}
	default:
		panic("Need to provide float64 interface{} as params")
	}
}
