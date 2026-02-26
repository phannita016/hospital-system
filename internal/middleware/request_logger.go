package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggingMiddleware logs request details and response time
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Record the start time
		startTime := time.Now()

		// Get request details
		path := c.Request.URL.Path
		method := c.Request.Method
		clientIP := c.ClientIP()

		// Process the request
		c.Next()

		// Calculate the duration
		duration := time.Since(startTime)
		statusCode := c.Writer.Status()

		// Log the request details
		log.Printf("[%s] %s %s - %d - %v - %s",
			method,
			path,
			clientIP,
			statusCode,
			duration,
			c.Errors.String(),
		)
	}
}
