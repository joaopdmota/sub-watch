package app_errors

import "fmt"

const (
	ERROR_UNKNOW               = "UNKNOW"
	ERROR_INTERNAL_SERVER_ERROR = "INTERNAL_SERVER_ERROR"
	ERROR_BAD_REQUEST          = "BAD_REQUEST"
	ERROR_UNPROCESSABLE_ENTITY = "UNPROCESSABLE_ENTITY"
	ERROR_NOT_FOUND            = "NOT_FOUND"
	ERROR_INVALID_AUTHENTICATION_TOKEN = "INVALID_AUTHENTICATION_TOKEN"
	
)

type Error struct {
	Code    int                    `json:"code"`
	Type    string                 `json:"type,omitempty"`
	Path    string                 `json:"path,omitempty"`
	Message string                 `json:"message,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
	Err     error                  `json:"-"`
}

func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *Error) Unwrap() error {
	return e.Err
}

type Errors []Error

type ErrorsResponseDTO struct {
	Errors Errors `json:"errors"`
}

func CreateErrors(errs ...Error) Errors {
	var errors Errors

	for _, err := range errs {
		errors = append(errors, err)
	}

	return errors
}

func NewInternalError(message string, err error) *Error {
	return &Error{
		Code:    500,
		Type:    ERROR_INTERNAL_SERVER_ERROR,
		Message: message,
		Err:     err,
	}
}

func NewNotFoundError(message string, err error) *Error {
	return &Error{
		Code:    404,
		Type:    ERROR_NOT_FOUND,
		Message: message,
		Err:     err,
	}
}

func NewBadRequestError(message string, err error) *Error {
	return &Error{
		Code:    400,
		Type:    ERROR_BAD_REQUEST,
		Message: message,
		Err:     err,
	}
}
