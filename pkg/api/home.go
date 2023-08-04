package api

import (
	"net/http"

	"github.com/fredrikaverpil/go-api-std/pkg/services"
)

func (s *Server) DefaultHandler(w http.ResponseWriter, r *http.Request) {
	message := services.Home()
	w.Write([]byte(message))
}
