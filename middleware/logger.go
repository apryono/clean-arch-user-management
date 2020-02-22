package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logging function use to log
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%v] Called with method %v at %v \n", r.URL.Path, r.Method, time.Now())
		next.ServeHTTP(w, r)
	})
}
