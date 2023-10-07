package rest

import "net/http"

func (s *Server) DefaultHandler(w http.ResponseWriter, r *http.Request) {
	message := "hello world"
	if _, err := w.Write([]byte(message)); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
