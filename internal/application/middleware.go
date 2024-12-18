package application

import (
	"log"
	"net/http"
)

func MethodMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)

		method := r.Method
		if method == http.MethodPost {
			next.ServeHTTP(w, r)
		}
	})
}
