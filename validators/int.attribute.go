package validators

import (
	"aio-server/exceptions"
	"fmt"

	"golang.org/x/exp/constraints"
)

type IntAttribute[T constraints.Signed] struct {
	FieldAttribute
	Value     T
	AllowZero bool
}

func (attribute *IntAttribute[T]) GetCode() string {
	return attribute.Code
}

func (attribute *IntAttribute[T]) GetErrors() exceptions.ResourceModifyErrors {
	return exceptions.ResourceModifyErrors{
		Field:  attribute.Code,
		Errors: attribute.Errors,
	}
}

func (attribute *IntAttribute[T]) ValidateRequired() {
	// because Zero-Value of int is 0
	if attribute.Value == 0 {
		fmt.Printf("allow zero %+v", attribute.AllowZero)
		if !attribute.AllowZero {
			attribute.AddError("is required")
		}
	}
}

func (attribute *IntAttribute[T]) AddError(message string) {
	attribute.Errors = append(attribute.Errors, ValidationMessage(attribute.Name, message))
}

func (attribute *IntAttribute[T]) ValidateLimit(min *int, max *int64) {
	value := attribute.Value

	if min != nil {
		if value < T(*min) {
			attribute.AddError(fmt.Sprintf("is too small. Min value is %d", *min))
		}
	}

	if max != nil {
		if int64(value) > *max {
			attribute.AddError(fmt.Sprintf("is too large. Max value is %d", *max))
		}
	}
}
