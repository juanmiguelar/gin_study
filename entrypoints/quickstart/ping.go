package quickstart

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) Ping(c *gin.Context) {
	result := c.GetString("example")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%v%v", "ping", result),
	})
}