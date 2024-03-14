package validators

import (
	"fmt"
	"strings"
	"time"
)

// StringAttribute represents a string attribute validator.
type StringAttribute struct {
	FieldAttribute

	Value string
}

// GetCode returns the code of the string attribute.
func (attribute *StringAttribute) GetCode() string {
	return attribute.Code
}

// GetErrors returns the errors of the string attribute.
func (attribute *StringAttribute) GetErrors() []interface{} {
	return attribute.Errors
}

// AddError adds an error to the string attribute.
func (attribute *StringAttribute) AddError(message interface{}) {
	attribute.Errors = append(attribute.Errors, ValidationMessage(attribute.Name, message))
}

// ValidateRequired validates if the string attribute is required.
func (attribute *StringAttribute) ValidateRequired() {
	if attribute.Value == "" || strings.TrimSpace(attribute.Value) == "" {
		attribute.AddError("is required")
	}
}

// ValidateLimit validates the length limits of the string attribute.
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

func (attribute *StringAttribute) ValidateFormat(formatter string, formatterRemind string) {
	// No need to implement yet
}

func (attribute *StringAttribute) Time() *time.Time {
	return nil
}

func (attribute *StringAttribute) IsClean() bool {
	return len(attribute.Errors) == 0
}
