package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/fredrikaverpil/go-api-std/stores"
)

type Server struct {
	listenAddr string
	store      stores.Store
	router     *mux.Router
}

func NewServer(listenAddr string, store stores.Store) *Server {
	server := Server{
		listenAddr: listenAddr,
		store:      store,
	}

	server.router = mux.NewRouter()
	server.router.HandleFunc("/", server.DefaultHandler)

	usersRouter := server.router.PathPrefix("/users").Subrouter()
	usersRouter.Methods(http.MethodPost).Path("").HandlerFunc(server.CreateUser)
	usersRouter.Methods(http.MethodGet).Path("/{id}").HandlerFunc(server.GetUser)

	return &server
}

func (s *Server) Run() error {
	return http.ListenAndServe(s.listenAddr, s.router)
}
