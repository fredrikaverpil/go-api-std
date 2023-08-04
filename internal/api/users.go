package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/fredrikaverpil/go-api-std/internal/lib"
	"github.com/fredrikaverpil/go-api-std/internal/services"
)

type CreateUserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// TODO: use JSON:API spec for JSON error responses

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
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

	user, err := services.CreateUser(s.store, payload.Username, payload.Password)
	if err != nil {
		ierr := err.(*lib.CustomError)
		switch ierr.Code {
		case lib.ErrNotFound:
			http.Error(w, ierr.Message, http.StatusNotFound)
		case lib.ErrConflict:
			http.Error(w, ierr.Message, http.StatusConflict)
		default:
			http.Error(w, ierr.Message, http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
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
		ierr := err.(*lib.CustomError)
		switch ierr.Code {
		case lib.ErrNotFound:
			http.Error(w, ierr.Message, http.StatusNotFound)
			return
		default:
			http.Error(w, ierr.Message, http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
