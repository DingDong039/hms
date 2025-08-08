package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger returns a middleware for logging request/response details
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		// Get status code and client IP
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method

		// Format the raw query if it exists
		if raw != "" {
			path = path + "?" + raw
		}

		// Log the request details
		fmt.Printf("[HMS] %v | %3d | %13v | %15s | %-7s %#v\n",
			end.Format("2006/01/02 - 15:04:05"),
			statusCode,
			latency,
			clientIP,
			method,
			path,
		)
	}
}
