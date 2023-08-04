package lib

import (
	"fmt"
)

type ErrorCode int

const (
	ErrInternal     ErrorCode = iota
	ErrNotFound     ErrorCode = iota
	ErrConflict     ErrorCode = iota
	ErrInvalidInput ErrorCode = iota
	// Add more error codes here
)

type CustomError struct {
	Code    ErrorCode
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

func InternalError(message string) error {
	return &CustomError{Code: ErrInternal, Message: message}
}

func NotFoundError(message string) error {
	return &CustomError{Code: ErrNotFound, Message: message}
}

func ConflictError(message string) error {
	return &CustomError{Code: ErrConflict, Message: message}
}

func InvalidInputError(message string) error {
	return &CustomError{Code: ErrInvalidInput, Message: message}
}
