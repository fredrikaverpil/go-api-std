package lib

type ErrorCode int

const (
	ErrInternal     ErrorCode = iota
	ErrNotFound     ErrorCode = iota
	ErrInvalidInput ErrorCode = iota
	// Add more error codes here
)

type internalError struct {
	Code    ErrorCode
	Message string
}

func InternalError(message string) *internalError {
	return &internalError{Code: ErrInternal, Message: message}
}

func NotFoundError(message string) *internalError {
	return &internalError{Code: ErrNotFound, Message: message}
}

func InvalidInputError(message string) *internalError {
	return &internalError{Code: ErrInvalidInput, Message: message}
}
