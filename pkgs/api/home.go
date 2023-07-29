package api

import (
	"net/http"
)

func (s *Server) DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, worldz!"))
}
