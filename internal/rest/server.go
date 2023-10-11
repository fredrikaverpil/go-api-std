package rest

import (
	"net/http"

	"github.com/fredrikaverpil/go-api-std/internal/services/user"
	"github.com/gorilla/mux"
)

type Server struct {
	listenAddr  string
	router      *mux.Router
	userService user.UserService
}

func NewServer(listenAddr string, userService user.UserService) *Server {
	server := Server{
		listenAddr:  listenAddr,
		userService: userService,
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
