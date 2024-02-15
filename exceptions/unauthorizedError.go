package exceptions

import (
	"aio-server/pkg/constants"
	"fmt"
)

// UnauthorizedError represents an unauthorized access error.
type UnauthorizedError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error returns the error message.
func (e *UnauthorizedError) Error() string {
	return fmt.Sprintf("error [%d]: %s", e.Code, e.Message)
}

// Extensions returns additional data associated with the error.
func (e *UnauthorizedError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}
}

// NewUnauthorizedError creates a new UnauthorizedError instance with the provided message.
// If the message is empty, it uses the default error message.
func NewUnauthorizedError(message string) *UnauthorizedError {
	if message == "" {
		message = constants.UnauthorizedErrorMsg
	}

	return &UnauthorizedError{
		Code:    constants.UnauthorizedErrorCode,
		Message: message,
	}
}
