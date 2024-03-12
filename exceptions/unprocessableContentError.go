package exceptions

import (
	"aio-server/pkg/constants"
	"fmt"
)

// ResourceModifyErrors represents errors associated with modifying a resource.
type ResourceModifyErrors struct {
	Field  string   `json:"field"`
	Errors []string `json:"errors"`
}

type ResourceModificationError map[string][]interface{}

// UnprocessableContentError represents an unprocessable content error.
type UnprocessableContentError struct {
	Code    int                       `json:"code"`
	Message string                    `json:"message"`
	Errors  ResourceModificationError `json:"errors"`
}

// Error returns the error message.
func (e UnprocessableContentError) Error() string {
	return fmt.Sprintf("error [%d]: %s", e.Code, e.Message)
}

// Extensions returns additional data associated with the error.
func (e UnprocessableContentError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
		"errors":  e.Errors,
	}
}

// NewUnprocessableContentError creates a new UnprocessableContentError instance with the provided message and errors.
// If the message is empty, it uses the default error message.
func NewUnprocessableContentError(message string, errors ResourceModificationError) *UnprocessableContentError {
	if message == "" {
		message = constants.UnprocessableContentErrorMsg
	}

	return &UnprocessableContentError{
		Code:    constants.UnprocessableContentErrorCode,
		Message: message,
		Errors:  errors,
	}
}

// AddError adds a new ResourceModifyErrors to the UnprocessableContentError.
func (e *UnprocessableContentError) AddError(field string, errors []interface{}) {
	if e.Errors == nil {
		e.Errors = ResourceModificationError{}
	}

	if e.Errors[field] == nil {
		e.Errors[field] = errors
	} else {

		e.Errors[field] = append(e.Errors[field], errors...)
	}
}
