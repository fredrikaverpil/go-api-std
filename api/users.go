package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type CreateUserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *Server) Users(w http.ResponseWriter, r *http.Request) {
	// TODO: use JSON:API spec for JSON error responses

	switch r.Method {

	case http.MethodPost:

		if len(r.URL.Path) > len("/users/") {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		var payload CreateUserPayload
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, "Failed to parse JSON payload", http.StatusBadRequest)
			return
		}
		if payload.Username == "" || payload.Password == "" {
			http.Error(w, "Username and password are required", http.StatusBadRequest)
			return
		}

		user, err := s.store.GetUserByUsername(payload.Username)
		if err == nil && user.Username == payload.Username {
			http.Error(w, "Username is already taken", http.StatusConflict)
			return
		}

		user, error := s.store.CreateUser(payload.Username, payload.Password)
		if error != nil {
			http.Error(w, error.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)

	case http.MethodGet:

		if len(r.URL.Path) <= len("/users/") {
			// do not accept /users or /users/ as valid paths
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		userIdStr := r.URL.Path[len("/users/"):]

		int64NumOfDigits := 19
		if len(userIdStr) >= int64NumOfDigits {
			http.Error(w, "Bad request: too many digits", http.StatusBadRequest)
			return
		}

		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			http.Error(w, "Bad request: not an integer", http.StatusBadRequest)
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
