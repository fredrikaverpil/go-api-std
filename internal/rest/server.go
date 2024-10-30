package rest

import (
	"net/http"

	"github.com/fredrikaverpil/go-api-std/internal/services/user"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type Server struct {
	listenAddr  string
	mux         *MiddlewareServeMux
	userService user.UserService
}

func NewServer(listenAddr string, userService user.UserService) *Server {
	server := Server{
		listenAddr:  listenAddr,
		userService: userService,
		mux:         NewMiddlewareServeMux(),
	}

	// Add middleware for every request
	server.mux.Use(LogMiddleware)

	// Default handler
	server.mux.HandleFunc("/", server.DefaultHandler)

	// serve all static files at /static from the ./static folder
	staticFolderHandler := http.FileServer(http.Dir("./static"))
	server.mux.Handle("/static/", http.StripPrefix("/static/", staticFolderHandler))

	// swagger docs at /swagger/index.html
	swaggerHandler := httpSwagger.Handler(httpSwagger.URL("/static/swagger.json"))
	server.mux.Handle("/swagger/", http.StripPrefix("/swagger/", swaggerHandler))

	// users
	server.mux.HandleFunc("POST /users", server.CreateUser)
	server.mux.HandleFunc("GET /users/{id}", server.GetUser)

	return &server
}

func (s *Server) Run() error {
	return http.ListenAndServe(s.listenAddr, s.mux)
}
