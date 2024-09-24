package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

// FiberLogger is the middleware for logging HTTP requests in Fiber
func FiberLogger(c *fiber.Ctx) error {
	start := time.Now()

	// Log the incoming request
	log.Printf("Started %s %s", c.Method(), c.OriginalURL())

	// Call the next handler in the stack
	err := c.Next()

	// Log the completion of the request
	log.Printf("Completed %s in %v", c.OriginalURL(), time.Since(start))

	return err
}
