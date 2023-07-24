package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		if len(r.URL.Path) <= len("/users/") {
			// do not accept /users or /users/ as valid paths
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		userIdStr := r.URL.Path[len("/users/"):]

		int64NumOfDigits := 19
		if len(userIdStr) >= int64NumOfDigits {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := s.store.GetUser(userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		validationErr := user.Validate()
		if validationErr != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)

	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
}
