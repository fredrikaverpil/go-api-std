/*
Domain errors.

The domain error is a custom error type that can be used to return errors from the domain layer and have them
automatically mapped to a desired API response in the API layer. For example, if the domain layer returns
a NotFoundError, then the API layer will automatically return a 404 Not Found response for the REST API.

Each domain error carries a code. This code is used by each API's error-to-response mapper. See each API's mapper
implementation for more details.

WARNING: Anytime a domain error is returned, it is assumed this error does not contain any sensitive data that can
be considered a security risk if leaked by mistake to the client. All domain errors are assumed to be safe. If you
are unsure, don't return a domain error. Instead, return a regular error. The intent is that the API mapper should
always produce an internal server error for regular errors without revealing the details of the underlying error
to the client. Logging should instead be used to capture the details of the error.
*/

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
	// Add more error codes here and then add the mapping in each APIs error-to-response mapper
)

// Domain error
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
