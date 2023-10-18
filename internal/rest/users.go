package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type createUserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// CreateUser godoc
// @Summary      Create user
// @Description  Create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        payload   body  createUserPayload  true  "User payload"
// @Success      200  {object}  models.User
// @Failure      400  {string}  string  "Bad Request"
// @Failure      409  {string}  string  "Conflict"
// @Failure      500  {string}  string  "Internal Server Error"
// @Router       /users [post]
func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	var payload createUserPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Failed to parse JSON payload", http.StatusBadRequest)
		return
	}
	if payload.Username == "" || payload.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	user, err := s.userService.CreateUser(payload.Username, payload.Password)
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

// GetUser godoc
// @Summary      Get user data
// @Description  get user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  models.User
// @Failure      404  {string}  string  "Not Found"
// @Failure      400  {string}  string  "Bad Request"
// @Failure      500  {string}  string  "Internal Server Error"
// @Router       /users/{id} [get]
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

	user, err := s.userService.GetUser(userId)
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
