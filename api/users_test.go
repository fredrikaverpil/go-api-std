package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fredrikaverpil/go-api-std/stores"
)

func TestGetUserByIdOk(t *testing.T) {
	expectedJsonBody := `{"id":1, "first_name":"John"}`

	store := stores.DummyStore{}
	server := NewServer(":8080", &store)
	url := "/users/1"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	server.handleGetUser(rr, req)

	assert.Exactly(t, rr.Code, http.StatusOK)
	assert.JSONEq(t, expectedJsonBody, rr.Body.String())
}

func TestGetUsersNotSupported(t *testing.T) {
	store := stores.DummyStore{}
	server := NewServer(":8080", &store)
	url := "/users/"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	server.handleGetUser(rr, req)

	assert.Exactly(t, rr.Code, http.StatusBadRequest)
}

func TestUsersNoSlash(t *testing.T) {
	store := stores.DummyStore{}
	server := NewServer(":8080", &store)
	url := "/users"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	server.handleGetUser(rr, req)

	assert.Exactly(t, rr.Code, http.StatusBadRequest)
}

func TestNonExistingUser(t *testing.T) {
	store := stores.DummyStore{}
	server := NewServer(":8080", &store)
	url := "/users/0"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	server.handleGetUser(rr, req)

	assert.Exactly(t, rr.Code, http.StatusNotFound)
}
