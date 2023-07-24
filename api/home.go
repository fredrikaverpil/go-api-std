package api

import "net/http"

func (s *Server) defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, worldz!"))
}
