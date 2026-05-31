package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

const slowThreshold = 200 * time.Millisecond

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)

		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		ip := c.ClientIP()

		if latency > slowThreshold {
			log.Printf(
				"[SLOW API] %s | %d | %v | %s | %s",
				method,
				status,
				latency,
				ip,
				path,
			)
			return
		}

		log.Printf(
			"[HTTP] %s | %d | %v | %s | %s",
			method,
			status,
			latency,
			ip,
			path,
		)
	}
}