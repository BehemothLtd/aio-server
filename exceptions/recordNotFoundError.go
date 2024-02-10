package exceptions

import (
	"aio-server/pkg/constants"
	"fmt"
)

type RecordNotFoundError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e RecordNotFoundError) Error() string {
	return fmt.Sprintf("error [%d]: %s", e.Code, e.Message)
}

func (e RecordNotFoundError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}
}

func NewRecordNotFoundError() RecordNotFoundError {
	return RecordNotFoundError{
		Code:    constants.NotFoundErrorCode,
		Message: constants.NotFoundErrorMsg,
	}
}
