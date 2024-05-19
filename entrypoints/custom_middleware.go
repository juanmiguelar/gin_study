package entrypoints

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Printf("Latency: [%v]", latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Printf("Status: [%v]", status)
	}
}