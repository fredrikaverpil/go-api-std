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

	server.router.Use(LogMiddleware)

	server.router.HandleFunc("/", server.DefaultHandler)

	staticFilesRouter := server.router.PathPrefix("/static").Subrouter()
	staticFilesRouter.PathPrefix("/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	docsRouter := server.router.PathPrefix("/docs").Subrouter()
	docsRouter.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/docs/index.html", http.StatusMovedPermanently)
	})
	docsRouter.PathPrefix("/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/static/swagger.json"),
	))

	usersRouter := server.router.PathPrefix("/users").Subrouter()
	usersRouter.Methods(http.MethodPost).Path("").HandlerFunc(server.CreateUser)
	usersRouter.Methods(http.MethodGet).Path("/{id}").HandlerFunc(server.GetUser)

	return &server
}

func (s *Server) Run() error {
	return http.ListenAndServe(s.listenAddr, s.router)
}
