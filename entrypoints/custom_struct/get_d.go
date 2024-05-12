package custom_struct

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type structD struct {
	NestedAnonyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func (h handler) GetDataD(c *gin.Context) {
	var b structD
	c.Bind(&b)
	c.JSON(http.StatusOK, gin.H{
		"x": b.NestedAnonyStruct,
		"d": b.FieldD,
	})
}
