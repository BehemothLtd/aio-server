package validators

import "aio-server/gql/inputs/insightInputs"

type NestedSlices interface {
	insightInputs.ProjectIssueStatusInputForProjectCreate | insightInputs.ProjectAssigneeInputForProjectCreate
}

type SliceAttribute[T NestedSlices] struct {
	FieldAttribute
	Value *[]T
}

// GetCode returns the code of the attribute.
func (attribute *SliceAttribute[T]) GetCode() string {
	return attribute.Code
}

// GetErrors returns the errors associated with the attribute.
func (attribute *SliceAttribute[T]) GetErrors() []interface{} {
	return attribute.Errors
}

// AddError adds an error message to the attribute.
func (attribute *SliceAttribute[T]) AddError(message interface{}) {
	attribute.Errors = append(attribute.Errors, ValidationMessage(attribute.Name, message))
}

// ValidateRequired validates if the attribute is required.
func (attribute *SliceAttribute[T]) ValidateRequired() {
	if attribute.Value == nil || len(*attribute.Value) == 0 {
		attribute.AddError("is required")
	}
}

// ValidateLimit validates if the attribute value is within the specified limits.
func (attribute *SliceAttribute[T]) ValidateLimit(min *int, max *int64) {
	// No need to implement
}
