package entrypoints

import (
	"github.com/gin-gonic/gin"
	"github.com/juanmiguelar/gin_study/entrypoints/ascii_json"
	"github.com/juanmiguelar/gin_study/entrypoints/custom_struct"
	"github.com/juanmiguelar/gin_study/entrypoints/quickstart"
)

func CreateRoutes() *gin.Engine {
	r := gin.Default()
	quickstart.NewHandler(r)
	ascii_json.NewHandler(r)
	custom_struct.NewHandler(r)
	return r
}
