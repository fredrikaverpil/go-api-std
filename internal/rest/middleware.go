package rest

import (
	"log"
	"net/http"
)

func LogMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL)

		next.ServeHTTP(w, r)
	}
}
