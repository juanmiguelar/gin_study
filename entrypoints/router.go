package entrypoints

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juanmiguelar/gin_study/entrypoints/ascii_json"
	"github.com/juanmiguelar/gin_study/entrypoints/custom_struct"
	"github.com/juanmiguelar/gin_study/entrypoints/quickstart"
	"github.com/juanmiguelar/gin_study/entrypoints/uri"
)

func CreateRoutes() *gin.Engine {
	r := gin.New()
	r.Use(Logger())
	quickstart.NewHandler(r)
	ascii_json.NewHandler(r)
	custom_struct.NewHandler(r)
	uri.NewHandler(r)
	return r
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		// after request
		latency := time.Since(t)
		log.Printf("Latency: %s", latency.String())

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
