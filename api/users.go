package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		userId := r.URL.Path[len("/users/"):]
		userIdInt, err := strconv.Atoi(userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := s.store.GetUser(userIdInt)
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
	} else {
		// POST, UPDATE, PATCH, DELETE...
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
}
