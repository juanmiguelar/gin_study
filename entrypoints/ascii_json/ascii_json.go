package ascii_json

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) Ascii(c *gin.Context){
	data := map[string]string{
		"lang": "GO语言",
		"tag":"<br>",
	}
	c.AsciiJSON(http.StatusOK, data)
}