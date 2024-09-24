package middleware

import (
	"log"
	"net/http"
	"time"
)

// HTTPLogger is the middleware for logging HTTP requests
func HTTPLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log the incoming request
		log.Printf("Started %s %s", r.Method, r.URL.Path)

		// Call the next handler
		next.ServeHTTP(w, r)

		// Log the completion of the request
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}
