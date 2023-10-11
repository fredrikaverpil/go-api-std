package rest

import (
	"log"
	"net/http"

	"github.com/fredrikaverpil/go-api-std/internal/domain"
)

// Map custom domain errors (and regular errors) to REST responses
func mapErrorToRESTResponse(err error, validHTTPStatuses []int, w http.ResponseWriter) {
	// To avoid leaking undesired errors, the caller to this function must provide a list of valid HTTP status codes.
	// If the error code is not in the list, then a generic 500 Internal Server Error is returned, without revealing
	// the error message.
	//
	// So, in short, use regular errors (errors.New()) for errors that should never reach the client.
	//
	// The reasoning behind this is that each endpoint usually needs to add documentation on which possible
	// HTTP status codes it can return. By forcing the caller to provide a list of valid HTTP status codes,
	// we can ensure that the caller has thought about which HTTP status codes are valid for the endpoint.
	//
	// One might argue that it is more important to instead provide a list of valid domain error codes, so to avoid
	// leaking internal errors. However, this is not always possible. For example, if the endpoint is a POST endpoint,
	// then it is not possible to know which error codes are valid, since the error codes are dependent on the data
	// provided in the request body.

	codeToHTTPStatus := map[domain.ErrorCode]int{
		domain.ErrInternal:        http.StatusInternalServerError,
		domain.ErrNotFound:        http.StatusNotFound,
		domain.ErrConflict:        http.StatusConflict,
		domain.ErrInvalidArgument: http.StatusBadRequest,
		// Add more mappings here as you add more custom error codes
	}

	if customErr, ok := err.(*domain.Error); ok {
		if httpStatus, exists := codeToHTTPStatus[customErr.Code]; exists {
			for _, validStatus := range validHTTPStatuses {
				if httpStatus == validStatus {
					http.Error(w, customErr.Message, httpStatus)
					return
				}
			}
		}
	}

	log.Printf("Internal server error: %v", err.Error())
	http.Error(w, "Internal server error", http.StatusInternalServerError)
}
