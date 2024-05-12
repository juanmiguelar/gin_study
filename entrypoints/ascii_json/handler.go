package ascii_json

import "github.com/gin-gonic/gin"

type handler struct {}

func NewHandler(r *gin.Engine){
	h := handler{}
	
	r.GET("/ascii", h.Ascii)
}