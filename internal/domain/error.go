package domain

import (
	"fmt"
)

type ErrorCode int

const (
	ErrInternal        ErrorCode = iota
	ErrNotFound        ErrorCode = iota
	ErrConflict        ErrorCode = iota
	ErrInvalidArgument ErrorCode = iota
	// Add more error codes here
)

type Error struct {
	Code    ErrorCode
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

func InternalError(message string) error {
	return &Error{Code: ErrInternal, Message: message}
}

func NotFoundError(message string) error {
	return &Error{Code: ErrNotFound, Message: message}
}

func ConflictError(message string) error {
	return &Error{Code: ErrConflict, Message: message}
}

func InvalidArgumentError(message string) error {
	return &Error{Code: ErrInvalidArgument, Message: message}
}
