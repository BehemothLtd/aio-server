package validators

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/pkg/specialTypes"
)

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
func (attribute *SliceAttribute[T]) GetErrors() specialTypes.FieldAttributeErrorType {
	return attribute.Errors
}

// AddError adds an error message to the attribute.
func (attribute *SliceAttribute[T]) AddError(message string) {
	attribute.Errors.Base = append(attribute.Errors.Base, ValidationMessage(attribute.Name, message))
}

func (attribute *SliceAttribute[T]) AddItemsError(index int, field string, message string) {
	if attribute.Errors.Items == nil {
		attribute.Errors.Items = map[int]map[string][]string{}
	}

	if attribute.Errors.Items[index] == nil {
		// init items errors
		attribute.Errors.Items[index] = make(map[string][]string)

		if attribute.Errors.Items[index][field] == nil {
			attribute.Errors.Items[index][field] = []string{}
		}
	}

	attribute.Errors.Items[index][field] = append(attribute.Errors.Items[index][field], message)
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
