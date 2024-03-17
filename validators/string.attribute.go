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

func (attribute *StringAttribute) ValidateFormat(formatter string, formatterRemind string) {
	// No need to implement yet
}

func (attribute *StringAttribute) Time() *time.Time {
	return nil
}

func (attribute *StringAttribute) IsClean() bool {
	return len(attribute.Errors) == 0
}

func (attribute *StringAttribute) ValidateMin(min interface{}) {
	switch v := min.(type) {
	case int64:
		if int64(len(attribute.Value)) < v {
			attribute.AddError(fmt.Sprintf("is too short. Min length is %d", min))
		} else {
			attribute.AddError(fmt.Sprintf("is invalid min length %d", min))
		}
	default:
		panic("Need to provide int64 interface{} as params")
	}
}

func (attribute *StringAttribute) ValidateMax(max interface{}) {
	switch v := max.(type) {
	case int64:
		if v < int64(len(attribute.Value)) {
			attribute.AddError(fmt.Sprintf("is too long. Max length is %d", max))
		} else {
			attribute.AddError(fmt.Sprintf("is invalid max length %d", max))
		}
	default:
		panic("Need to provide int64 interface{} as params")
	}
}
