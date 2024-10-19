package rest

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fredrikaverpil/go-api-std/internal/services/user"
	"github.com/fredrikaverpil/go-api-std/internal/stores"
	"gotest.tools/v3/assert"
)

func TestX(t *testing.T) {
	assert.Assert(t, true)
}

func TestCreateUserOk(t *testing.T) {
	expectedJsonBody := `{"id":1,"username":"john"}`

	store := stores.NewDummyStore()
	userService := user.NewService(store)
	server := NewServer(":8080", *userService)
	url := "/users"
	body := bytes.NewBufferString(`{"username":"john", "password":"secret"}`)
	req, err := http.NewRequest("POST", url, body)
	assert.NilError(t, err)

	rr := httptest.NewRecorder()

	server.router.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusCreated)
	assert.Equal(t, strings.TrimSpace(rr.Body.String()), strings.TrimSpace(expectedJsonBody))
}

func TestCreateUserNoUsername(t *testing.T) {
	store := stores.NewDummyStore()
	userService := user.NewService(store)
	server := NewServer(":8080", *userService)
	url := "/users"
	body := bytes.NewBufferString(`{"password":"secret"}`)
	req, err := http.NewRequest("POST", url, body)
	assert.NilError(t, err)

	rr := httptest.NewRecorder()

	server.router.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusBadRequest)
}

func TestCreateUserNoPassword(t *testing.T) {
	store := stores.NewDummyStore()
	userService := user.NewService(store)
	server := NewServer(":8080", *userService)
	url := "/users"
	body := bytes.NewBufferString(`{"username":"john"}`)
	req, err := http.NewRequest("POST", url, body)
	assert.NilError(t, err)

	rr := httptest.NewRecorder()

	server.router.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusBadRequest)
}

func TestCreateUserUsernameTaken(t *testing.T) {
	store := stores.NewDummyStore()
	userService := user.NewService(store)
	server := NewServer(":8080", *userService)
	url := "/users"
	body := bytes.NewBufferString(`{"username":"foo", "password":"secret"}`)
	req, err := http.NewRequest("POST", url, body)
	assert.NilError(t, err)

	rr := httptest.NewRecorder()

	server.router.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusCreated)

	// User 2, with the same username
	body = bytes.NewBufferString(`{"username":"foo", "password":"secret"}`)
	req, err = http.NewRequest("POST", url, body)
	assert.NilError(t, err)

	rr = httptest.NewRecorder()

	server.router.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusConflict)
}

func TestGetUserByIdOk(t *testing.T) {
	expectedJsonBody := `{"id":1,"username":"john"}`

	store := stores.NewDummyStore()
	user_, err := store.CreateUser("john", "secret")
	assert.NilError(t, err)
	assert.Equal(t, user_.ID, 1)

	userService := user.NewService(store)
	server := NewServer(":8080", *userService)
	url := "/users/1"
	req, err := http.NewRequest("GET", url, nil)
	assert.NilError(t, err)

	rr := httptest.NewRecorder()

	server.router.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, strings.TrimSpace(rr.Body.String()), strings.TrimSpace(expectedJsonBody))
}

func TestGetUsersWithSlash(t *testing.T) {
	store := stores.NewDummyStore()
	userService := user.NewService(store)
	server := NewServer(":8080", *userService)
	url := "/users/"
	req, err := http.NewRequest("GET", url, nil)
	assert.NilError(t, err)

	rr := httptest.NewRecorder()

	server.router.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
}

func TestUsersNoSlash(t *testing.T) {
	store := stores.NewDummyStore()
	userService := user.NewService(store)
	server := NewServer(":8080", *userService)
	url := "/users"
	req, err := http.NewRequest("GET", url, nil)
	assert.NilError(t, err)

	rr := httptest.NewRecorder()

	server.router.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
}

func TestNonExistingUser(t *testing.T) {
	store := stores.NewDummyStore()
	userService := user.NewService(store)
	server := NewServer(":8080", *userService)
	url := "/users/0"
	req, err := http.NewRequest("GET", url, nil)
	assert.NilError(t, err)

	rr := httptest.NewRecorder()

	server.router.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusNotFound)
}
