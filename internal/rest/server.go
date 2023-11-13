package rest

import (
	"net/http"

	"github.com/fredrikaverpil/go-api-std/internal/services/user"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
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
		router:      mux.NewRouter(),
	}

	// middleware for all requests
	server.router.Use(LogMiddleware)

	// catch-all
	server.router.HandleFunc("/", server.DefaultHandler)

	// serve all static files at /static from the ./static folder
	staticFolderHandler := http.FileServer(http.Dir("./static"))
	staticHandler := http.StripPrefix("/static/", staticFolderHandler)
	server.router.PathPrefix("/static/").Handler(staticHandler)

	// swagger docs at /swagger/index.html
	swaggerHandler := httpSwagger.Handler(httpSwagger.URL("/static/swagger.json"))
	docsHandler := http.StripPrefix("/swagger/", swaggerHandler)
	server.router.PathPrefix("/swagger/").Handler(docsHandler)

	// users
	usersRouter := server.router.PathPrefix("/users").Subrouter()
	usersRouter.Methods(http.MethodPost).HandlerFunc(server.CreateUser)
	usersRouter.Methods(http.MethodGet).Path("/{id}").HandlerFunc(server.GetUser)

	return &server
}

func (s *Server) Run() error {
	return http.ListenAndServe(s.listenAddr, s.router)
}
