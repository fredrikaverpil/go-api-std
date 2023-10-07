package rest

import (
	"log"
	"net/http"

	"github.com/fredrikaverpil/go-api-std/internal/domain"
)

func mapErrorToRESTResponse(err error, validHTTPStatuses []int, w http.ResponseWriter) {
	codeToHTTPStatus := map[domain.ErrorCode]int{
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
