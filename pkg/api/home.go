package api

import "net/http"

func (s *Server) DefaultHandler(w http.ResponseWriter, r *http.Request) {
	message := "hello world"
	w.Write([]byte(message))
}
