package api

import (
	"log"
	"net/http"
)

func MiddlewareAuth(next http.Handler) http.Handler {
	log.Println("MiddlewareAuth")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
