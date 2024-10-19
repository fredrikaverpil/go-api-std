package rest

import (
	"log"
	"net/http"
)

// MiddlewareServeMux wraps http.ServeMux and applies middleware to all requests
type MiddlewareServeMux struct {
	*http.ServeMux
	middlewares []func(http.HandlerFunc) http.HandlerFunc
}

func NewMiddlewareServeMux() *MiddlewareServeMux {
	return &MiddlewareServeMux{ServeMux: http.NewServeMux()}
}

func (m *MiddlewareServeMux) Use(middleware func(http.HandlerFunc) http.HandlerFunc) {
	m.middlewares = append(m.middlewares, middleware)
}

func (m *MiddlewareServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var handler http.HandlerFunc = m.ServeMux.ServeHTTP

	// Apply middlewares in reverse order
	for i := len(m.middlewares) - 1; i >= 0; i-- {
		handler = m.middlewares[i](handler)
	}

	handler.ServeHTTP(w, r)
}

// LogMiddleware is now a middleware function
func LogMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL)
		next.ServeHTTP(w, r)
	}
}
