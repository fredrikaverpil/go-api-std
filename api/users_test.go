package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fredrikaverpil/go-api-std/stores"
)

func TestUserAPI(t *testing.T) {
	store := stores.DummyStore{}
	server := NewServer(":8080", &store)

	url := "/users/1"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	server.handleGetUser(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("got %v want %v", rr.Code, http.StatusOK)
	}

	expectedJsonBody := `{"id":1, "first_name":"John"}`

	assert.JSONEq(t, expectedJsonBody, rr.Body.String())
}
