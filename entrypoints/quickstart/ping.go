package quickstart

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}