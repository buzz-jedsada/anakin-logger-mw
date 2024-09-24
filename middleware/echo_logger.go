package middleware

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

// EchoLogger is the middleware for logging HTTP requests in Echo
func EchoLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		// Log the incoming request
		log.Printf("Started %s %s", c.Request().Method, c.Request().RequestURI)

		// Call the next handler in the stack
		err := next(c)

		// Log the completion of the request
		log.Printf("Completed %s in %v", c.Request().RequestURI, time.Since(start))

		return err
	}
}
