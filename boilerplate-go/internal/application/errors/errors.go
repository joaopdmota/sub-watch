package app_errors

import (
	"errors"
	"fmt"
)

var (
	ErrUnknown              = errors.New("unknown error")
	ErrBadRequest           = errors.New("bad request")
	ErrUnprocessableEntity = errors.New("unprocessable entity")
	ErrNotFound             = errors.New("not found")
)

type AppError struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Path    string `json:"path,omitempty"`
	Message string `json:"message,omitempty"`
	Err     error  `json:"-"`
}

func (e AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e AppError) Unwrap() error {
	return e.Err
}

type ErrorsResponseDTO struct {
	Errors []AppError `json:"errors"`
}
