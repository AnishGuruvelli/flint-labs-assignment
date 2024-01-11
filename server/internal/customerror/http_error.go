package customerror

import (
	"fmt"
)

type CustomError struct {
	IsSuccess bool        `json:"is_success"`
	Message   string      `json:"message"`
	Code      string      `json:"code"`
	Data      []ErrorData `json:"data"`
}

type InternalErrorHandler struct {
	Error     error
	Code      int
	ErrorCode string
}

type ErrorData struct {
	Code        string      `json:"code"`
	Description string      `json:"description"`
	MetaData    interface{} `json:"metadata"`
}

func NewCustomError(msg string, code string, data []ErrorData) *CustomError {
	return &CustomError{
		IsSuccess: false,
		Code:      code,
		Message:   msg,
		Data:      data,
	}
}

func (ce *CustomError) Error() string {
	return fmt.Sprintf("New Error - %s", ce.Message)
}
