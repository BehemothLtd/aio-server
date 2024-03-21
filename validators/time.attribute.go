package validators

import (
	"fmt"
	"time"
)

type TimeAttribute struct {
	FieldAttribute
	Value     string
	TimeValue *time.Time
}

// GetCode returns the code of the attribute.
func (attribute *TimeAttribute) GetCode() string {
	return attribute.Code
}

// GetErrors returns the errors of the string attribute.
func (attribute *TimeAttribute) GetErrors() []interface{} {
	return attribute.Errors
}

// AddError adds an error to the string attribute.
func (attribute *TimeAttribute) AddError(message interface{}) {
	attribute.Errors = append(attribute.Errors, ValidationMessage(attribute.Name, message))
}

// ValidateRequired validates if the string attribute is required.
func (attribute *TimeAttribute) ValidateRequired() {
	if attribute.Value == "" {
		attribute.AddError("is required")
	}
}

func (attribute *TimeAttribute) ValidateFormat(formatter string, formatterRemind string) {
	if attribute.Value != "" {
		if timeValue, err := time.Parse(formatter, attribute.Value); err != nil {
			attribute.AddError(fmt.Sprintf("need to be formatted as %s", formatterRemind))
		} else {
			attribute.TimeValue = &timeValue
		}
	}
}

func (attribute *TimeAttribute) Time() *time.Time {
	return attribute.TimeValue
}

func (attribute *TimeAttribute) IsClean() bool {
	return len(attribute.Errors) == 0
}

func (attribute *TimeAttribute) ValidateMin(min interface{}) {
	switch v := min.(type) {
	case time.Time:
		if v.After(*attribute.TimeValue) {
			attribute.AddError(fmt.Sprintf("is invalid. Need to be after %+v", v))
		}
	default:
		panic("Need to provide time interface{} as params")
	}
}

func (attribute *TimeAttribute) ValidateMax(max interface{}) {
	switch v := max.(type) {
	case time.Time:
		if v.Before(*attribute.TimeValue) {
			attribute.AddError(fmt.Sprintf("is invalid. Need to be after %+v", v))
		}
	default:
		panic("Need to provide time interface{} as params")
	}
}
