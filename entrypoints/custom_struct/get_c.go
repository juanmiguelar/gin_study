package custom_struct

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type structC struct {
	NestedStructPointer *StructA
	FieldC              string `form:"field_c"`
}

func (h handler) GetDataC(c *gin.Context) {
	var b structC
	c.Bind(&b)
	c.JSON(http.StatusOK, gin.H{
		"a": b.NestedStructPointer,
		"c": b.FieldC,
	})
}