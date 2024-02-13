package exceptions

import (
	"aio-server/pkg/constants"
	"fmt"
)

type UnauthorizedError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e UnauthorizedError) Error() string {
	return fmt.Sprintf("error [%d]: %s", e.Code, e.Message)
}

func (e UnauthorizedError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}
}

func NewUnauthorizedError(message string) UnauthorizedError {
	var returnMessage string

	if message == "" {
		returnMessage = constants.UnauthorizedErrorMsg
	} else {
		returnMessage = message
	}

	return UnauthorizedError{
		Code:    constants.UnauthorizedErrorCode,
		Message: returnMessage,
	}
}
