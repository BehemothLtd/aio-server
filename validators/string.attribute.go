package validators

import (
	"aio-server/exceptions"
	"fmt"
)

type StringAttribute struct {
	FieldAttribute
	Value string
}

func (attribute *StringAttribute) GetCode() string {
	return attribute.Code
}

func (attribute *StringAttribute) GetErrors() exceptions.ResourceModifyErrors {
	return exceptions.ResourceModifyErrors{
		Field:  attribute.Code,
		Errors: attribute.Errors,
	}
}

func (attribute *StringAttribute) ValidateRequired() {
	if attribute.Value == "" {
		attribute.AddError("is required")
	}
}

func (attribute *StringAttribute) AddError(message string) {
	attribute.Errors = append(attribute.Errors, ValidationMessage(attribute.Name, message))
}

func (attribute *StringAttribute) ValidateLimit(min *int, max *int64) {
	value := attribute.Value

	if min != nil {
		if len(value) < *min {
			attribute.AddError(fmt.Sprintf("is too short. Min characters is %d", *min))
		}
	}

	if max != nil {
		if int64(len(value)) > *max {
			attribute.AddError(fmt.Sprintf("is too long. Max characters is %d", *max))
		}
	}
}
