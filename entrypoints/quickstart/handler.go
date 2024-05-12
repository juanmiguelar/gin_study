package quickstart

import "github.com/gin-gonic/gin"

type handler struct {

}

func NewHandler(r *gin.Engine){
	h := handler{}
	
	r.GET("/ping", h.Ping)
}