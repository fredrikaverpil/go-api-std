package rest

import (
	"net/http"

	"github.com/fredrikaverpil/go-api-std/internal/services/user"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type Server struct {
	listenAddr  string
	router      *MiddlewareServeMux
	userService user.UserService
}

func NewServer(listenAddr string, userService user.UserService) *Server {
	server := Server{
		listenAddr:  listenAddr,
		userService: userService,
		router:      NewMiddlewareServeMux(),
	}

	// Add middleware for every request
	server.router.Use(LogMiddleware)

	// Default handler
	server.router.HandleFunc("/", server.DefaultHandler)

	// serve all static files at /static from the ./static folder
	staticFolderHandler := http.FileServer(http.Dir("./static"))
	server.router.Handle("/static/", http.StripPrefix("/static/", staticFolderHandler))

	// swagger docs at /swagger/index.html
	swaggerHandler := httpSwagger.Handler(httpSwagger.URL("/static/swagger.json"))
	server.router.Handle("/swagger/", http.StripPrefix("/swagger/", swaggerHandler))

	// users
	server.router.HandleFunc("POST /users", server.CreateUser)
	server.router.HandleFunc("GET /users/{id}", server.GetUser)

	return &server
}

func (s *Server) Run() error {
	return http.ListenAndServe(s.listenAddr, s.router)
}
