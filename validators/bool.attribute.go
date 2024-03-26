package validators

import (
	"time"
)

// BoolAttribute represents a bool attribute validator.
type BoolAttribute struct {
	FieldAttribute

	Value bool
}

// GetCode returns the code of the bool attribute.
func (attribute *BoolAttribute) GetCode() string {
	return attribute.Code
}

// GetErrors returns the errors of the bool attribute.
func (attribute *BoolAttribute) GetErrors() []interface{} {
	return attribute.Errors
}

// AddError adds an error to the bool attribute.
func (attribute *BoolAttribute) AddError(message interface{}) {
	attribute.Errors = append(attribute.Errors, ValidationMessage(attribute.Name, message))
}

// ValidateRequired validates if the bool attribute is required.
func (attribute *BoolAttribute) ValidateRequired() {
}

func (attribute *BoolAttribute) ValidateFormat(formatter string, formatterRemind string) {
	// No need to implement yet
}

func (attribute *BoolAttribute) Time() *time.Time {
	return nil
}

func (attribute *BoolAttribute) IsClean() bool {
	return len(attribute.Errors) == 0
}

func (attribute *BoolAttribute) ValidateMin(min interface{}) {
}

func (attribute *BoolAttribute) ValidateMax(max interface{}) {
}
