package rest

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/fredrikaverpil/go-api-std/internal/services"
	"github.com/gorilla/mux"
)

type CreateUserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func validatePayload(payload *CreateUserPayload) error {
	if payload.Username == "" || payload.Password == "" {
		return errors.New("username and password are required")
	}
	return nil
}

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	var payload CreateUserPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Failed to parse JSON payload", http.StatusBadRequest)
		return
	}
	if err := validatePayload(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := services.CreateUser(s.store, payload.Username, payload.Password)
	if err != nil {
		validHTTPStatuses := []int{http.StatusNotFound, http.StatusConflict}
		mapErrorToRESTResponse(err, validHTTPStatuses, w)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to write JSON response", http.StatusInternalServerError)
	}
}

func (s *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdStr := vars["id"]

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

	user, err := services.GetUser(s.store, userId)
	// TODO: do not allow getting the user unless the user id is part of the JWT
	if err != nil {
		validHTTPStatuses := []int{http.StatusNotFound}
		mapErrorToRESTResponse(err, validHTTPStatuses, w)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to write JSON response", http.StatusInternalServerError)
	}
}
