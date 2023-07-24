package api

import (
	"net/http"

	"github.com/fredrikaverpil/go-api-std/stores"
)

type Server struct {
	listenAddr string
	store      stores.Store
}

func NewServer(listenAddr string, store stores.Store) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *Server) ListenAddr() string {
	return s.listenAddr
}

func (s *Server) Run() error {
	http.HandleFunc("/", s.defaultHandler)
	http.HandleFunc("/users/", s.handleGetUser)
	return http.ListenAndServe(s.listenAddr, nil)
}
