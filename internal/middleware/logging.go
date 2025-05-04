package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logging is middleware which can be used on a http handler function to
// properly log request and respond times.
func Logging(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("started %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
		log.Printf("completed %s in %v", r.URL.Path, time.Since(start))
	}
}
