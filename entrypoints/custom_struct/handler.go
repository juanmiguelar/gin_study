package custom_struct

import "github.com/gin-gonic/gin"

type handler struct {}

func NewHandler(r *gin.Engine){
	h := handler{}
	
	r.GET("/getb", h.GetDataB)
    r.GET("/getc", h.GetDataC)
    r.GET("/getd", h.GetDataD)
}